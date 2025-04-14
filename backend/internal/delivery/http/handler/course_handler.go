package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/service"
	"github.com/shironxn/inkarya/pkg"
	"gorm.io/gorm"
)

type CourseHandler interface {
	// Course
	GetAllCourses(c *fiber.Ctx) error
	GetCourseByID(c *fiber.Ctx) error

	// Lesson
	GetLessonByID(c *fiber.Ctx) error

	// Enrollment
	EnrollCourse(c *fiber.Ctx) error
	GetEnrollByID(c *fiber.Ctx) error
	GetEnrollByCourseID(c *fiber.Ctx) error
	GetEnrollByUserID(c *fiber.Ctx) error
	UnenrollCourse(c *fiber.Ctx) error
}

type courseHandler struct {
	service   service.CourseService
	validator pkg.ValidatorService
	jwt       pkg.JWTService
}

func NewCourseHandler(service service.CourseService, validator pkg.ValidatorService, jwt pkg.JWTService) CourseHandler {
	return &courseHandler{
		service:   service,
		validator: validator,
		jwt:       jwt,
	}
}

func (h *courseHandler) GetAllCourses(c *fiber.Ctx) error {
	result, err := h.service.GetAllCourses()
	if err != nil {
		return err
	}

	var courses []dto.CourseResponse
	for _, course := range result {
		courses = append(courses, dto.CourseResponse{
			ID:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			ImageURL:    course.ImageURL,
			CategoryID:  course.CategoryID,
			CreatedAt:   course.CreatedAt,
			UpdatedAt:   course.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "courses retrieved successfully",
		Data:    courses,
	})
}

func (h *courseHandler) GetCourseByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid course id")
	}

	course, err := h.service.GetCourseByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "course not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "course retrieved successfully",
		Data: dto.CourseResponse{
			ID:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			ImageURL:    course.ImageURL,
			CategoryID:  course.CategoryID,
			CreatedAt:   course.CreatedAt,
			UpdatedAt:   course.UpdatedAt,
		},
	})
}

// Lesson handlers
func (h *courseHandler) GetLessonByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid lesson id")
	}

	lesson, err := h.service.GetLessonByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "lesson not found")
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "lesson retrieved successfully",
		Data: dto.CourseLessonResponse{
			ID:        lesson.ID,
			Title:     lesson.Title,
			Content:   lesson.Content,
			CourseID:  lesson.CourseID,
			CreatedAt: lesson.CreatedAt,
			UpdatedAt: lesson.UpdatedAt,
		},
	})
}

// Enrollment handlers
func (h *courseHandler) EnrollCourse(c *fiber.Ctx) error {
	courseID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid course id")
	}

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req := dto.CourseEnrollmentCreateRequest{
		CourseID: uint(courseID),
		UserID:   userID,
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Success: false,
			Status:  fiber.StatusBadRequest,
			Message: "validation failed",
			Errors:  err,
		})
	}

	if err := h.service.EnrollCourse(&req); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid course id or user id")
		}
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusCreated,
		Message: "user enrolled successfully",
	})
}

func (h *courseHandler) GetEnrollByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid enrollment id")
	}

	enrollment, err := h.service.GetEnrollByID(uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "enrollment retrieved successfully",
		Data: dto.CourseEnrollmentResponse{
			ID:        enrollment.ID,
			UserID:    enrollment.UserID,
			CourseID:  enrollment.CourseID,
			CreatedAt: enrollment.CreatedAt,
			UpdatedAt: enrollment.UpdatedAt,
		},
	})
}

func (h *courseHandler) GetEnrollByCourseID(c *fiber.Ctx) error {
	courseID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid course id")
	}

	enrollments, err := h.service.GetEnrollByCourseID(uint(courseID))
	if err != nil {
		return err
	}

	var enrollmentsResponses []dto.CourseEnrollmentResponse
	for _, enrollment := range enrollments {
		enrollmentsResponses = append(enrollmentsResponses, dto.CourseEnrollmentResponse{
			ID:        enrollment.ID,
			UserID:    enrollment.UserID,
			CourseID:  enrollment.CourseID,
			CreatedAt: enrollment.CreatedAt,
			UpdatedAt: enrollment.UpdatedAt,
		})
	}
	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "enrollments retrieved successfully",
		Data:    enrollmentsResponses,
	})
}

func (h *courseHandler) GetEnrollByUserID(c *fiber.Ctx) error {
	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	enrollments, err := h.service.GetEnrollByUserID(userID)
	if err != nil {
		return err
	}

	var enrollmentsResponses []dto.CourseEnrollmentResponse
	for _, enrollment := range enrollments {
		enrollmentsResponses = append(enrollmentsResponses, dto.CourseEnrollmentResponse{
			ID: enrollment.ID,
		})
	}
	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "enrollments retrieved successfully",
		Data:    enrollmentsResponses,
	})
}

func (h *courseHandler) UnenrollCourse(c *fiber.Ctx) error {
	courseID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid course id")
	}

	// Get user ID from JWT token
	token := c.Locals("user").(*jwt.Token)
	userID, err := h.jwt.GetUserID(token)
	if err != nil {
		return err
	}

	req := dto.CourseEnrollmentDeleteRequest{
		ID:     uint(courseID),
		UserID: userID,
	}

	if err := h.service.UnenrollCourse(&req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "enrollment not found")
		}
		if err.Error() == "user is not enrolled in this course" {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "user unenrolled successfully",
	})
}
