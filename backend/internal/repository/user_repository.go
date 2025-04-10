package repository

import (
	"github.com/shironxn/inkarya/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindAll() ([]domain.User, error)
	FindByID(id uint) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *domain.User) error {
	return r.DB.Updates(&user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.User{}, id).Error
}
