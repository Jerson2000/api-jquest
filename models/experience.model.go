package models

import "time"

type Experience struct {
	Id          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	CandidateId int        `gorm:"index;not null" json:"candidateId"`
	CompanyName string     `gorm:"size:150;not null" json:"companyName"`
	JobTitle    string     `gorm:"size:150;not null" json:"jobTitle"`
	Description string     `gorm:"type:text" json:"description"`
	StartDate   time.Time  `json:"startDate"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	IsCurrent   bool       `json:"isCurrent"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}
