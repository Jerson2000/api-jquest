package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jerson2000/jquest/internal/config"
	"github.com/jerson2000/jquest/internal/dtos"
	"github.com/jerson2000/jquest/internal/models"
	"github.com/jerson2000/jquest/internal/repositories"
	"github.com/jerson2000/jquest/internal/responses"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(dto dtos.AuthLoginRequestDto) responses.ResultResponse[dtos.AuthResponseDto]
	Signup(dto dtos.AuthSignupRequestDto) responses.ResultResponse[dtos.AuthResponseDto]
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService() AuthService {
	repo := repositories.NewUserRepository(config.Database)
	return &authService{userRepo: repo}
}

func (a authService) Login(dto dtos.AuthLoginRequestDto) responses.ResultResponse[dtos.AuthResponseDto] {
	isUserExist, err := a.userRepo.GetByEmail(dto.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.Failure[dtos.AuthResponseDto](
				http.StatusBadRequest,
				"incorrect email or password",
			)
		}
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	err = bcrypt.CompareHashAndPassword([]byte(isUserExist.Password), []byte(dto.Password))

	if err != nil {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusBadRequest,
			"incorrect email or password",
		)
	}

	token, err := generateToken(isUserExist)

	if err != nil {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return responses.Success(int(http.StatusOK), dtos.AuthResponseDto{Token: token})
}

func (a authService) Signup(dto dtos.AuthSignupRequestDto) responses.ResultResponse[dtos.AuthResponseDto] {

	newUser := models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	if existing, _ := a.userRepo.GetByEmail(dto.Email); existing.Id != 0 {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusBadRequest,
			"email already in use",
		)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	newUser.Password = string(hashed)
	createdUser, err := a.userRepo.Create(newUser)
	if err != nil {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	token, err := generateToken(createdUser)

	return responses.Success(int(http.StatusOK), dtos.AuthResponseDto{Token: token})
}

func generateToken(user models.User) (string, error) {
	expiry := time.Now().Add(1 * time.Hour).UTC()
	issueAt := time.Now().UTC()

	claims := models.JWTClaims{
		Id:   user.Id,
		Name: user.Name,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
			IssuedAt:  jwt.NewNumericDate(issueAt),
			Issuer:    "jquest",
		},
	}
	signed := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := signed.SignedString(config.JWTKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
