package repositories

import (
	"context"

	"github.com/jerson2000/jquest/models"
	"gorm.io/gorm"
)

type SavedJobRepository interface {
	Create(ctx context.Context, savedJob models.SavedJob) (models.SavedJob, error)
	Delete(ctx context.Context, candidateId, jobId int) error
	GetByCandidateID(ctx context.Context, candidateId int, page, limit int) ([]models.SavedJob, int64, error)
	IsSaved(ctx context.Context, candidateId, jobId int) (bool, error)
}

type savedJobRepo struct {
	db *gorm.DB
}

func NewSavedJobRepository(db *gorm.DB) SavedJobRepository {
	return &savedJobRepo{db}
}

func (r *savedJobRepo) Create(ctx context.Context, savedJob models.SavedJob) (models.SavedJob, error) {
	err := r.db.WithContext(ctx).Create(&savedJob).Error
	return savedJob, err
}

func (r *savedJobRepo) Delete(ctx context.Context, candidateId, jobId int) error {
	return r.db.WithContext(ctx).Where("candidate_id = ? AND job_id = ?", candidateId, jobId).Delete(&models.SavedJob{}).Error
}

func (r *savedJobRepo) GetByCandidateID(ctx context.Context, candidateId int, page, limit int) ([]models.SavedJob, int64, error) {
	var savedJobs []models.SavedJob
	var count int64

	db := r.db.WithContext(ctx).Model(&models.SavedJob{}).Where("candidate_id = ?", candidateId)
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Preload("Job").Preload("Job.Company").Offset(offset).Limit(limit).Find(&savedJobs).Error
	return savedJobs, count, err
}

func (r *savedJobRepo) IsSaved(ctx context.Context, candidateId, jobId int) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.SavedJob{}).Where("candidate_id = ? AND job_id = ?", candidateId, jobId).Count(&count).Error
	return count > 0, err
}
