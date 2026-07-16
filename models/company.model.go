package models

import (
	"time"

	"github.com/jerson2000/jquest/dtos"
	"gorm.io/gorm"
)

type Company struct {
	Id          int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"size:200;uniqueIndex;not null" json:"name"`
	Industry    string         `gorm:"size:150" json:"industry"`
	Website     string         `gorm:"size:255" json:"website"`
	Location    string         `gorm:"size:200" json:"location"`
	CompanySize string         `gorm:"size:50" json:"companySize"`
	Description string         `gorm:"type:text" json:"description,omitempty"`
	LogoURL     string         `gorm:"size:255" json:"logoUrl,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

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
		Description: c.Description,
		LogoURL:     c.LogoURL,
	}
}

func ToCompanyResponseDtoList(companies []Company) []dtos.CompanyResponseDto {
	dtoList := make([]dtos.CompanyResponseDto, len(companies))
	for i, c := range companies {
		dtoList[i] = c.ToCompanyResponseDto()
	}
	return dtoList
}
