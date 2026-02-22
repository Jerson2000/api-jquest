package repositories

import (
	"context"

	"github.com/jerson2000/jquest/models"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Create(ctx context.Context, application models.Application) (models.Application, error)
	GetByCandidateAndJob(ctx context.Context, candidateId, jobId int) (models.Application, error)
	GetByID(ctx context.Context, id int) (models.Application, error)
	GetByJobID(ctx context.Context, jobId, page, limit int) ([]models.Application, int64, error)
	GetByCandidateID(ctx context.Context, candidateId, page, limit int) ([]models.Application, int64, error)
	UpdateStatus(ctx context.Context, id int, status string) (models.Application, error)
	WithTx(tx *gorm.DB) ApplicationRepository
}

type applicationRepo struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &applicationRepo{db}
}

func (r *applicationRepo) WithTx(tx *gorm.DB) ApplicationRepository {
	return &applicationRepo{db: tx}
}

func (r *applicationRepo) Create(ctx context.Context, application models.Application) (models.Application, error) {
	err := r.db.WithContext(ctx).Create(&application).Error
	return application, err
}

func (r *applicationRepo) GetByCandidateAndJob(ctx context.Context, candidateId, jobId int) (models.Application, error) {
	var application models.Application
	err := r.db.WithContext(ctx).Where("candidate_id = ? AND job_id = ?", candidateId, jobId).First(&application).Error
	return application, err
}

func (r *applicationRepo) GetByID(ctx context.Context, id int) (models.Application, error) {
	var application models.Application
	err := r.db.WithContext(ctx).Preload("Candidate").First(&application, id).Error
	return application, err
}

func (r *applicationRepo) GetByJobID(ctx context.Context, jobId, page, limit int) ([]models.Application, int64, error) {
	var applications []models.Application
	var count int64

	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).Model(&models.Application{}).Where("job_id = ?", jobId).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.WithContext(ctx).Preload("Candidate").Where("job_id = ?", jobId).Offset(offset).Limit(limit).Find(&applications).Error
	return applications, count, err
}

func (r *applicationRepo) GetByCandidateID(ctx context.Context, candidateId, page, limit int) ([]models.Application, int64, error) {
	var applications []models.Application
	var count int64

	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).Model(&models.Application{}).Where("candidate_id = ?", candidateId).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.WithContext(ctx).Preload("Job").Preload("Job.Company").Where("candidate_id = ?", candidateId).Offset(offset).Limit(limit).Find(&applications).Error
	return applications, count, err
}

func (r *applicationRepo) UpdateStatus(ctx context.Context, id int, status string) (models.Application, error) {
	var application models.Application
	if err := r.db.WithContext(ctx).First(&application, id).Error; err != nil {
		return application, err
	}
	application.Status = status
	err := r.db.WithContext(ctx).Save(&application).Error
	return application, err
}
