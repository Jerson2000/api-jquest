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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CompanyService interface {
	CreateCompany(ctx context.Context, userId int, dto dtos.CompanyCreateRequestDto) responses.ResultResponse[dtos.CompanyResponseDto]
	GetCompanies(ctx context.Context) responses.ResultResponse[[]dtos.CompanyResponseDto]
	ApplyAsRecruiter(ctx context.Context, companyDto dtos.CompanyApplyRequestDto) responses.ResultResponse[dtos.CompanyResponseDto]
}

type companyService struct {
	companyRepo   repositories.CompanyRepository
	recruiterRepo repositories.RecruiterRepository
	userRepo      repositories.UserRepository
}

func NewCompanyService() CompanyService {
	return &companyService{
		companyRepo:   repositories.NewCompanyRepository(config.Database),
		recruiterRepo: repositories.NewRecruiterRepository(config.Database),
		userRepo:      repositories.NewUserRepository(config.Database),
	}
}

func (s *companyService) CreateCompany(ctx context.Context, userId int, dto dtos.CompanyCreateRequestDto) responses.ResultResponse[dtos.CompanyResponseDto] {
	company := models.Company{
		Name:        dto.Name,
		Industry:    dto.Industry,
		Website:     *dto.Website,
		Location:    *dto.Location,
		CompanySize: *dto.CompanySize,
	}

	newCompany, err := s.companyRepo.Create(ctx, company)
	if err != nil {
		return responses.Failure[dtos.CompanyResponseDto](http.StatusInternalServerError, "Failed to create company")
	}

	return responses.Success(http.StatusCreated, newCompany.ToCompanyResponseDto())
}

func (s *companyService) GetCompanies(ctx context.Context) responses.ResultResponse[[]dtos.CompanyResponseDto] {
	companies, err := s.companyRepo.GetAll(ctx)
	if err != nil {
		return responses.Failure[[]dtos.CompanyResponseDto](http.StatusInternalServerError, "Failed to fetch companies")
	}

	return responses.Success(http.StatusOK, models.ToCompanyResponseDtoList(companies))
}

func (s *companyService) GetCompanyByID(ctx context.Context, id int) responses.ResultResponse[dtos.CompanyResponseDto] {
	company, err := s.companyRepo.GetByID(ctx, id)
	if err != nil {
		return responses.Failure[dtos.CompanyResponseDto](http.StatusNotFound, "Company not found")
	}
	return responses.Success(http.StatusOK, company.ToCompanyResponseDto())
}

func (s *companyService) UpdateCompany(ctx context.Context, id int, dto dtos.CompanyCreateRequestDto) responses.ResultResponse[dtos.CompanyResponseDto] {
	company, err := s.companyRepo.GetByID(ctx, id)
	if err != nil {
		return responses.Failure[dtos.CompanyResponseDto](http.StatusNotFound, "Company not found")
	}

	company.Name = dto.Name
	company.Industry = dto.Industry
	company.Website = *dto.Website
	company.Location = *dto.Location
	company.CompanySize = *dto.CompanySize

	updatedCompany, err := s.companyRepo.Update(ctx, id, company)
	if err != nil {
		return responses.Failure[dtos.CompanyResponseDto](http.StatusInternalServerError, "Failed to update company")
	}

	return responses.Success(http.StatusOK, updatedCompany.ToCompanyResponseDto())
}

func (s *companyService) DeleteCompany(ctx context.Context, id int) responses.ResultResponse[dtos.CompanyResponseDto] {
	err := s.companyRepo.Delete(ctx, id)
	if err != nil {
		return responses.Failure[dtos.CompanyResponseDto](http.StatusInternalServerError, "Failed to delete company")
	}
	return responses.Success(http.StatusNoContent, dtos.CompanyResponseDto{})
}

func (s *companyService) ApplyAsRecruiter(ctx context.Context, companyDto dtos.CompanyApplyRequestDto) responses.ResultResponse[dtos.CompanyResponseDto] {

	newUser := models.User{
		Name:     companyDto.User.Name,
		Email:    companyDto.User.Email,
		Password: companyDto.User.Password,
		Sex:      companyDto.User.Sex,
		Phone:    companyDto.User.Phone,
		Role:     enums.RECRUITER,
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(companyDto.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.Failure[dtos.CompanyResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	newUser.Password = string(hashed)
	user, err := s.userRepo.Create(ctx, newUser)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.Failure[dtos.CompanyResponseDto](
				http.StatusBadRequest,
				"user not found",
			)
		}
		return responses.Failure[dtos.CompanyResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	newCompany := models.Company{
		Name:        companyDto.Company.Name,
		Industry:    companyDto.Company.Industry,
		Website:     *companyDto.Company.Website,
		Location:    *companyDto.Company.Location,
		CompanySize: *companyDto.Company.CompanySize,
	}
	company, err := s.companyRepo.Create(ctx, newCompany)

	if err != nil {
		return responses.Failure[dtos.CompanyResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	recruiter := models.Recruiter{
		UserId:     user.Id,
		CompanyId:  company.Id,
		Position:   "Recruiter",
		IsVerified: false,
	}
	_, err = s.recruiterRepo.Create(ctx, recruiter)

	if err != nil {
		return responses.Failure[dtos.CompanyResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return responses.Success(http.StatusCreated, company.ToCompanyResponseDto())
}
