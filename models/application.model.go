package models

import "time"

type Application struct {
	Id          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	CandidateId int        `gorm:"index:idx_application_candidate_job,unique;not null" json:"candidateId"`
	JobId       int        `gorm:"index:idx_application_candidate_job,unique;not null" json:"jobId"`
	Status      string     `gorm:"size:50;index;not null" json:"status"`
	AppliedAt   time.Time  `json:"appliedAt"`
	Candidate   Candidate  `gorm:"foreignKey:CandidateId"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `gorm:"index" json:"deletedAt,omitempty"`

	Job *Job `gorm:"foreignKey:JobId" json:"applications,omitempty"`
}
