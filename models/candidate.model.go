package models

import (
	"time"

	"github.com/jerson2000/jquest/dtos"
)

type Candidate struct {
	Id              int        `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName       string     `gorm:"size:100;not null" json:"firstName"`
	LastName        string     `gorm:"size:100;not null" json:"lastName"`
	Email           string     `gorm:"size:150;uniqueIndex;not null" json:"email"`
	Phone           string     `gorm:"size:30" json:"phone"`
	LinkedInURL     string     `gorm:"size:255" json:"linkedinUrl"`
	ResumeURL       string     `gorm:"size:255" json:"resumeUrl"`
	TotalExperience float32    `json:"totalExperience"`
	CurrentTitle    string     `gorm:"size:150" json:"currentTitle"`
	CurrentLocation string     `gorm:"size:150" json:"currentLocation"`
	UserId          int        `gorm:"not null" json:"userId"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	DeletedAt       *time.Time `gorm:"index" json:"deletedAt,omitempty"`

	User User `gorm:"foreignKey:UserId" json:"user"`

	Experiences  []Experience  `gorm:"foreignKey:CandidateId" json:"experiences"`
	Applications []Application `gorm:"foreignKey:CandidateId" json:"applications"`
}

func (c *Candidate) ToCandidateResponseDto() dtos.CandidateResponseDto {
	return dtos.CandidateResponseDto{
		Id:              c.Id,
		FirstName:       c.FirstName,
		LastName:        c.LastName,
		Email:           c.Email,
		Phone:           c.Phone,
		LinkedInURL:     c.LinkedInURL,
		ResumeURL:       c.ResumeURL,
		TotalExperience: c.TotalExperience,
		CurrentTitle:    c.CurrentTitle,
		CurrentLocation: c.CurrentLocation,
	}
}
