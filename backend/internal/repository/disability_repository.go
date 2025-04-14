package repository

import (
	"github.com/shironxn/inkarya/internal/domain"
	"gorm.io/gorm"
)

type DisabilityRepository interface {
	FindAll() ([]domain.Disability, error)
	FindByID(id uint) (domain.Disability, error)
}

type disabilityRepository struct {
	db *gorm.DB
}

func NewDisabilityRepository(db *gorm.DB) DisabilityRepository {
	return &disabilityRepository{db: db}
}

func (r *disabilityRepository) FindAll() ([]domain.Disability, error) {
	var disabilities []domain.Disability
	if err := r.db.Find(&disabilities).Error; err != nil {
		return nil, err
	}
	return disabilities, nil
}

func (r *disabilityRepository) FindByID(id uint) (domain.Disability, error) {
	var disability domain.Disability
	if err := r.db.First(&disability, id).Error; err != nil {
		return domain.Disability{}, err
	}
	return disability, nil
}
