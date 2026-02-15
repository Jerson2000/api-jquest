package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/enums"
	"github.com/jerson2000/jquest/models"
	"github.com/jerson2000/jquest/repositories"
	"github.com/jerson2000/jquest/responses"
	"gorm.io/gorm"
)

type RecruiterService interface {
	CreateRecruiter(ctx context.Context, userId int, dto dtos.RecruiterCreateRequestDto) responses.ResultResponse[dtos.RecruiterResponseDto]
	GetRecruiters(ctx context.Context) responses.ResultResponse[[]dtos.RecruiterResponseDto]
	GetByUserID(ctx context.Context, userId int) responses.ResultResponse[dtos.RecruiterResponseDto]
	GetByCompanyID(ctx context.Context, companyId int) responses.ResultResponse[[]dtos.RecruiterResponseDto]
}

type recruiterService struct {
	recruiterRepo repositories.RecruiterRepository
	companyRepo   repositories.CompanyRepository
	userRepo      repositories.UserRepository
}

func NewRecruiterService() RecruiterService {
	recruiterRepo := repositories.NewRecruiterRepository(config.Database)
	companyRepo := repositories.NewCompanyRepository(config.Database)
	userRepo := repositories.NewUserRepository(config.Database)
	return &recruiterService{recruiterRepo, companyRepo, userRepo}
}

func (s *recruiterService) CreateRecruiter(ctx context.Context, userId int, dto dtos.RecruiterCreateRequestDto) responses.ResultResponse[dtos.RecruiterResponseDto] {
	// 1. Check if user is already a recruiter?
	_, err := s.recruiterRepo.GetByUserID(ctx, userId)
	if err == nil {
		return responses.Failure[dtos.RecruiterResponseDto](http.StatusBadRequest, "User is already a recruiter")
	}

	// 2. Check if company exists
	_, err = s.companyRepo.GetByID(ctx, dto.CompanyId)
	if err != nil {
		return responses.Failure[dtos.RecruiterResponseDto](http.StatusNotFound, "Company not found")
	}

	// 3. Create Recruiter
	recruiter := models.Recruiter{
		UserId:     userId,
		CompanyId:  dto.CompanyId,
		Position:   dto.Position,
		IsVerified: true,
	}

	newRecruiter, err := s.recruiterRepo.Create(ctx, recruiter)
	if err != nil {
		return responses.Failure[dtos.RecruiterResponseDto](http.StatusInternalServerError, "Failed to create recruiter profile")
	}

	// 4. Update User Role to RECRUITER
	user, err := s.userRepo.GetByID(ctx, userId)
	if err == nil {
		user.Role = enums.RECRUITER
		s.userRepo.Update(ctx, userId, user)
	}

	return responses.Success(http.StatusCreated, newRecruiter.ToRecruiterResponseDto())
}

func (s *recruiterService) GetRecruiters(ctx context.Context) responses.ResultResponse[[]dtos.RecruiterResponseDto] {
	recruiters, err := s.recruiterRepo.GetRecruiters(ctx)
	if err != nil {
		return responses.Failure[[]dtos.RecruiterResponseDto](http.StatusInternalServerError, "Failed to fetch recruiters")
	}
	return responses.Success(http.StatusOK, models.ToRecruiterResponseDtoList(recruiters))
}

func (s *recruiterService) GetByUserID(ctx context.Context, userId int) responses.ResultResponse[dtos.RecruiterResponseDto] {
	recruiter, err := s.recruiterRepo.GetByUserID(ctx, userId)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.Failure[dtos.RecruiterResponseDto](
				http.StatusBadRequest,
				"user not found",
			)
		}
		return responses.Failure[dtos.RecruiterResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return responses.Success(int(http.StatusOK), recruiter.ToRecruiterResponseDto())

}

func (s *recruiterService) GetByCompanyID(ctx context.Context, companyId int) responses.ResultResponse[[]dtos.RecruiterResponseDto] {
	recruiters, err := s.recruiterRepo.GetByCompanyID(ctx, companyId)
	if err != nil {
		return responses.Failure[[]dtos.RecruiterResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return responses.Success(int(http.StatusOK), models.ToRecruiterResponseDtoList(recruiters))
}
