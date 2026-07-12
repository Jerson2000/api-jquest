package services

import (
	"context"
	"net/http"
	"time"

	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/models"
	"github.com/jerson2000/jquest/repositories"
	"github.com/jerson2000/jquest/responses"
)

type SavedJobService interface {
	SaveJob(ctx context.Context, userId int, dto dtos.SavedJobCreateRequestDto) responses.ResultResponse[any]
	UnsaveJob(ctx context.Context, userId int, jobId int) responses.ResultResponse[any]
	GetSavedJobs(ctx context.Context, userId int, page, limit int) responses.ResultResponse[responses.PaginatedResponse[dtos.SavedJobResponseDto]]
}

type savedJobService struct {
	savedJobRepo  repositories.SavedJobRepository
	candidateRepo repositories.CandidateRepository
	jobRepo       repositories.JobRepository
}

func NewSavedJobService(
	savedJobRepo repositories.SavedJobRepository,
	candidateRepo repositories.CandidateRepository,
	jobRepo repositories.JobRepository,
) SavedJobService {
	return &savedJobService{
		savedJobRepo:  savedJobRepo,
		candidateRepo: candidateRepo,
		jobRepo:       jobRepo,
	}
}

func (s *savedJobService) SaveJob(ctx context.Context, userId int, dto dtos.SavedJobCreateRequestDto) responses.ResultResponse[any] {
	candidate, err := s.candidateRepo.GetByUserID(ctx, userId)
	if err != nil {
		return responses.Failure[any](http.StatusNotFound, "Candidate profile not found")
	}

	_, err = s.jobRepo.GetByID(ctx, dto.JobId)
	if err != nil {
		return responses.Failure[any](http.StatusNotFound, "Job not found")
	}

	alreadySaved, _ := s.savedJobRepo.IsSaved(ctx, candidate.Id, dto.JobId)
	if alreadySaved {
		return responses.Failure[any](http.StatusBadRequest, "Job already saved")
	}

	savedJob := models.SavedJob{
		CandidateId: candidate.Id,
		JobId:       dto.JobId,
		SavedAt:     time.Now().UTC(),
	}

	_, err = s.savedJobRepo.Create(ctx, savedJob)
	if err != nil {
		return responses.Failure[any](http.StatusInternalServerError, "Failed to save job")
	}

	return responses.Success[any](http.StatusCreated, "Job saved successfully")
}

func (s *savedJobService) UnsaveJob(ctx context.Context, userId int, jobId int) responses.ResultResponse[any] {
	candidate, err := s.candidateRepo.GetByUserID(ctx, userId)
	if err != nil {
		return responses.Failure[any](http.StatusNotFound, "Candidate profile not found")
	}

	alreadySaved, _ := s.savedJobRepo.IsSaved(ctx, candidate.Id, jobId)
	if !alreadySaved {
		return responses.Failure[any](http.StatusBadRequest, "Job not saved")
	}

	err = s.savedJobRepo.Delete(ctx, candidate.Id, jobId)
	if err != nil {
		return responses.Failure[any](http.StatusInternalServerError, "Failed to unsave job")
	}

	return responses.Success[any](http.StatusOK, "Job unsaved successfully")
}

func (s *savedJobService) GetSavedJobs(ctx context.Context, userId int, page, limit int) responses.ResultResponse[responses.PaginatedResponse[dtos.SavedJobResponseDto]] {
	candidate, err := s.candidateRepo.GetByUserID(ctx, userId)
	if err != nil {
		return responses.Failure[responses.PaginatedResponse[dtos.SavedJobResponseDto]](http.StatusNotFound, "Candidate profile not found")
	}

	savedJobs, total, err := s.savedJobRepo.GetByCandidateID(ctx, candidate.Id, page, limit)
	if err != nil {
		return responses.Failure[responses.PaginatedResponse[dtos.SavedJobResponseDto]](http.StatusInternalServerError, "Failed to retrieve saved jobs")
	}

	dtosList := make([]dtos.SavedJobResponseDto, len(savedJobs))
	for i, sj := range savedJobs {
		dtosList[i] = dtos.SavedJobResponseDto{
			Id:          sj.Id,
			CandidateId: sj.CandidateId,
			JobId:       sj.JobId,
			SavedAt:     sj.SavedAt,
			Job:         sj.Job.ToJobResponseDto(),
		}
	}

	return responses.PaginatedSuccess(http.StatusOK, dtosList, total, page, limit)
}

