package models

import (
	"time"

	"github.com/jerson2000/jquest/dtos"
)

type Company struct {
	Id          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"size:200;uniqueIndex;not null" json:"name"`
	Industry    string     `gorm:"size:150" json:"industry"`
	Website     string     `gorm:"size:255" json:"website"`
	Location    string     `gorm:"size:200" json:"location"`
	CompanySize string     `gorm:"size:50" json:"companySize"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `gorm:"index" json:"deletedAt,omitempty"`

	Jobs []Job `gorm:"foreignKey:CompanyId" json:"jobs"`
}

func (c *Company) ToCompanyResponseDto() dtos.CompanyResponseDto {
	return dtos.CompanyResponseDto{
		Id:          c.Id,
		Name:        c.Name,
		Industry:    c.Industry,
		Website:     c.Website,
		Location:    c.Location,
		CompanySize: c.CompanySize,
	}
}

func ToCompanyResponseDtoList(companies []Company) []dtos.CompanyResponseDto {
	dtoList := make([]dtos.CompanyResponseDto, len(companies))
	for i, c := range companies {
		dtoList[i] = c.ToCompanyResponseDto()
	}
	return dtoList
}
