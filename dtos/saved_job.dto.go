package dtos

import (
	"time"
)

type SavedJobCreateRequestDto struct {
	JobId int `json:"jobId" binding:"required"`
}

type SavedJobResponseDto struct {
	Id          int            `json:"id"`
	CandidateId int            `json:"candidateId"`
	JobId       int            `json:"jobId"`
	SavedAt     time.Time      `json:"savedAt"`
	Job         JobResponseDto `json:"job,omitempty"`
}
