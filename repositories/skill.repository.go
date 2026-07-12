package repositories

import (
	"context"

	"github.com/jerson2000/jquest/models"
	"gorm.io/gorm"
)

type SkillRepository interface {
	Create(ctx context.Context, skill models.Skill) (models.Skill, error)
	GetOrCreate(ctx context.Context, name string) (models.Skill, error)
	GetByID(ctx context.Context, id int) (models.Skill, error)
	GetAll(ctx context.Context) ([]models.Skill, error)
	AddSkillToCandidate(ctx context.Context, candidateId, skillId int) error
	RemoveSkillFromCandidate(ctx context.Context, candidateId, skillId int) error
	AddSkillToJob(ctx context.Context, jobId, skillId int) error
	RemoveSkillFromJob(ctx context.Context, jobId, skillId int) error
}

type skillRepo struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepo{db}
}

func (r *skillRepo) Create(ctx context.Context, skill models.Skill) (models.Skill, error) {
	err := r.db.WithContext(ctx).Create(&skill).Error
	return skill, err
}

func (r *skillRepo) GetOrCreate(ctx context.Context, name string) (models.Skill, error) {
	var skill models.Skill
	err := r.db.WithContext(ctx).Where("name = ?", name).FirstOrCreate(&skill, models.Skill{Name: name}).Error
	return skill, err
}

func (r *skillRepo) GetByID(ctx context.Context, id int) (models.Skill, error) {
	var skill models.Skill
	err := r.db.WithContext(ctx).First(&skill, id).Error
	return skill, err
}

func (r *skillRepo) GetAll(ctx context.Context) ([]models.Skill, error) {
	var skills []models.Skill
	err := r.db.WithContext(ctx).Find(&skills).Error
	return skills, err
}

func (r *skillRepo) AddSkillToCandidate(ctx context.Context, candidateId, skillId int) error {
	return r.db.WithContext(ctx).Exec("INSERT INTO candidate_skills (candidate_id, skill_id) VALUES (?, ?) ON CONFLICT DO NOTHING", candidateId, skillId).Error
}

func (r *skillRepo) RemoveSkillFromCandidate(ctx context.Context, candidateId, skillId int) error {
	return r.db.WithContext(ctx).Exec("DELETE FROM candidate_skills WHERE candidate_id = ? AND skill_id = ?", candidateId, skillId).Error
}

func (r *skillRepo) AddSkillToJob(ctx context.Context, jobId, skillId int) error {
	return r.db.WithContext(ctx).Exec("INSERT INTO job_skills (job_id, skill_id) VALUES (?, ?) ON CONFLICT DO NOTHING", jobId, skillId).Error
}

func (r *skillRepo) RemoveSkillFromJob(ctx context.Context, jobId, skillId int) error {
	return r.db.WithContext(ctx).Exec("DELETE FROM job_skills WHERE job_id = ? AND skill_id = ?", jobId, skillId).Error
}
