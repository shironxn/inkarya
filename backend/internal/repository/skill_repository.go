package repository

import (
	"github.com/shironxn/inkarya/internal/domain"
	"gorm.io/gorm"
)

type SkillRepository interface {
	FindAll() ([]domain.Skill, error)
	FindByID(id uint) (domain.Skill, error)
}

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{
		db: db,
	}
}

func (r *skillRepository) FindAll() ([]domain.Skill, error) {
	var skills []domain.Skill
	if err := r.db.Find(&skills).Error; err != nil {
		return nil, err
	}
	return skills, nil
}

func (r *skillRepository) FindByID(id uint) (domain.Skill, error) {
	var skill domain.Skill
	if err := r.db.First(&skill, id).Error; err != nil {
		return domain.Skill{}, err
	}
	return skill, nil
}
