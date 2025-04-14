package service

import (
	"github.com/google/uuid"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/domain"
	"github.com/shironxn/inkarya/internal/repository"
)

type JobService interface {
	// Job
	GetAllJobs() ([]domain.Job, error)
	GetJobByID(id uint) (*domain.Job, error)
	GetJobsByCompanyID(companyID uint) ([]domain.Job, error)
	SearchJobs(query string) ([]domain.Job, error)

	// Job Application
	ApplyForJob(req dto.JobApplicationRequest) error
	GetJobApplications() ([]domain.JobApplication, error)
	GetJobApplicationsByUserID(userID uuid.UUID) ([]domain.JobApplication, error)
	GetJobApplicationByID(id uint) (*domain.JobApplication, error)

	// Saved Jobs
	SaveJob(req dto.JobSaveRequest) error
	UnsaveJob(req dto.JobSaveRequest) error
	GetSavedJobs(userID uuid.UUID) ([]domain.Job, error)
}

type jobService struct {
	repo repository.JobRepository
}

func NewJobService(repo repository.JobRepository) JobService {
	return &jobService{
		repo: repo,
	}
}

// Job Implementation
func (s *jobService) GetAllJobs() ([]domain.Job, error) {
	return s.repo.FindAllJobs()
}

func (s *jobService) GetJobByID(id uint) (*domain.Job, error) {
	return s.repo.FindJobByID(id)
}

func (s *jobService) GetJobsByCompanyID(companyID uint) ([]domain.Job, error) {
	return s.repo.FindJobsByCompanyID(companyID)
}

func (s *jobService) SearchJobs(query string) ([]domain.Job, error) {
	return s.repo.SearchJobs(query)
}

// Job Application Implementation
func (s *jobService) ApplyForJob(req dto.JobApplicationRequest) error {
	return s.repo.ApplyForJob(req.UserID, req.JobID)
}

func (s *jobService) GetJobApplications() ([]domain.JobApplication, error) {
	return s.repo.FindJobApplications()
}

func (s *jobService) GetJobApplicationsByUserID(userID uuid.UUID) ([]domain.JobApplication, error) {
	return s.repo.FindJobApplicationsByUserID(userID)
}

func (s *jobService) GetJobApplicationByID(id uint) (*domain.JobApplication, error) {
	return s.repo.FindJobApplicationByID(id)
}

// Saved Jobs Implementation
func (s *jobService) SaveJob(req dto.JobSaveRequest) error {
	return s.repo.SaveJob(req.UserID, req.JobID)
}

func (s *jobService) UnsaveJob(req dto.JobSaveRequest) error {
	return s.repo.UnsaveJob(req.UserID, req.JobID)
}

func (s *jobService) GetSavedJobs(userID uuid.UUID) ([]domain.Job, error) {
	return s.repo.FindSavedJobs(userID)
}
