package repository

import (
	"github.com/google/uuid"
	"github.com/shironxn/inkarya/internal/domain"
	"gorm.io/gorm"
)

type JobRepository interface {
	// Job
	FindAllJobs() ([]domain.Job, error)
	FindJobByID(id uint) (*domain.Job, error)
	FindJobsByCompanyID(companyID uint) ([]domain.Job, error)
	SearchJobs(query string) ([]domain.Job, error)

	// Job Application
	ApplyForJob(userID uuid.UUID, jobID uint) error
	FindJobApplications() ([]domain.JobApplication, error)
	FindJobApplicationsByUserID(userID uuid.UUID) ([]domain.JobApplication, error)
	FindJobApplicationByID(id uint) (*domain.JobApplication, error)

	// Saved Jobs
	SaveJob(userID uuid.UUID, jobID uint) error
	UnsaveJob(userID uuid.UUID, jobID uint) error
	FindSavedJobs(userID uuid.UUID) ([]domain.Job, error)
}

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{db: db}
}

// Job Implementation
func (r *jobRepository) CreateJob(job *domain.Job) error {
	return r.db.Create(job).Error
}

func (r *jobRepository) FindJobApplications() ([]domain.JobApplication, error) {
	var applications []domain.JobApplication
	if err := r.db.Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func (r *jobRepository) FindAllJobs() ([]domain.Job, error) {
	var jobs []domain.Job
	if err := r.db.Preload("Skills").Preload("Disabilities").Preload("Company").Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *jobRepository) FindJobByID(id uint) (*domain.Job, error) {
	var job domain.Job
	if err := r.db.Preload("Skills").Preload("Disabilities").Preload("Company").First(&job, id).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *jobRepository) FindJobsByCompanyID(companyID uint) ([]domain.Job, error) {
	var jobs []domain.Job
	if err := r.db.Preload("Skills").Preload("Disabilities").Preload("Company").Where("company_id = ?", companyID).Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *jobRepository) SearchJobs(query string) ([]domain.Job, error) {
	var jobs []domain.Job
	if err := r.db.Preload("Skills").Preload("Disabilities").Preload("Company").
		Where("title ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%").
		Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *jobRepository) UpdateJob(job *domain.Job) error {
	return r.db.Model(&domain.Job{}).Where("id = ?", job.ID).Updates(job).Error
}

func (r *jobRepository) DeleteJob(id uint) error {
	return r.db.Delete(&domain.Job{}, id).Error
}

// Job Application Implementation
func (r *jobRepository) ApplyForJob(userID uuid.UUID, jobID uint) error {
	application := &domain.JobApplication{
		UserID: userID,
		JobID:  jobID,
	}
	return r.db.Create(application).Error
}

func (r *jobRepository) FindJobApplicationsByUserID(userID uuid.UUID) ([]domain.JobApplication, error) {
	var applications []domain.JobApplication
	if err := r.db.Where("user_id = ?", userID).Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func (r *jobRepository) FindJobApplicationByID(id uint) (*domain.JobApplication, error) {
	var application domain.JobApplication
	if err := r.db.First(&application, id).Error; err != nil {
		return nil, err
	}
	return &application, nil
}

// Saved Jobs Implementation
func (r *jobRepository) SaveJob(userID uuid.UUID, jobID uint) error {
	savedJob := &domain.SavedJob{
		UserID: userID,
		JobID:  jobID,
	}
	return r.db.Create(savedJob).Error
}

func (r *jobRepository) UnsaveJob(userID uuid.UUID, jobID uint) error {
	return r.db.Where("user_id = ? AND job_id = ?", userID, jobID).
		Delete(&domain.SavedJob{}).Error
}

func (r *jobRepository) FindSavedJobs(userID uuid.UUID) ([]domain.Job, error) {
	var jobs []domain.Job
	if err := r.db.Joins("JOIN saved_jobs ON saved_jobs.job_id = jobs.id").
		Where("saved_jobs.user_id = ?", userID).
		Preload("Skills").
		Preload("Disabilities").
		Preload("Company").
		Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}
