package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/models"
	"github.com/jerson2000/jquest/repositories"
	"github.com/jerson2000/jquest/responses"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers(ctx context.Context) responses.ResultResponse[[]dtos.UserResponseDto]
	GetUser(ctx context.Context, id int) responses.ResultResponse[dtos.UserResponseDto]
	CreateUser(ctx context.Context, user dtos.UserCreateRequestDto) responses.ResultResponse[dtos.UserResponseDto]
	UpdateUser(ctx context.Context, id int, user dtos.UserUpdateRequestDto) responses.ResultResponse[dtos.UserResponseDto]
	DeleteUser(ctx context.Context, id int) responses.ResultResponse[string]
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService() UserService {
	repo := repositories.NewUserRepository(config.Database)
	return &userService{repo}
}

func (s *userService) GetUsers(ctx context.Context) responses.ResultResponse[[]dtos.UserResponseDto] {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return responses.Failure[[]dtos.UserResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	dtoList := models.ToUserResponseDtoList(users)

	return responses.Success(int(http.StatusOK), dtoList)
}

func (s *userService) GetUser(ctx context.Context, id int) responses.ResultResponse[dtos.UserResponseDto] {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.Failure[dtos.UserResponseDto](
				http.StatusBadRequest,
				"user not found",
			)
		}
		return responses.Failure[dtos.UserResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	return responses.Success(int(http.StatusOK), user.ToUserResponseDto())
}

func (s *userService) CreateUser(ctx context.Context, user dtos.UserCreateRequestDto) responses.ResultResponse[dtos.UserResponseDto] {

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Sex:      user.Sex,
		Phone:    user.Phone,
	}

	if existing, _ := s.repo.GetByEmail(ctx, user.Email); existing.Id != 0 {
		return responses.Failure[dtos.UserResponseDto](
			http.StatusBadRequest,
			"email already in use",
		)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.Failure[dtos.UserResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	newUser.Password = string(hashed)

	createdUser, err := s.repo.Create(ctx, newUser)
	if err != nil {
		return responses.Failure[dtos.UserResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return responses.Success(int(http.StatusOK), createdUser.ToUserResponseDto())
}

func (s *userService) UpdateUser(ctx context.Context, id int, user dtos.UserUpdateRequestDto) responses.ResultResponse[dtos.UserResponseDto] {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.Failure[dtos.UserResponseDto](
			http.StatusNotFound,
			"user not found",
		)
	}

	if user.Email != "" && user.Email != existing.Email {
		if other, _ := s.repo.GetByEmail(ctx, user.Email); other.Id != 0 && other.Id != id {
			return responses.Failure[dtos.UserResponseDto](
				http.StatusBadRequest,
				"email already in use",
			)
		}
		existing.Email = user.Email
	}

	if user.Name != "" {
		existing.Name = user.Name
	}
	existing.Phone = user.Phone
	existing.Sex = user.Sex

	updated, err := s.repo.Update(ctx, id, existing)
	if err != nil {
		return responses.Failure[dtos.UserResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return responses.Success(int(http.StatusOK), updated.ToUserResponseDto())
}

func (s *userService) DeleteUser(ctx context.Context, id int) responses.ResultResponse[string] {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return responses.Failure[string](
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	return responses.Success(int(http.StatusNoContent), "")
}
