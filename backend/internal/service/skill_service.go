package service

import (
	"github.com/shironxn/inkarya/internal/domain"
	"github.com/shironxn/inkarya/internal/repository"
)

type SkillService interface {
	GetAll() ([]domain.Skill, error)
	GetByID(id uint) (domain.Skill, error)
}

type skillService struct {
	repo repository.SkillRepository
}

func NewSkillService(repo repository.SkillRepository) SkillService {
	return &skillService{
		repo: repo,
	}
}

func (s *skillService) GetAll() ([]domain.Skill, error) {
	return s.repo.FindAll()
}

func (s *skillService) GetByID(id uint) (domain.Skill, error) {
	return s.repo.FindByID(id)
}
