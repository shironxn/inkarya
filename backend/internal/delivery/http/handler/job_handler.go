package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/service"
	"github.com/shironxn/inkarya/pkg"
	"gorm.io/gorm"
)

// Job Handler Interface
type JobHandler interface {
	// Job
	GetAllJobs(c *fiber.Ctx) error
	GetJobByID(c *fiber.Ctx) error
	GetJobsByCompanyID(c *fiber.Ctx) error
	SearchJobs(c *fiber.Ctx) error

	// Job Application
	ApplyForJob(c *fiber.Ctx) error
	GetJobApplications(c *fiber.Ctx) error
	GetJobApplicationsByUserID(c *fiber.Ctx) error
	GetJobApplicationByID(c *fiber.Ctx) error

	// Saved Jobs
	SaveJob(c *fiber.Ctx) error
	UnsaveJob(c *fiber.Ctx) error
	GetSavedJobs(c *fiber.Ctx) error
}

type jobHandler struct {
	service   service.JobService
	validator pkg.ValidatorService
	jwt       pkg.JWTService
}

func NewJobHandler(service service.JobService, validator pkg.ValidatorService, jwt pkg.JWTService) JobHandler {
	return &jobHandler{
		service:   service,
		validator: validator,
		jwt:       jwt,
	}
}

// Job Implementation
func (h *jobHandler) GetAllJobs(c *fiber.Ctx) error {
	jobs, err := h.service.GetAllJobs()
	if err != nil {
		return err
	}

	var jobResponses []dto.JobResponse
	for i, job := range jobs {
		jobResponses = append(jobResponses, dto.JobResponse{
			ID:        job.ID,
			CompanyID: job.CompanyID,
			Company: dto.CompanyResponse{
				ID:          job.Company.ID,
				Name:        job.Company.Name,
				AvatarURL:   job.Company.AvatarURL,
				Location:    job.Company.Location,
				Description: job.Company.Description,
			},
			Title:       job.Title,
			Description: job.Description,
			Location:    job.Location,
			Education:   job.Education,
			SalaryMin:   job.SalaryMin,
			SalaryMax:   job.SalaryMax,
			Skills:      make([]dto.SkillResponse, len(job.Skills)),
			CreatedAt:   job.CreatedAt,
			UpdatedAt:   job.UpdatedAt,
		})
		for j, skill := range job.Skills {
			jobResponses[i].Skills[j] = dto.SkillResponse{
				ID:          skill.ID,
				Name:        skill.Name,
				Description: skill.Description,
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "jobs retrieved successfully",
		Data:    jobResponses,
	})
}

func (h *jobHandler) GetJobByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid job id")
	}

	job, err := h.service.GetJobByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "job not found")
		}
		return err
	}

	jobResponse := dto.JobResponse{
		ID:        job.ID,
		CompanyID: job.CompanyID,
		Company: dto.CompanyResponse{
			ID:          job.Company.ID,
			Name:        job.Company.Name,
			AvatarURL:   job.Company.AvatarURL,
			Location:    job.Company.Location,
			Description: job.Company.Description,
		},
		Title:       job.Title,
		Description: job.Description,
		Location:    job.Location,
		Education:   job.Education,
		SalaryMin:   job.SalaryMin,
		SalaryMax:   job.SalaryMax,
		Skills:      make([]dto.SkillResponse, len(job.Skills)),
		CreatedAt:   job.CreatedAt,
		UpdatedAt:   job.UpdatedAt,
	}

	for i, skill := range job.Skills {
		jobResponse.Skills[i] = dto.SkillResponse{
			ID:          skill.ID,
			Name:        skill.Name,
			Description: skill.Description,
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "job retrieved successfully",
		Data:    jobResponse,
	})
}

func (h *jobHandler) GetJobsByCompanyID(c *fiber.Ctx) error {
	companyID, err := c.ParamsInt("company_id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid company id")
	}

	jobs, err := h.service.GetJobsByCompanyID(uint(companyID))
	if err != nil {
		return err
	}

	jobResponses := make([]dto.JobResponse, len(jobs))
	for i, job := range jobs {
		jobResponses[i] = dto.JobResponse{
			ID:        job.ID,
			CompanyID: job.CompanyID,
			Company: dto.CompanyResponse{
				ID:          job.Company.ID,
				Name:        job.Company.Name,
				AvatarURL:   job.Company.AvatarURL,
				Location:    job.Company.Location,
				Description: job.Company.Description,
			},
			Title:       job.Title,
			Description: job.Description,
			Location:    job.Location,
			Education:   job.Education,
			SalaryMin:   job.SalaryMin,
			SalaryMax:   job.SalaryMax,
			Skills:      make([]dto.SkillResponse, len(job.Skills)),
			CreatedAt:   job.CreatedAt,
			UpdatedAt:   job.UpdatedAt,
		}
		for j, skill := range job.Skills {
			jobResponses[i].Skills[j] = dto.SkillResponse{
				ID:          skill.ID,
				Name:        skill.Name,
				Description: skill.Description,
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "jobs retrieved successfully",
		Data:    jobResponses,
	})
}

func (h *jobHandler) SearchJobs(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return fiber.NewError(fiber.StatusBadRequest, "search query is required")
	}

	jobs, err := h.service.SearchJobs(query)
	if err != nil {
		return err
	}

	var jobResponses []dto.JobResponse
	for i, job := range jobs {
		jobResponses = append(jobResponses, dto.JobResponse{
			ID:        job.ID,
			CompanyID: job.CompanyID,
			Company: dto.CompanyResponse{
				ID:          job.Company.ID,
				Name:        job.Company.Name,
				AvatarURL:   job.Company.AvatarURL,
				Location:    job.Company.Location,
				Description: job.Company.Description,
			},
			Title:       job.Title,
			Description: job.Description,
			Location:    job.Location,
			Education:   job.Education,
			SalaryMin:   job.SalaryMin,
			SalaryMax:   job.SalaryMax,
			Skills:      make([]dto.SkillResponse, len(job.Skills)),
			CreatedAt:   job.CreatedAt,
			UpdatedAt:   job.UpdatedAt,
		})
		for j, skill := range job.Skills {
			jobResponses[i].Skills[j] = dto.SkillResponse{
				ID:          skill.ID,
				Name:        skill.Name,
				Description: skill.Description,
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "jobs retrieved successfully",
		Data:    jobResponses,
	})
}

// Job Application Implementation
func (h *jobHandler) ApplyForJob(c *fiber.Ctx) error {
	jobID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid job id")
	}

	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req := dto.JobApplicationRequest{
		UserID: userID,
		JobID:  uint(jobID),
	}

	if err := h.service.ApplyForJob(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "job application submitted successfully",
	})
}

func (h *jobHandler) GetJobApplications(c *fiber.Ctx) error {
	applications, err := h.service.GetJobApplications()
	if err != nil {
		return err
	}

	var applicationResponses []dto.JobApplicationResponse
	for _, application := range applications {
		applicationResponses = append(applicationResponses, dto.JobApplicationResponse{
			ID:        application.ID,
			JobID:     application.JobID,
			UserID:    application.UserID,
			CreatedAt: application.CreatedAt,
			UpdatedAt: application.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "job applications retrieved successfully",
		Data:    applicationResponses,
	})
}

func (h *jobHandler) GetJobApplicationsByUserID(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	applications, err := h.service.GetJobApplicationsByUserID(userID)
	if err != nil {
		return err
	}

	var applicationResponses []dto.JobApplicationResponse
	for _, application := range applications {
		applicationResponses = append(applicationResponses, dto.JobApplicationResponse{
			ID:        application.ID,
			JobID:     application.JobID,
			UserID:    application.UserID,
			CreatedAt: application.CreatedAt,
			UpdatedAt: application.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "job applications retrieved successfully",
		Data:    applicationResponses,
	})
}

func (h *jobHandler) GetJobApplicationByID(c *fiber.Ctx) error {
	applicationID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid application id")
	}

	application, err := h.service.GetJobApplicationByID(uint(applicationID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "job application not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "job application retrieved successfully",
		Data: dto.JobApplicationResponse{
			ID:        application.ID,
			JobID:     application.JobID,
			UserID:    application.UserID,
			CreatedAt: application.CreatedAt,
			UpdatedAt: application.UpdatedAt,
		},
	})
}

// Saved Jobs Implementation
func (h *jobHandler) SaveJob(c *fiber.Ctx) error {
	jobID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid job id")
	}

	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req := dto.JobSaveRequest{
		UserID: userID,
		JobID:  uint(jobID),
	}

	if err := h.service.SaveJob(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "job saved successfully",
	})
}

func (h *jobHandler) UnsaveJob(c *fiber.Ctx) error {
	jobID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid job id")
	}

	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req := dto.JobSaveRequest{
		UserID: userID,
		JobID:  uint(jobID),
	}

	if err := h.service.UnsaveJob(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "job unsaved successfully",
	})
}

func (h *jobHandler) GetSavedJobs(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	jobs, err := h.service.GetSavedJobs(userID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "saved jobs retrieved successfully",
		Data:    jobs,
	})
}
