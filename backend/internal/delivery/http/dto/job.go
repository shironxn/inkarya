package dto

import "time"

type JobCreateRequest struct {
	CompanyID   uint   `json:"company_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Education   string `json:"education"`
	SalaryMin   *int   `json:"salary_min"`
	SalaryMax   *int   `json:"salary_max"`
	SkillIDs    []uint `json:"skill_ids"`
}

type JobUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Education   string `json:"education"`
	SalaryMin   *int   `json:"salary_min"`
	SalaryMax   *int   `json:"salary_max"`
	SkillIDs    []uint `json:"skill_ids"`
}

type CompanyResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	AvatarURL   string `json:"avatar_url"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

type JobResponse struct {
	ID          uint            `json:"id"`
	CompanyID   uint            `json:"company_id"`
	Company     CompanyResponse `json:"company"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Location    string          `json:"location"`
	Education   string          `json:"education"`
	SalaryMin   *int            `json:"salary_min"`
	SalaryMax   *int            `json:"salary_max"`
	Skills      []SkillResponse `json:"skills"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
