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
	if err := r.DB.Preload("Skills").Preload("Disabilities").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindUserByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	if err := r.DB.Preload("Skills").Preload("Disabilities").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *domain.User) error {
	// Start a transaction
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Clear existing relationships
	if err := tx.Model(user).Association("Skills").Clear(); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(user).Association("Disabilities").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	// Update user and its relationships
	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *userRepository) DeleteUser(id uuid.UUID) error {
	return r.DB.Unscoped().Delete(&domain.User{}, id).Error
}
