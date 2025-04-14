package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/domain"
	"github.com/shironxn/inkarya/internal/repository"
	"gorm.io/gorm"
)

// Forum Service Interface
type ForumService interface {
	// Forum
	CreateForum(req dto.ForumCreateRequest) error
	GetAllForums() ([]domain.Forum, error)
	GetForumByID(id uint) (*domain.Forum, error)
	UpdateForum(req dto.ForumUpdateRequest) error
	DeleteForum(req dto.ForumDeleteRequest) error

	// Forum Category
	GetAllCategories() ([]domain.ForumCategory, error)
	GetCategoryByID(id uint) (*domain.ForumCategory, error)

	// Forum Comment
	CreateComment(req dto.ForumCommentCreateRequest) error
	GetCommentsByForumID(forumID uint) ([]domain.ForumComment, error)
	GetCommentByID(id uint) (*domain.ForumComment, error)
	UpdateComment(req dto.ForumCommentUpdateRequest) error
	DeleteComment(id uint) error
}

type forumService struct {
	repo repository.ForumRepository
}

func NewForumService(repo repository.ForumRepository) ForumService {
	return &forumService{
		repo: repo,
	}
}

// Forum Implementation
func (s *forumService) CreateForum(req dto.ForumCreateRequest) error {
	return s.repo.CreateForum(&domain.Forum{
		UserID:     req.UserID,
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
	})
}

func (s *forumService) GetAllForums() ([]domain.Forum, error) {
	return s.repo.FindAllForums()
}

func (s *forumService) GetForumByID(id uint) (*domain.Forum, error) {
	return s.repo.FindForumByID(id)
}

func (s *forumService) UpdateForum(req dto.ForumUpdateRequest) error {
	// Check if forum exists and belongs to the user
	forum, err := s.repo.FindForumByID(req.ID)
	if err != nil {
		return err
	}

	// Verify ownership
	if forum.UserID != req.UserID {
		return fiber.ErrUnauthorized
	}

	return s.repo.UpdateForum(&domain.Forum{
		Model: gorm.Model{
			ID: req.ID,
		},
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
	})
}

func (s *forumService) DeleteForum(req dto.ForumDeleteRequest) error {
	// Check if forum exists and belongs to the user
	forum, err := s.repo.FindForumByID(req.ID)
	if err != nil {
		return err
	}

	// Verify ownership
	if forum.UserID != req.UserID {
		return fiber.ErrUnauthorized
	}

	return s.repo.DeleteForum(req.ID)
}

// Forum Category Implementation
func (s *forumService) GetAllCategories() ([]domain.ForumCategory, error) {
	return s.repo.FindAllCategories()
}

func (s *forumService) GetCategoryByID(id uint) (*domain.ForumCategory, error) {
	return s.repo.FindCategoryByID(id)
}

// Forum Comment Implementation
func (s *forumService) CreateComment(req dto.ForumCommentCreateRequest) error {
	return s.repo.CreateComment(&domain.ForumComment{
		Content: req.Content,
		ForumID: req.ForumID,
		UserID:  req.UserID,
	})
}

func (s *forumService) GetCommentsByForumID(forumID uint) ([]domain.ForumComment, error) {
	return s.repo.FindCommentsByForumID(forumID)
}

func (s *forumService) GetCommentByID(id uint) (*domain.ForumComment, error) {
	return s.repo.FindCommentByID(id)
}

func (s *forumService) UpdateComment(req dto.ForumCommentUpdateRequest) error {
	return s.repo.UpdateComment(&domain.ForumComment{
		Model: gorm.Model{
			ID: req.ID,
		},
		Content: req.Content,
		UserID:  req.UserID,
	})
}

func (s *forumService) DeleteComment(id uint) error {
	return s.repo.DeleteComment(id)
}
