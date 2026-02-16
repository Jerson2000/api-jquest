package repositories

import (
	"context"

	"github.com/jerson2000/jquest/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(context.Context) ([]models.User, error)
	GetByID(context.Context, int) (models.User, error)
	GetByEmail(context.Context, string) (models.User, error)
	Create(context.Context, models.User) (models.User, error)
	Update(context.Context, int, models.User) (models.User, error)
	Delete(context.Context, int) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) GetByID(ctx context.Context, id int) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error
	return user, err
}

func (r *userRepo) Create(ctx context.Context, user models.User) (models.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepo) Update(ctx context.Context, id int, user models.User) (models.User, error) {
	var existing models.User
	if err := r.db.WithContext(ctx).First(&existing, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}

	existing.Name = user.Name
	existing.Email = user.Email
	existing.Role = user.Role
	// existing.Password // Password update usually separate or handled if provided? Assuming dto handles mapping
	existing.Phone = user.Phone
	existing.Sex = user.Sex
	existing.IsActive = user.IsActive
	existing.IsVerified = user.IsVerified

	if err := r.db.WithContext(ctx).Save(&existing).Error; err != nil {
		return models.User{}, err
	}
	return existing, nil
}

func (r *userRepo) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id).Error
}
