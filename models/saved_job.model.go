package models

import (
	"time"
)

type SavedJob struct {
	Id          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CandidateId int       `gorm:"uniqueIndex:idx_candidate_job;not null" json:"candidateId"`
	JobId       int       `gorm:"uniqueIndex:idx_candidate_job;not null" json:"jobId"`
	SavedAt     time.Time `gorm:"not null" json:"savedAt"`

	Candidate Candidate `gorm:"foreignKey:CandidateId" json:"candidate,omitempty"`
	Job       Job       `gorm:"foreignKey:JobId" json:"job,omitempty"`
}
