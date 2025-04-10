package dto

type CreateUserRequest struct {
	Name         string  `json:"name" validate:"required"`
	Email        *string `json:"email" validate:"required,email"`
	AvatarURL    *string `json:"avatar_url"`
	Bio          *string `json:"bio"`
	Interest     string  `json:"interest" validate:"required"`
	DOB          string  `json:"dob" validate:"required"`
	Phone        *string `json:"phone"`
	Location     string  `json:"location" validate:"required"`
	Status       *string `json:"status"`
	Availability *string `json:"availability"`
	ResumeURL    *string `json:"resume_url"`
}

type UpdateUserRequest struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Email        *string `json:"email"`
	AvatarURL    *string `json:"avatar_url"`
	Bio          *string `json:"bio"`
	Interest     string  `json:"interest"`
	DOB          string  `json:"dob"`
	Phone        *string `json:"phone"`
	Location     string  `json:"location"`
	Status       *string `json:"status"`
	Availability *string `json:"availability"`
	ResumeURL    *string `json:"resume_url"`
}

type UserResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Email        *string `json:"email"`
	AvatarURL    *string `json:"avatar_url"`
	Bio          *string `json:"bio"`
	Interest     string  `json:"interest"`
	DOB          string  `json:"dob"`
	Phone        *string `json:"phone"`
	Location     string  `json:"location"`
	Status       *string `json:"status"`
	Availability *string `json:"availability"`
	ResumeURL    *string `json:"resume_url"`
}

type UserMinimalResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
