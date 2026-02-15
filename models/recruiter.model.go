package models

import (
	"time"

	"github.com/jerson2000/jquest/dtos"
)

type Recruiter struct {
	Id         int        `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId     int        `gorm:"not null" json:"userId"`
	CompanyId  int        `gorm:"not null" json:"companyId"`
	Position   string     `gorm:"size:100;not null" json:"position"`
	IsVerified bool       `gorm:"default:false" json:"isVerified"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `gorm:"index" json:"deletedAt,omitempty"`

	User    *User    `gorm:"foreignKey:UserId" json:"user,omitempty"`
	Company *Company `gorm:"foreignKey:CompanyId" json:"company,omitempty"`
}

func (r *Recruiter) ToRecruiterResponseDto() dtos.RecruiterResponseDto {
	dto := dtos.RecruiterResponseDto{
		Id:         r.Id,
		UserId:     r.UserId,
		CompanyId:  r.CompanyId,
		Position:   r.Position,
		IsVerified: r.IsVerified,
	}

	if r.Company != nil {
		companyDto := r.Company.ToCompanyResponseDto()
		dto.Company = &companyDto
	}

	return dto
}

func ToRecruiterResponseDtoList(recruiters []Recruiter) []dtos.RecruiterResponseDto {
	dtosList := make([]dtos.RecruiterResponseDto, 0, len(recruiters))

	for _, recruiter := range recruiters {
		dtosList = append(dtosList, recruiter.ToRecruiterResponseDto())
	}

	return dtosList
}
