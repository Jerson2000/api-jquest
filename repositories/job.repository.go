package repositories

import (
	"context"

	"github.com/jerson2000/jquest/models"
	"gorm.io/gorm"
)

type JobRepository interface {
	Create(context.Context, models.Job) (models.Job, error)
	GetByID(context.Context, int) (models.Job, error)
	GetAll(context.Context, int, int) ([]models.Job, int64, error)
	GetByCompanyID(context.Context, int) ([]models.Job, error)
	Update(context.Context, int, models.Job) (models.Job, error)
	Delete(context.Context, int) error
}

type jobRepo struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepo{db}
}

func (r *jobRepo) Create(ctx context.Context, job models.Job) (models.Job, error) {
	if err := r.db.WithContext(ctx).Create(&job).Error; err != nil {
		return models.Job{}, err
	}
	return job, nil
}

func (r *jobRepo) GetByID(ctx context.Context, id int) (models.Job, error) {
	var job models.Job

	err := r.db.WithContext(ctx).Preload("Company").First(&job, "id = ?", id).Error
	return job, err
}

func (r *jobRepo) GetAll(ctx context.Context, page, limit int) ([]models.Job, int64, error) {
	var jobs []models.Job
	var count int64

	if err := r.db.WithContext(ctx).Model(&models.Job{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).Preload("Company").Offset(offset).Limit(limit).Find(&jobs).Error
	return jobs, count, err
}

func (r *jobRepo) GetByCompanyID(ctx context.Context, companyId int) ([]models.Job, error) {
	var jobs []models.Job
	err := r.db.WithContext(ctx).Find(&jobs, "company_id = ?", companyId).Error
	return jobs, err
}

func (r *jobRepo) Update(ctx context.Context, id int, job models.Job) (models.Job, error) {
	var existing models.Job
	if err := r.db.WithContext(ctx).First(&existing, "id = ?", id).Error; err != nil {
		return models.Job{}, err
	}

	existing.Title = job.Title
	existing.Description = job.Description
	existing.Location = job.Location
	existing.JobType = job.JobType
	existing.Experience = job.Experience
	existing.SalaryMin = job.SalaryMin
	existing.SalaryMax = job.SalaryMax
	existing.Status = job.Status
	existing.Deadline = job.Deadline

	if err := r.db.WithContext(ctx).Save(&existing).Error; err != nil {
		return models.Job{}, err
	}
	return existing, nil
}

func (r *jobRepo) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Job{}, "id = ?", id).Error
}
