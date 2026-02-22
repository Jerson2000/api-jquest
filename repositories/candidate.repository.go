package repositories

import (
	"context"

	"github.com/jerson2000/jquest/models"
	"gorm.io/gorm"
)

type CandidateRepository interface {
	Create(ctx context.Context, candidate models.Candidate) (models.Candidate, error)
	GetByUserID(ctx context.Context, userId int) (models.Candidate, error)
}

type candidateRepo struct {
	db *gorm.DB
}

func NewCandidateRepository(db *gorm.DB) CandidateRepository {
	return &candidateRepo{db}
}

func (r *candidateRepo) Create(ctx context.Context, candidate models.Candidate) (models.Candidate, error) {
	err := r.db.WithContext(ctx).Create(&candidate).Error
	return candidate, err
}

func (r *candidateRepo) GetByUserID(ctx context.Context, userId int) (models.Candidate, error) {
	var candidate models.Candidate
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).First(&candidate).Error
	return candidate, err
}
