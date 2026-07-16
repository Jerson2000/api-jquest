package services

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

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

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUsers(ctx context.Context) responses.ResultResponse[[]dtos.UserResponseDto] {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		slog.Error("failed to get all users", "error", err)
		return responses.Failure[[]dtos.UserResponseDto](
			http.StatusInternalServerError,
			"failed to retrieve users",
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
		slog.Error("failed to get user by ID", "id", id, "error", err)
		return responses.Failure[dtos.UserResponseDto](
			http.StatusInternalServerError,
			"failed to retrieve user",
		)
	}
	return responses.Success(int(http.StatusOK), user.ToUserResponseDto())
}

func (s *userService) CreateUser(ctx context.Context, user dtos.UserCreateRequestDto) responses.ResultResponse[dtos.UserResponseDto] {

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Gender:   user.Gender,
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
		slog.Error("failed to hash password during user creation", "error", err)
		return responses.Failure[dtos.UserResponseDto](
			http.StatusInternalServerError,
			"failed to create user",
		)
	}
	newUser.Password = string(hashed)

	createdUser, err := s.repo.Create(ctx, newUser)
	if err != nil {
		slog.Error("failed to create user record", "error", err)
		return responses.Failure[dtos.UserResponseDto](
			http.StatusInternalServerError,
			"failed to create user",
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
	existing.Gender = user.Gender

	updated, err := s.repo.Update(ctx, id, existing)
	if err != nil {
		slog.Error("failed to update user record", "id", id, "error", err)
		return responses.Failure[dtos.UserResponseDto](
			http.StatusInternalServerError,
			"failed to update user",
		)
	}

	return responses.Success(int(http.StatusOK), updated.ToUserResponseDto())
}

func (s *userService) DeleteUser(ctx context.Context, id int) responses.ResultResponse[string] {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		slog.Error("failed to delete user record", "id", id, "error", err)
		return responses.Failure[string](
			http.StatusInternalServerError,
			"failed to delete user",
		)
	}
	return responses.Success(int(http.StatusNoContent), "")
}
