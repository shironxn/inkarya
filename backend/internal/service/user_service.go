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
	repo           repository.UserRepository
	skillRepo      repository.SkillRepository
	disabilityRepo repository.DisabilityRepository
}

func NewUserService(repo repository.UserRepository, skillRepo repository.SkillRepository, disabilityRepo repository.DisabilityRepository) UserService {
	return &userService{
		repo:           repo,
		skillRepo:      skillRepo,
		disabilityRepo: disabilityRepo,
	}
}

// User Implementation
func (s *userService) CreateUser(req dto.UserCreateRequest) error {
	// Get skills
	var skills []domain.Skill
	for _, skillID := range req.Skills {
		skill, err := s.skillRepo.FindByID(skillID)
		if err != nil {
			return err
		}
		skills = append(skills, skill)
	}

	// Get disabilities
	var disabilities []domain.Disability
	for _, disabilityID := range req.Disabilities {
		disability, err := s.disabilityRepo.FindByID(disabilityID)
		if err != nil {
			return err
		}
		disabilities = append(disabilities, disability)
	}

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
		Skills:       skills,
		Disabilities: disabilities,
	}
	return s.repo.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]domain.User, error) {
	users, err := s.repo.FindAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) GetUserByID(id uuid.UUID) (*domain.User, error) {
	user, err := s.repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(req dto.UserUpdateRequest) error {
	// Get existing user
	user, err := s.repo.FindUserByID(req.ID)
	if err != nil {
		return err
	}

	// Get skills
	var skills []domain.Skill
	for _, skillID := range req.Skills {
		skill, err := s.skillRepo.FindByID(skillID)
		if err != nil {
			return err
		}
		skills = append(skills, skill)
	}

	// Get disabilities
	var disabilities []domain.Disability
	for _, disabilityID := range req.Disabilities {
		disability, err := s.disabilityRepo.FindByID(disabilityID)
		if err != nil {
			return err
		}
		disabilities = append(disabilities, disability)
	}

	// Update user fields
	user.Name = req.Name
	user.Email = req.Email
	user.AvatarURL = req.AvatarURL
	user.Bio = req.Bio
	user.Interest = req.Interest
	user.DOB = req.DOB
	user.Phone = req.Phone
	user.Location = req.Location
	user.Status = req.Status
	user.Availability = req.Availability
	user.ResumeURL = req.ResumeURL
	user.Skills = skills
	user.Disabilities = disabilities

	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
