package models

import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	Id          int            `gorm:"primaryKey;autoIncrement" json:"id"`
	CandidateId int            `gorm:"index:idx_application_candidate_job,unique;not null" json:"candidateId"`
	JobId       int            `gorm:"index:idx_application_candidate_job,unique;not null" json:"jobId"`
	Status      string         `gorm:"size:50;index;not null" json:"status"`
	AppliedAt   time.Time      `json:"appliedAt"`
	CoverLetter string         `gorm:"type:text" json:"coverLetter,omitempty"`
	ResumeURL   string         `gorm:"size:255" json:"resumeUrl,omitempty"`
	Candidate   Candidate      `gorm:"foreignKey:CandidateId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Job *Job `gorm:"foreignKey:JobId" json:"applications,omitempty"`
}
