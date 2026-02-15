package repositories

import (
	"context"

	"github.com/jerson2000/jquest/models"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(context.Context, models.Company) (models.Company, error)
	GetByID(context.Context, int) (models.Company, error)
	GetAll(context.Context) ([]models.Company, error)
	Update(context.Context, int, models.Company) (models.Company, error)
	Delete(context.Context, int) error
}

type companyRepo struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepo{db}
}

func (r *companyRepo) Create(ctx context.Context, company models.Company) (models.Company, error) {
	if err := r.db.WithContext(ctx).Create(&company).Error; err != nil {
		return models.Company{}, err
	}
	return company, nil
}

func (r *companyRepo) GetByID(ctx context.Context, id int) (models.Company, error) {
	var company models.Company
	err := r.db.WithContext(ctx).First(&company, "id = ?", id).Error
	return company, err
}

func (r *companyRepo) GetAll(ctx context.Context) ([]models.Company, error) {
	var companies []models.Company
	err := r.db.WithContext(ctx).Find(&companies).Error
	return companies, err
}

func (r *companyRepo) Update(ctx context.Context, id int, company models.Company) (models.Company, error) {
	var existing models.Company
	if err := r.db.WithContext(ctx).First(&existing, "id = ?", id).Error; err != nil {
		return models.Company{}, err
	}

	existing.Name = company.Name
	existing.Industry = company.Industry
	existing.Website = company.Website
	existing.Location = company.Location
	existing.CompanySize = company.CompanySize

	if err := r.db.WithContext(ctx).Save(&existing).Error; err != nil {
		return models.Company{}, err
	}
	return existing, nil
}

func (r *companyRepo) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Company{}, "id = ?", id).Error
}
