package repositories

import (
	"context"

	"github.com/jerson2000/jquest/models"
	"gorm.io/gorm"
)

type RecruiterRepository interface {
	Create(context.Context, models.Recruiter) (models.Recruiter, error)
	GetByUserID(context.Context, int) (models.Recruiter, error)
	GetByCompanyID(context.Context, int) ([]models.Recruiter, error)
	Delete(context.Context, int) error
	GetRecruiters(context.Context) ([]models.Recruiter, error)
}

type recruiterRepo struct {
	db *gorm.DB
}

func NewRecruiterRepository(db *gorm.DB) RecruiterRepository {
	return &recruiterRepo{db}
}

func (r *recruiterRepo) Create(ctx context.Context, recruiter models.Recruiter) (models.Recruiter, error) {
	if err := r.db.WithContext(ctx).Create(&recruiter).Error; err != nil {
		return models.Recruiter{}, err
	}
	return recruiter, nil
}

func (r *recruiterRepo) GetByUserID(ctx context.Context, userId int) (models.Recruiter, error) {
	var recruiter models.Recruiter
	err := r.db.WithContext(ctx).Preload("Company").First(&recruiter, "user_id = ?", userId).Error
	return recruiter, err
}

func (r *recruiterRepo) GetByCompanyID(ctx context.Context, companyId int) ([]models.Recruiter, error) {
	var recruiters []models.Recruiter
	err := r.db.WithContext(ctx).Find(&recruiters, "company_id = ?", companyId).Error
	return recruiters, err
}

func (r *recruiterRepo) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Recruiter{}, "id = ?", id).Error
}

func (r *recruiterRepo) GetRecruiters(ctx context.Context) ([]models.Recruiter, error) {
	var recruiters []models.Recruiter
	err := r.db.WithContext(ctx).Preload("Company").Find(&recruiters).Error
	return recruiters, err
}
