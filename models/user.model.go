package models

import (
	"time"

	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/enums"
)

type User struct {
	Id         int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string     `gorm:"size:100;not null" json:"name"`
	Email      string     `gorm:"size:150;uniqueIndex;not null" json:"email"`
	Role       enums.Role `gorm:"type:varchar(20);not null;default:candidate" json:"role"`
	Password   string     `gorm:"not null" json:"password,omitempty"`
	Phone      string     `gorm:"size:20;default:null" json:"phone,omitempty"`
	Sex        string     `gorm:"default:null" json:"gender,omitempty"`
	IsActive   bool       `gorm:"default:true" json:"isActive"`
	IsVerified bool       `gorm:"default:false" json:"isVerified"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"deletedAt,omitempty"`

	Candidate *Candidate `gorm:"foreignKey:UserId" json:"candidate,omitempty"`
}

func (u *User) ToUserResponseDto() dtos.UserResponseDto {
	return dtos.UserResponseDto{
		Id:         u.Id,
		Name:       u.Name,
		Email:      u.Email,
		Role:       u.Role,
		Phone:      u.Phone,
		Sex:        u.Sex,
		IsActive:   u.IsActive,
		IsVerified: u.IsVerified,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func ToUserResponseDtoList(users []User) []dtos.UserResponseDto {
	dtoList := make([]dtos.UserResponseDto, len(users))
	for i, u := range users {
		dtoList[i] = u.ToUserResponseDto()
	}
	return dtoList
}
