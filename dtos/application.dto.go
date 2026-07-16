package dtos

import (
	"time"
)

type ApplicationCreateRequestDto struct {
	JobId       int    `json:"jobId" binding:"required"`
	CoverLetter string `json:"coverLetter"`
	ResumeURL   string `json:"resumeUrl"`
}

type ApplicationUpdateStatusRequestDto struct {
	Status string `json:"status" binding:"required,oneof=pending reviewing interviewig offered rejected accepted"`
}

type ApplicationResponseDto struct {
	Id          int                   `json:"id"`
	CandidateId int                   `json:"candidateId"`
	JobId       int                   `json:"jobId"`
	Status      string                `json:"status"`
	AppliedAt   time.Time             `json:"appliedAt"`
	CoverLetter string                `json:"coverLetter,omitempty"`
	ResumeURL   string                `json:"resumeUrl,omitempty"`
	Job         JobResponseDto        `json:"job,omitempty"`
	Candidate   *CandidateResponseDto `json:"candidate,omitempty"`
}
