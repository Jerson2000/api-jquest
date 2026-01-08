package dtos

import (
	"time"

	"github.com/jerson2000/jquest/internal/enums"
)

type UserCreateRequestDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	Phone    string `json:"phone,omitempty"`
	Sex      string `json:"gender,omitempty" binding:"required"`
}

type UserUpdateRequestDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone,omitempty"`
	Sex   string `json:"gender,omitempty"`
}

type UserResponseDto struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Role       enums.Role `json:"role"`
	Phone      string     `json:"phone,omitempty"`
	Sex        string     `json:"gender,omitempty"`
	IsActive   bool       `json:"isActive"`
	IsVerified bool       `json:"isVerified"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
