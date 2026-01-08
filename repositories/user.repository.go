package repositories

import (
	"github.com/jerson2000/jquest/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(int) (models.User, error)
	GetByEmail(email string) (models.User, error)
	Create(models.User) (models.User, error)
	Update(int, models.User) (models.User, error)
	Delete(int) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) GetByID(id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepo) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}

func (r *userRepo) Create(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepo) Update(id int, user models.User) (models.User, error) {
	var existing models.User
	if err := r.db.First(&existing, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}

	existing.Name = user.Name
	existing.Email = user.Email
	existing.Role = user.Role
	existing.Phone = user.Phone
	existing.Sex = user.Sex
	existing.IsActive = user.IsActive
	existing.IsVerified = user.IsVerified

	if err := r.db.Save(&existing).Error; err != nil {
		return models.User{}, err
	}
	return existing, nil
}

func (r *userRepo) Delete(id int) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
