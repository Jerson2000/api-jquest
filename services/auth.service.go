package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jerson2000/jquest/config"
	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/models"
	"github.com/jerson2000/jquest/repositories"
	"github.com/jerson2000/jquest/responses"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(dto dtos.AuthLoginRequestDto) responses.ResultResponse[dtos.AuthResponseDto]
	Signup(dto dtos.AuthSignupRequestDto) responses.ResultResponse[dtos.AuthResponseDto]
	Refresh(dto dtos.AuthRefreshRequestDto) responses.ResultResponse[dtos.AuthResponseDto]
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService() AuthService {
	repo := repositories.NewUserRepository(config.Database)

	return &authService{
		userRepo: repo,
	}
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

	token, refreshToken, err := a.generateTokens(isUserExist)

	if err != nil {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return responses.Success(int(http.StatusOK), dtos.AuthResponseDto{
		Token:        token,
		RefreshToken: refreshToken,
	})
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

	token, refreshToken, err := a.generateTokens(createdUser)

	return responses.Success(int(http.StatusOK), dtos.AuthResponseDto{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

func (a authService) Refresh(dto dtos.AuthRefreshRequestDto) responses.ResultResponse[dtos.AuthResponseDto] {
	token, err := jwt.ParseWithClaims(dto.RefreshToken, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWTKey, nil
	})

	if err != nil || !token.Valid {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusUnauthorized,
			"invalid refresh token",
		)
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusUnauthorized,
			"invalid token claims",
		)
	}

	// verify
	key := fmt.Sprintf("user:%d:refresh", claims.Id)
	var cachedToken string
	err = config.CacheStore.Get(key, &cachedToken)

	if err != nil || cachedToken != dto.RefreshToken {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusUnauthorized,
			"refresh token expired or invalid",
		)
	}

	user, err := a.userRepo.GetByID(claims.Id)
	if err != nil {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusUnauthorized,
			"user not found",
		)
	}

	newToken, newRefreshToken, err := a.generateTokens(user)
	if err != nil {
		return responses.Failure[dtos.AuthResponseDto](
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return responses.Success(int(http.StatusOK), dtos.AuthResponseDto{
		Token:        newToken,
		RefreshToken: newRefreshToken,
	})
}

func (a *authService) generateTokens(user models.User) (string, string, error) {
	// 15 minutes
	accessExpiry := time.Now().Add(15 * time.Minute).UTC()
	accessToken, err := createToken(user, accessExpiry)
	if err != nil {
		return "", "", err
	}

	// 24 hours
	refreshExpiry := time.Now().Add(24 * time.Hour).UTC()
	refreshToken, err := createToken(user, refreshExpiry)
	if err != nil {
		return "", "", err
	}

	// cached
	key := fmt.Sprintf("user:%d:refresh", user.Id)
	err = config.CacheStore.Set(key, refreshToken, 25*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func createToken(user models.User, expiry time.Time) (string, error) {
	issueAt := time.Now().UTC()
	claims := models.JWTClaims{
		Id:   user.Id,
		Name: user.Name,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
			IssuedAt:  jwt.NewNumericDate(issueAt),
			Issuer:    "jquest",
			ID:        uuid.NewString(),
		},
	}
	signed := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return signed.SignedString(config.JWTKey)
}
