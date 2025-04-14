package repository

import (
	"github.com/google/uuid"
	"github.com/shironxn/inkarya/internal/domain"
	"gorm.io/gorm"
)

// User Repository Interface
type UserRepository interface {
	// User
	CreateUser(user *domain.User) error
	FindAllUsers() ([]domain.User, error)
	FindUserByID(id uuid.UUID) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id uuid.UUID) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

// User Implementation
func (r *userRepository) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepository) FindAllUsers() ([]domain.User, error) {
	var users []domain.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindUserByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *domain.User) error {
	return r.DB.Updates(&user).Error
}

func (r *userRepository) DeleteUser(id uuid.UUID) error {
	return r.DB.Unscoped().Delete(&domain.User{}, id).Error
}
