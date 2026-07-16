package models

import (
	"time"

	"github.com/jerson2000/jquest/dtos"
	"gorm.io/gorm"
)

type Education struct {
	Id           int            `gorm:"primaryKey;autoIncrement" json:"id"`
	CandidateId  int            `gorm:"index;not null" json:"candidateId"`
	SchoolName   string         `gorm:"size:150;not null" json:"schoolName"`
	Degree       string         `gorm:"size:100;not null" json:"degree"`
	FieldOfStudy string         `gorm:"size:100;not null" json:"fieldOfStudy"`
	StartDate    time.Time      `json:"startDate"`
	EndDate      *time.Time     `json:"endDate,omitempty"`
	IsCurrent    bool           `json:"isCurrent"`
	Description  string         `gorm:"type:text" json:"description"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

func (e *Education) ToEducationResponseDto() dtos.EducationResponseDto {
	return dtos.EducationResponseDto{
		Id:           e.Id,
		CandidateId:  e.CandidateId,
		SchoolName:   e.SchoolName,
		Degree:       e.Degree,
		FieldOfStudy: e.FieldOfStudy,
		StartDate:    e.StartDate,
		EndDate:      e.EndDate,
		IsCurrent:    e.IsCurrent,
		Description:  e.Description,
	}
}

func ToEducationResponseDtoList(educations []Education) []dtos.EducationResponseDto {
	dtoList := make([]dtos.EducationResponseDto, len(educations))
	for i, e := range educations {
		dtoList[i] = e.ToEducationResponseDto()
	}
	return dtoList
}
