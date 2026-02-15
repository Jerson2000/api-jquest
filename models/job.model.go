package models

import (
	"time"

	"github.com/jerson2000/jquest/dtos"
	"github.com/jerson2000/jquest/enums"
)

type Job struct {
	Id           int                `gorm:"primaryKey;autoIncrement" json:"id"`
	CompanyId    int                `gorm:"index;not null" json:"companyId"`
	Title        string             `gorm:"size:200;not null" json:"title"`
	Description  string             `gorm:"type:text;not null" json:"description"`
	Location     string             `gorm:"size:200" json:"location"`
	JobType      enums.JobType      `gorm:"size:50;index" json:"jobType"`
	Experience   int                `json:"experience"`
	SalaryMin    *int               `json:"salaryMin,omitempty"`
	SalaryMax    *int               `json:"salaryMax,omitempty"`
	Currency     string             `gorm:"size:3" json:"currency"` // e.g., "USD", "PHP"
	SalaryPeriod enums.SalaryPeriod `gorm:"size:20" json:"salaryPeriod"`
	Status       enums.JobStatus    `gorm:"size:50;index" json:"status"`
	Deadline     time.Time          `gorm:"index" json:"deadline"`
	PublishedAt  *time.Time         `gorm:"index" json:"publishedAt,omitempty"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
	DeletedAt    *time.Time         `gorm:"index" json:"deletedAt,omitempty"`

	Company      Company       `gorm:"foreignKey:CompanyId" json:"company"`
	Applications []Application `gorm:"foreignKey:JobId" json:"applications"`
}

func (j *Job) ToJobResponseDto() dtos.JobResponseDto {
	return dtos.JobResponseDto{
		Id:          j.Id,
		CompanyId:   j.CompanyId,
		Title:       j.Title,
		Description: j.Description,
		Location:    j.Location,
		JobType:     j.JobType,
		Experience:  j.Experience,
		SalaryMin:   j.SalaryMin,
		SalaryMax:   j.SalaryMax,
		Status:      j.Status,
		Deadline:    j.Deadline,
		Company:     j.Company.ToCompanyResponseDto(),
		CreatedAt:   j.CreatedAt,
	}
}

func ToJobResponseDtoList(jobs []Job) []dtos.JobResponseDto {
	dtoList := make([]dtos.JobResponseDto, len(jobs))
	for i, j := range jobs {
		dtoList[i] = j.ToJobResponseDto()
	}
	return dtoList
}
