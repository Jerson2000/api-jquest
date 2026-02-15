package dtos

import (
	"time"

	"github.com/jerson2000/jquest/enums"
)

type JobCreateJobRequestDto struct {
	CompanyId   int             `json:"companyId" binding:"required"`
	Title       string          `json:"title" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Location    string          `json:"location" binding:"required"`
	JobType     enums.JobType   `json:"jobType" binding:"required,oneof=full-time part-time contract internship"`
	Experience  int             `json:"experience" binding:"required,numeric"`
	SalaryMin   *int            `json:"salaryMin"`
	SalaryMax   *int            `json:"salaryMax"`
	Status      enums.JobStatus `json:"status" binding:"omitempty,oneof=draft open close"`
	Deadline    time.Time       `json:"deadline"`
	PublishedAt *time.Time      `json:"publishedAt,omitempty"`
}

type JobResponseDto struct {
	Id          int                `json:"id"`
	CompanyId   int                `json:"companyId"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Location    string             `json:"location"`
	JobType     enums.JobType      `json:"jobType"`
	Experience  int                `json:"experience"`
	SalaryMin   *int               `json:"salaryMin,omitempty"`
	SalaryMax   *int               `json:"salaryMax,omitempty"`
	Status      enums.JobStatus    `json:"status"`
	Deadline    time.Time          `json:"deadline"`
	PublishedAt *time.Time         `json:"publishedAt,omitempty"`
	Company     CompanyResponseDto `json:"company"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	DeletedAt   *time.Time         `json:"deletedAt,omitempty"`
}
