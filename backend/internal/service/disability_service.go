package service

import (
	"github.com/shironxn/inkarya/internal/domain"
	"github.com/shironxn/inkarya/internal/repository"
)

type DisabilityService interface {
	GetAll() ([]domain.Disability, error)
	GetByID(id uint) (domain.Disability, error)
}

type disabilityService struct {
	disabilityRepo repository.DisabilityRepository
}

func NewDisabilityService(disabilityRepo repository.DisabilityRepository) DisabilityService {
	return &disabilityService{
		disabilityRepo: disabilityRepo,
	}
}

func (s *disabilityService) GetAll() ([]domain.Disability, error) {
	return s.disabilityRepo.FindAll()
}

func (s *disabilityService) GetByID(id uint) (domain.Disability, error) {
	return s.disabilityRepo.FindByID(id)
}
