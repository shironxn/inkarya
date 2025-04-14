package service

import (
	"github.com/google/uuid"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/domain"
	"github.com/shironxn/inkarya/internal/repository"
)

// User Service Interface
type UserService interface {
	// User
	CreateUser(req dto.UserCreateRequest) error
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
	UpdateUser(req dto.UserUpdateRequest) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// User Implementation
func (s *userService) CreateUser(req dto.UserCreateRequest) error {
	user := &domain.User{
		ID:           req.ID,
		Name:         req.Name,
		Email:        req.Email,
		AvatarURL:    req.AvatarURL,
		Bio:          req.Bio,
		Interest:     req.Interest,
		DOB:          req.DOB,
		Phone:        req.Phone,
		Location:     req.Location,
		Status:       req.Status,
		Availability: req.Availability,
		ResumeURL:    req.ResumeURL,
	}
	return s.repo.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]domain.User, error) {
	return s.repo.FindAllUsers()
}

func (s *userService) GetUserByID(id uuid.UUID) (*domain.User, error) {
	return s.repo.FindUserByID(id)
}

func (s *userService) UpdateUser(req dto.UserUpdateRequest) error {
	user := &domain.User{
		ID:           req.ID,
		Name:         req.Name,
		Email:        req.Email,
		AvatarURL:    req.AvatarURL,
		Bio:          req.Bio,
		Interest:     req.Interest,
		DOB:          req.DOB,
		Phone:        req.Phone,
		Location:     req.Location,
		Status:       req.Status,
		Availability: req.Availability,
		ResumeURL:    req.ResumeURL,
	}
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
