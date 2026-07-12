package services

import (
	"context"
	"net/http"

	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/repositories"
	"github.com/jerson2000/jquest/responses"
)

type SkillService interface {
	AddSkillToCandidate(ctx context.Context, userId int, skillName string) responses.ResultResponse[any]
	RemoveSkillFromCandidate(ctx context.Context, userId int, skillName string) responses.ResultResponse[any]
	AddSkillToJob(ctx context.Context, recruiterUserId, jobId int, skillName string) responses.ResultResponse[any]
	RemoveSkillFromJob(ctx context.Context, recruiterUserId, jobId int, skillName string) responses.ResultResponse[any]
	GetSkills(ctx context.Context) responses.ResultResponse[[]dtos.SkillResponseDto]
}

type skillService struct {
	skillRepo     repositories.SkillRepository
	candidateRepo repositories.CandidateRepository
	jobRepo       repositories.JobRepository
	recruiterRepo repositories.RecruiterRepository
}

func NewSkillService(
	skillRepo repositories.SkillRepository,
	candidateRepo repositories.CandidateRepository,
	jobRepo repositories.JobRepository,
	recruiterRepo repositories.RecruiterRepository,
) SkillService {
	return &skillService{
		skillRepo:     skillRepo,
		candidateRepo: candidateRepo,
		jobRepo:       jobRepo,
		recruiterRepo: recruiterRepo,
	}
}

func (s *skillService) AddSkillToCandidate(ctx context.Context, userId int, skillName string) responses.ResultResponse[any] {
	candidate, err := s.candidateRepo.GetByUserID(ctx, userId)
	if err != nil {
		return responses.Failure[any](http.StatusNotFound, "Candidate profile not found")
	}

	skill, err := s.skillRepo.GetOrCreate(ctx, skillName)
	if err != nil {
		return responses.Failure[any](http.StatusInternalServerError, "Failed to get/create skill")
	}

	err = s.skillRepo.AddSkillToCandidate(ctx, candidate.Id, skill.Id)
	if err != nil {
		return responses.Failure[any](http.StatusInternalServerError, "Failed to link skill to candidate")
	}

	return responses.Success[any](http.StatusOK, "Skill added to candidate profile")
}

func (s *skillService) RemoveSkillFromCandidate(ctx context.Context, userId int, skillName string) responses.ResultResponse[any] {
	candidate, err := s.candidateRepo.GetByUserID(ctx, userId)
	if err != nil {
		return responses.Failure[any](http.StatusNotFound, "Candidate profile not found")
	}

	skill, err := s.skillRepo.GetOrCreate(ctx, skillName)
	if err != nil {
		return responses.Failure[any](http.StatusNotFound, "Skill not found")
	}

	err = s.skillRepo.RemoveSkillFromCandidate(ctx, candidate.Id, skill.Id)
	if err != nil {
		return responses.Failure[any](http.StatusInternalServerError, "Failed to remove skill from candidate")
	}

	return responses.Success[any](http.StatusOK, "Skill removed from candidate profile")
}

func (s *skillService) AddSkillToJob(ctx context.Context, recruiterUserId, jobId int, skillName string) responses.ResultResponse[any] {
	recruiter, err := s.recruiterRepo.GetByUserID(ctx, recruiterUserId)
	if err != nil {
		return responses.Failure[any](http.StatusForbidden, "Only recruiters can update job skills")
	}

	job, err := s.jobRepo.GetByID(ctx, jobId)
	if err != nil {
		return responses.Failure[any](http.StatusNotFound, "Job not found")
	}

	if job.CompanyId != recruiter.CompanyId {
		return responses.Failure[any](http.StatusForbidden, "You do not own this job posting")
	}

	skill, err := s.skillRepo.GetOrCreate(ctx, skillName)
	if err != nil {
		return responses.Failure[any](http.StatusInternalServerError, "Failed to get/create skill")
	}

	err = s.skillRepo.AddSkillToJob(ctx, job.Id, skill.Id)
	if err != nil {
		return responses.Failure[any](http.StatusInternalServerError, "Failed to add skill to job")
	}

	return responses.Success[any](http.StatusOK, "Skill added to job post")
}

func (s *skillService) RemoveSkillFromJob(ctx context.Context, recruiterUserId, jobId int, skillName string) responses.ResultResponse[any] {
	recruiter, err := s.recruiterRepo.GetByUserID(ctx, recruiterUserId)
	if err != nil {
		return responses.Failure[any](http.StatusForbidden, "Only recruiters can update job skills")
	}

	job, err := s.jobRepo.GetByID(ctx, jobId)
	if err != nil {
		return responses.Failure[any](http.StatusNotFound, "Job not found")
	}

	if job.CompanyId != recruiter.CompanyId {
		return responses.Failure[any](http.StatusForbidden, "You do not own this job posting")
	}

	skill, err := s.skillRepo.GetOrCreate(ctx, skillName)
	if err != nil {
		return responses.Failure[any](http.StatusNotFound, "Skill not found")
	}

	err = s.skillRepo.RemoveSkillFromJob(ctx, job.Id, skill.Id)
	if err != nil {
		return responses.Failure[any](http.StatusInternalServerError, "Failed to remove skill from job")
	}

	return responses.Success[any](http.StatusOK, "Skill removed from job post")
}

func (s *skillService) GetSkills(ctx context.Context) responses.ResultResponse[[]dtos.SkillResponseDto] {
	skills, err := s.skillRepo.GetAll(ctx)
	if err != nil {
		return responses.Failure[[]dtos.SkillResponseDto](http.StatusInternalServerError, "Failed to retrieve skills")
	}

	dtosList := make([]dtos.SkillResponseDto, len(skills))
	for i, sk := range skills {
		dtosList[i] = dtos.SkillResponseDto{
			Id:   sk.Id,
			Name: sk.Name,
		}
	}

	return responses.Success(http.StatusOK, dtosList)
}
