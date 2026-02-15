package services

import (
	"context"
	"net/http"
	"time"

	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/enums"
	"github.com/jerson2000/jquest/models"
	"github.com/jerson2000/jquest/repositories"
	"github.com/jerson2000/jquest/responses"
)

type JobService interface {
	CreateJobPost(ctx context.Context, userId int, dto dtos.JobCreateJobRequestDto) responses.ResultResponse[dtos.JobResponseDto]
	GetJobPost(ctx context.Context, page, limit int) responses.ResultResponse[responses.PaginatedResponse[dtos.JobResponseDto]]
	GetJobPostById(ctx context.Context, id int) responses.ResultResponse[dtos.JobResponseDto]
	UpdateJobPost(ctx context.Context, userId int, id int, dto dtos.JobCreateJobRequestDto) responses.ResultResponse[dtos.JobResponseDto]
	DeleteJobPost(ctx context.Context, userId int, id int) responses.ResultResponse[dtos.JobResponseDto]
}

type jobService struct {
	jobRepo       repositories.JobRepository
	recruiterRepo repositories.RecruiterRepository
}

func NewJobService() JobService {
	jobRepo := repositories.NewJobRepository(config.Database)
	recruiterRepo := repositories.NewRecruiterRepository(config.Database)
	return &jobService{jobRepo, recruiterRepo}
}

func (s *jobService) CreateJobPost(ctx context.Context, userId int, dto dtos.JobCreateJobRequestDto) responses.ResultResponse[dtos.JobResponseDto] {
	// 1. Verify user is recruiter for the company
	recruiter, err := s.recruiterRepo.GetByUserID(ctx, userId)
	if err != nil {
		return responses.Failure[dtos.JobResponseDto](http.StatusForbidden, "You are not a registered recruiter")
	}

	if recruiter.CompanyId != dto.CompanyId {
		return responses.Failure[dtos.JobResponseDto](http.StatusForbidden, "You cannot post jobs for this company")
	}

	job := models.Job{
		CompanyId:   dto.CompanyId,
		Title:       dto.Title,
		Description: dto.Description,
		Location:    dto.Location,
		JobType:     dto.JobType,
		Experience:  dto.Experience,
		SalaryMin:   dto.SalaryMin,
		SalaryMax:   dto.SalaryMax,
		Status:      enums.JobStatusDraft,
		Deadline:    dto.Deadline,
	}

	if dto.Status != "" {
		job.Status = enums.JobStatus(dto.Status)
	}
	if dto.PublishedAt != nil && dto.Status == enums.JobStatusOpen {
		job.PublishedAt = dto.PublishedAt
	}
	if dto.PublishedAt == nil && dto.Status == enums.JobStatusOpen {
		now := time.Now().UTC()
		job.PublishedAt = &now
	}

	newJob, err := s.jobRepo.Create(ctx, job)
	if err != nil {
		return responses.Failure[dtos.JobResponseDto](http.StatusInternalServerError, "Failed to create job post")
	}

	fullJob, err := s.jobRepo.GetByID(ctx, newJob.Id)
	if err != nil {
		return responses.Failure[dtos.JobResponseDto](http.StatusInternalServerError, "Failed to fetch created job details")
	}

	return responses.Success(http.StatusCreated, fullJob.ToJobResponseDto())
}

func (s *jobService) GetJobPost(ctx context.Context, page, limit int) responses.ResultResponse[responses.PaginatedResponse[dtos.JobResponseDto]] {
	jobs, count, err := s.jobRepo.GetAll(ctx, page, limit)
	if err != nil {
		return responses.Failure[responses.PaginatedResponse[dtos.JobResponseDto]](http.StatusInternalServerError, "Failed to fetch jobs")
	}
	return responses.PaginatedSuccess(http.StatusOK, models.ToJobResponseDtoList(jobs), count, page, limit)
}

func (s *jobService) GetJobPostById(ctx context.Context, id int) responses.ResultResponse[dtos.JobResponseDto] {
	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return responses.Failure[dtos.JobResponseDto](http.StatusNotFound, "Job not found")
	}
	return responses.Success(http.StatusOK, job.ToJobResponseDto())
}

func (s *jobService) UpdateJobPost(ctx context.Context, userId int, id int, dto dtos.JobCreateJobRequestDto) responses.ResultResponse[dtos.JobResponseDto] {
	recruiter, err := s.recruiterRepo.GetByUserID(ctx, userId)
	if err != nil {
		return responses.Failure[dtos.JobResponseDto](http.StatusInternalServerError, err.Error())
	}

	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return responses.Failure[dtos.JobResponseDto](http.StatusNotFound, "Job not found")
	}

	if recruiter.CompanyId != job.CompanyId {
		return responses.Failure[dtos.JobResponseDto](http.StatusForbidden, "You don't have permission to access this resource")
	}

	job.Title = dto.Title
	job.Description = dto.Description
	job.Location = dto.Location
	job.JobType = dto.JobType
	job.Experience = dto.Experience
	job.SalaryMin = dto.SalaryMin
	job.SalaryMax = dto.SalaryMax
	job.Status = dto.Status
	job.Deadline = dto.Deadline
	job.PublishedAt = dto.PublishedAt

	updatedJob, err := s.jobRepo.Update(ctx, id, job)
	if err != nil {
		return responses.Failure[dtos.JobResponseDto](http.StatusInternalServerError, "Failed to update job post")
	}

	return responses.Success(http.StatusOK, updatedJob.ToJobResponseDto())
}

func (s *jobService) DeleteJobPost(ctx context.Context, userId int, id int) responses.ResultResponse[dtos.JobResponseDto] {
	recruiter, err := s.recruiterRepo.GetByUserID(ctx, userId)

	if err != nil {
		return responses.Failure[dtos.JobResponseDto](http.StatusInternalServerError, err.Error())
	}

	job, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return responses.Failure[dtos.JobResponseDto](http.StatusNotFound, "Job not found")
	}

	if recruiter.CompanyId != job.CompanyId {
		return responses.Failure[dtos.JobResponseDto](http.StatusForbidden, "You don't have permission to access this resource!")
	}

	err = s.jobRepo.Delete(ctx, id)
	if err != nil {
		return responses.Failure[dtos.JobResponseDto](http.StatusInternalServerError, "Failed to delete job post")
	}
	return responses.Success(http.StatusNoContent, dtos.JobResponseDto{})
}
