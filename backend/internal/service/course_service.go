package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/domain"
	"github.com/shironxn/inkarya/internal/repository"
)

// Course Service Interface
type CourseService interface {
	// Course
	GetAllCourses() ([]domain.Course, error)
	GetCourseByID(id uint) (*domain.Course, error)

	// Lesson
	GetLessonByID(id uint) (*domain.CourseLesson, error)
	GetAllLessonsByCourseID(courseID uint) ([]domain.CourseLesson, error)

	// Enrollment
	EnrollCourse(req *dto.CourseEnrollmentCreateRequest) error
	GetEnrollByID(id uint) (*domain.CourseEnrollment, error)
	GetEnrollByCourseID(courseID uint) ([]domain.CourseEnrollment, error)
	GetEnrollByUserID(userID uuid.UUID) ([]domain.CourseEnrollment, error)
	UnenrollCourse(req *dto.CourseEnrollmentDeleteRequest) error
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{
		repo: repo,
	}
}

// Course Implementation
func (s *courseService) GetAllCourses() ([]domain.Course, error) {
	return s.repo.FindAllCourses()
}

func (s *courseService) GetCourseByID(id uint) (*domain.Course, error) {
	return s.repo.FindCourseByID(id)
}

// Course Lesson Implementation
func (s *courseService) GetLessonByID(id uint) (*domain.CourseLesson, error) {
	return s.repo.FindLessonByID(id)
}

func (s *courseService) GetAllLessonsByCourseID(courseID uint) ([]domain.CourseLesson, error) {
	return s.repo.FindAllLessonsByCourseID(courseID)
}

// Course Enrollment Implementation
func (s *courseService) EnrollCourse(req *dto.CourseEnrollmentCreateRequest) error {
	enrollment := &domain.CourseEnrollment{
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}
	return s.repo.CreateEnroll(enrollment)
}

func (s *courseService) GetEnrollByID(id uint) (*domain.CourseEnrollment, error) {
	return s.repo.FindEnrollByID(id)
}

func (s *courseService) GetEnrollByCourseID(courseID uint) ([]domain.CourseEnrollment, error) {
	return s.repo.FindEnrollByCourseID(courseID)
}

func (s *courseService) GetEnrollByUserID(userID uuid.UUID) ([]domain.CourseEnrollment, error) {
	return s.repo.FindEnrollByUserID(userID)
}

func (s *courseService) UnenrollCourse(req *dto.CourseEnrollmentDeleteRequest) error {
	// Check if user is enrolled in the course
	enrollment, err := s.repo.FindEnrollByID(req.ID)
	if err != nil {
		return err
	}

	if enrollment.UserID != req.UserID {
		return fiber.ErrUnauthorized
	}

	return s.repo.DeleteEnroll(enrollment.ID)
}
