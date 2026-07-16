package dtos

import "time"

type EducationCreateRequestDto struct {
	SchoolName   string     `json:"schoolName" binding:"required"`
	Degree       string     `json:"degree" binding:"required"`
	FieldOfStudy string     `json:"fieldOfStudy" binding:"required"`
	StartDate    time.Time  `json:"startDate" binding:"required"`
	EndDate      *time.Time `json:"endDate"`
	IsCurrent    bool       `json:"isCurrent"`
	Description  string     `json:"description"`
}

type EducationResponseDto struct {
	Id           int        `json:"id"`
	CandidateId  int        `json:"candidateId"`
	SchoolName   string     `json:"schoolName"`
	Degree       string     `json:"degree"`
	FieldOfStudy string     `json:"fieldOfStudy"`
	StartDate    time.Time  `json:"startDate"`
	EndDate      *time.Time `json:"endDate,omitempty"`
	IsCurrent    bool       `json:"isCurrent"`
	Description  string     `json:"description"`
}
