package models

import (
	"time"

	"github.com/jerson2000/jquest/dtos"
	"gorm.io/gorm"
)

type Candidate struct {
	Id              int            `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName       string         `gorm:"size:100;not null" json:"firstName"`
	LastName        string         `gorm:"size:100;not null" json:"lastName"`
	Phone           string         `gorm:"size:30" json:"phone"`
	LinkedInURL     string         `gorm:"size:255" json:"linkedinUrl"`
	ResumeURL       string         `gorm:"size:255" json:"resumeUrl"`
	TotalExperience float32        `json:"totalExperience"`
	CurrentTitle    string         `gorm:"size:150" json:"currentTitle"`
	CurrentLocation string         `gorm:"size:150" json:"currentLocation"`
	Bio             string         `gorm:"type:text" json:"bio,omitempty"`
	ExpectedSalary  *int           `json:"expectedSalary,omitempty"`
	UserId          int            `gorm:"uniqueIndex;not null" json:"userId"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	User User `gorm:"foreignKey:UserId" json:"user"`

	Experiences  []Experience  `gorm:"foreignKey:CandidateId" json:"experiences"`
	Educations   []Education   `gorm:"foreignKey:CandidateId" json:"educations"`
	Applications []Application `gorm:"foreignKey:CandidateId" json:"applications"`
	Skills       []Skill       `gorm:"many2many:candidate_skills;" json:"skills,omitempty"`
}

func (c *Candidate) ToCandidateResponseDto() dtos.CandidateResponseDto {
	email := ""
	if c.User.Email != "" {
		email = c.User.Email
	}
	
	skills := make([]string, len(c.Skills))
	for i, s := range c.Skills {
		skills[i] = s.Name
	}

	educationsList := ToEducationResponseDtoList(c.Educations)

	return dtos.CandidateResponseDto{
		Id:              c.Id,
		FirstName:       c.FirstName,
		LastName:        c.LastName,
		Email:           email,
		Phone:           c.Phone,
		LinkedInURL:     c.LinkedInURL,
		ResumeURL:       c.ResumeURL,
		TotalExperience: c.TotalExperience,
		CurrentTitle:    c.CurrentTitle,
		CurrentLocation: c.CurrentLocation,
		Bio:             c.Bio,
		ExpectedSalary:  c.ExpectedSalary,
		Skills:          skills,
		Educations:      educationsList,
	}
}

