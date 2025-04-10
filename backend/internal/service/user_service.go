package service

import (
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/domain"
	"github.com/shironxn/inkarya/internal/repository"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user dto.CreateUserRequest) error
	GetAll() ([]domain.User, error)
	GetByID(id uint) (*domain.User, error)
	Update(user dto.UpdateUserRequest) error
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Register(user dto.CreateUserRequest) error {
	return s.repo.Create(&domain.User{
		Name:         user.Name,
		Email:        user.Email,
		AvatarURL:    user.AvatarURL,
		Bio:          user.Bio,
		Interest:     user.Interest,
		DOB:          user.DOB,
		Phone:        user.Phone,
		Location:     user.Location,
		Status:       user.Status,
		Availability: user.Availability,
		ResumeURL:    user.ResumeURL,
	})
}

func (s *userService) GetAll() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetByID(id uint) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) Update(user dto.UpdateUserRequest) error {
	return s.repo.Update(&domain.User{
		Model: gorm.Model{
			ID: user.ID,
		},
		Name:         user.Name,
		Email:        user.Email,
		AvatarURL:    user.AvatarURL,
		Bio:          user.Bio,
		Interest:     user.Interest,
		DOB:          user.DOB,
		Phone:        user.Phone,
		Location:     user.Location,
		Status:       user.Status,
		Availability: user.Availability,
		ResumeURL:    user.ResumeURL,
	})
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}
