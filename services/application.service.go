package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/enums"
	"github.com/jerson2000/jquest/models"
	"github.com/jerson2000/jquest/repositories"
	"github.com/jerson2000/jquest/responses"
	"gorm.io/gorm"
)

type ApplicationService interface {
	ApplyJob(ctx context.Context, userId int, dto dtos.ApplicationCreateRequestDto) responses.ResultResponse[dtos.ApplicationResponseDto]
	GetMyApplications(ctx context.Context, userId, page, limit int) responses.ResultResponse[responses.PaginatedResponse[dtos.ApplicationResponseDto]]
	GetJobApplications(ctx context.Context, userId, jobId, page, limit int) responses.ResultResponse[responses.PaginatedResponse[dtos.ApplicationResponseDto]]
	UpdateApplicationStatus(ctx context.Context, userId int, applicationId int, dto dtos.ApplicationUpdateStatusRequestDto) responses.ResultResponse[dtos.ApplicationResponseDto]
}

type applicationService struct {
	applicationRepo repositories.ApplicationRepository
	jobRepo         repositories.JobRepository
	candidateRepo   repositories.CandidateRepository
	recruiterRepo   repositories.RecruiterRepository
	userRepo        repositories.UserRepository
}

func NewApplicationService() ApplicationService {
	db := config.Database
	return &applicationService{
		applicationRepo: repositories.NewApplicationRepository(db),
		jobRepo:         repositories.NewJobRepository(db),
		candidateRepo:   repositories.NewCandidateRepository(db),
		recruiterRepo:   repositories.NewRecruiterRepository(db),
		userRepo:        repositories.NewUserRepository(db),
	}
}

func (s *applicationService) ApplyJob(ctx context.Context, userId int, dto dtos.ApplicationCreateRequestDto) responses.ResultResponse[dtos.ApplicationResponseDto] {
	var responseDto dtos.ApplicationResponseDto

	err := config.Database.Transaction(func(tx *gorm.DB) error {
		// Use repositories with transaction
		candidateRepo := repositories.NewCandidateRepository(tx)
		applicationRepo := repositories.NewApplicationRepository(tx)
		jobRepo := repositories.NewJobRepository(tx) // Read-only but good practice to use same tx

		// 1. Get/Create Candidate
		candidate, err := candidateRepo.GetByUserID(ctx, userId)
		if err != nil {
			// Create candidate profile if not exists
			userRepo := repositories.NewUserRepository(tx) // Use tx
			user, err := userRepo.GetByID(ctx, userId)
			if err != nil {
				return err // Rollback
			}
			newCandidate := models.Candidate{
				UserId:    userId,
				FirstName: user.Name,
				LastName:  "", // Placeholder, ideally specific
				Email:     user.Email,
			}
			candidate, err = candidateRepo.Create(ctx, newCandidate)
			if err != nil {
				return err
			}
		}

		// 2. Check Job
		_, err = jobRepo.GetByID(ctx, dto.JobId)
		if err != nil {
			return err
		}

		// 3. Check if already applied
		_, err = applicationRepo.GetByCandidateAndJob(ctx, candidate.Id, dto.JobId)
		if err == nil {
			return errors.New("Already applied")
		}

		// 4. Create Application
		application := models.Application{
			CandidateId: candidate.Id,
			JobId:       dto.JobId,
			Status:      string(enums.ApplicationStatusPending),
			AppliedAt:   time.Now(),
		}

		newApp, err := applicationRepo.Create(ctx, application)
		if err != nil {
			return err
		}

		responseDto = s.toDto(newApp)
		return nil
	})

	if err != nil {
		if err.Error() == "Already applied" {
			return responses.Failure[dtos.ApplicationResponseDto](http.StatusConflict, "You have already applied for this job")
		}
		return responses.Failure[dtos.ApplicationResponseDto](http.StatusInternalServerError, "Failed to submit application: "+err.Error())
	}

	return responses.Success(http.StatusCreated, responseDto)
}

func (s *applicationService) GetMyApplications(ctx context.Context, userId, page, limit int) responses.ResultResponse[responses.PaginatedResponse[dtos.ApplicationResponseDto]] {
	candidate, err := s.candidateRepo.GetByUserID(ctx, userId)
	if err != nil {
		return responses.PaginatedSuccess(http.StatusOK, []dtos.ApplicationResponseDto{}, 0, page, limit)
	}

	applications, count, err := s.applicationRepo.GetByCandidateID(ctx, candidate.Id, page, limit)
	if err != nil {
		return responses.Failure[responses.PaginatedResponse[dtos.ApplicationResponseDto]](http.StatusInternalServerError, "Failed to fetch applications")
	}

	dtosList := make([]dtos.ApplicationResponseDto, len(applications))
	for i, app := range applications {
		dtosList[i] = s.toDto(app)
	}

	return responses.PaginatedSuccess(http.StatusOK, dtosList, count, page, limit)
}

func (s *applicationService) GetJobApplications(ctx context.Context, userId, jobId, page, limit int) responses.ResultResponse[responses.PaginatedResponse[dtos.ApplicationResponseDto]] {
	recruiter, err := s.recruiterRepo.GetByUserID(ctx, userId)
	if err != nil {
		return responses.Failure[responses.PaginatedResponse[dtos.ApplicationResponseDto]](http.StatusForbidden, "You are not a recruiter")
	}

	job, err := s.jobRepo.GetByID(ctx, jobId)
	if err != nil {
		return responses.Failure[responses.PaginatedResponse[dtos.ApplicationResponseDto]](http.StatusNotFound, "Job not found")
	}

	if job.CompanyId != recruiter.CompanyId {
		return responses.Failure[responses.PaginatedResponse[dtos.ApplicationResponseDto]](http.StatusForbidden, "You do not have permission to view applications for this job")
	}

	applications, count, err := s.applicationRepo.GetByJobID(ctx, jobId, page, limit)
	if err != nil {
		return responses.Failure[responses.PaginatedResponse[dtos.ApplicationResponseDto]](http.StatusInternalServerError, "Failed to fetch applications")
	}

	dtosList := make([]dtos.ApplicationResponseDto, len(applications))
	for i, app := range applications {
		dtosList[i] = s.toDto(app)
	}

	return responses.PaginatedSuccess(http.StatusOK, dtosList, count, page, limit)
}

func (s *applicationService) UpdateApplicationStatus(ctx context.Context, userId int, applicationId int, dto dtos.ApplicationUpdateStatusRequestDto) responses.ResultResponse[dtos.ApplicationResponseDto] {
	recruiter, err := s.recruiterRepo.GetByUserID(ctx, userId)
	if err != nil {
		return responses.Failure[dtos.ApplicationResponseDto](http.StatusForbidden, "You are not a recruiter")
	}

	app, err := s.applicationRepo.GetByID(ctx, applicationId)
	if err != nil {
		return responses.Failure[dtos.ApplicationResponseDto](http.StatusNotFound, "Application not found")
	}

	job, err := s.jobRepo.GetByID(ctx, app.JobId)
	if err != nil {
		return responses.Failure[dtos.ApplicationResponseDto](http.StatusNotFound, "Associated job not found")
	}

	if job.CompanyId != recruiter.CompanyId {
		return responses.Failure[dtos.ApplicationResponseDto](http.StatusForbidden, "You do not have permission to update this application")
	}

	app, err = s.applicationRepo.UpdateStatus(ctx, applicationId, dto.Status)
	if err != nil {
		return responses.Failure[dtos.ApplicationResponseDto](http.StatusInternalServerError, "Failed to update application status")
	}

	return responses.Success(http.StatusOK, s.toDto(app))
}

func (s *applicationService) toDto(app models.Application) dtos.ApplicationResponseDto {
	dto := dtos.ApplicationResponseDto{
		Id:          app.Id,
		CandidateId: app.CandidateId,
		JobId:       app.JobId,
		Status:      app.Status,
		AppliedAt:   app.AppliedAt,
	}

	if app.Job != nil {
		dto.Job = app.Job.ToJobResponseDto()
	}

	if app.Candidate.Id != 0 {
		candidateDto := app.Candidate.ToCandidateResponseDto()
		dto.Candidate = &candidateDto
	}

	return dto
}
