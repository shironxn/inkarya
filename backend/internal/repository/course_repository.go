package repository

import (
	"github.com/google/uuid"
	"github.com/shironxn/inkarya/internal/domain"

	"gorm.io/gorm"
)

type CourseRepository interface {
	// Course
	FindAllCourses() ([]domain.Course, error)
	FindCourseByID(id uint) (*domain.Course, error)

	// Lesson
	FindLessonByID(id uint) (*domain.CourseLesson, error)
	FindAllLessonsByCourseID(courseID uint) ([]domain.CourseLesson, error)

	// Enrollment
	CreateEnroll(enrollment *domain.CourseEnrollment) error
	FindEnrollByID(id uint) (*domain.CourseEnrollment, error)
	FindEnrollByCourseID(courseID uint) ([]domain.CourseEnrollment, error)
	FindEnrollByUserID(userID uuid.UUID) ([]domain.CourseEnrollment, error)
	DeleteEnroll(id uint) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) FindAllCourses() ([]domain.Course, error) {
	var courses []domain.Course
	err := r.db.Preload("Category").Find(&courses).Error
	return courses, err
}

func (r *courseRepository) FindCourseByID(id uint) (*domain.Course, error) {
	var course domain.Course
	err := r.db.Preload("Category").First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) FindLessonByID(id uint) (*domain.CourseLesson, error) {
	var lesson domain.CourseLesson
	err := r.db.First(&lesson, id).Error
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (r *courseRepository) FindAllLessonsByCourseID(courseID uint) ([]domain.CourseLesson, error) {
	var lessons []domain.CourseLesson
	if err := r.db.Where("course_id = ?", courseID).Order("`order` asc").Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}

func (r *courseRepository) CreateEnroll(enrollment *domain.CourseEnrollment) error {
	return r.db.Create(enrollment).Error
}

func (r *courseRepository) FindEnrollByID(id uint) (*domain.CourseEnrollment, error) {
	var enrollment domain.CourseEnrollment
	err := r.db.First(&enrollment, id).Error
	if err != nil {
		return nil, err
	}
	return &enrollment, nil
}

func (r *courseRepository) FindEnrollByCourseID(courseID uint) ([]domain.CourseEnrollment, error) {
	var enrollments []domain.CourseEnrollment
	err := r.db.Where("course_id = ?", courseID).Find(&enrollments).Error
	return enrollments, err
}

func (r *courseRepository) FindEnrollByUserID(userID uuid.UUID) ([]domain.CourseEnrollment, error) {
	var enrollments []domain.CourseEnrollment
	err := r.db.Where("user_id = ?", userID).Find(&enrollments).Error
	return enrollments, err
}

func (r *courseRepository) DeleteEnroll(id uint) error {
	return r.db.Delete(&domain.CourseEnrollment{}, id).Error
}
