package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
	"github.com/shironxn/inkarya/internal/domain"
	"github.com/shironxn/inkarya/internal/repository"
	"gorm.io/gorm"
)

type PostService interface {
	// Post
	CreatePost(req dto.PostCreateRequest) error
	GetAllPosts() ([]domain.Post, error)
	GetPostByID(id uint) (*domain.Post, error)
	UpdatePost(req dto.PostUpdateRequest) error
	DeletePost(req dto.PostDeleteRequest) error

	// Post Comment
	CreateComment(req dto.PostCommentCreateRequest) error
	GetCommentsByPostID(postID uint) ([]domain.PostComment, error)
	UpdateComment(req dto.PostCommentUpdateRequest) error
	DeleteComment(req dto.PostCommentDeleteRequest) error

	// Post Like
	LikePost(req dto.PostLikeRequest) error
	UnlikePost(req dto.PostUnlikeRequest) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{
		repo: repo,
	}
}

// Post Implementation
func (s *postService) CreatePost(req dto.PostCreateRequest) error {
	post := &domain.Post{
		UserID:   req.UserID,
		Title:    req.Title,
		Content:  req.Content,
		ImageUrl: req.ImageUrl,
	}
	return s.repo.CreatePost(post)
}

func (s *postService) GetAllPosts() ([]domain.Post, error) {
	return s.repo.FindAllPosts()
}

func (s *postService) GetPostByID(id uint) (*domain.Post, error) {
	return s.repo.FindPostByID(id)
}

func (s *postService) UpdatePost(req dto.PostUpdateRequest) error {
	post, err := s.repo.FindPostByID(req.ID)
	if err != nil {
		return err
	}

	if post.UserID != req.UserID {
		return fiber.ErrUnauthorized
	}

	return s.repo.UpdatePost(&domain.Post{
		Model: gorm.Model{
			ID: req.ID,
		},
		Title:    req.Title,
		Content:  req.Content,
		ImageUrl: req.ImageUrl,
	})
}

func (s *postService) DeletePost(req dto.PostDeleteRequest) error {
	post, err := s.repo.FindPostByID(req.ID)
	if err != nil {
		return err
	}

	if post.UserID != req.UserID {
		return fiber.ErrUnauthorized
	}

	return s.repo.DeletePost(req.ID)
}

// Post Comment Implementation
func (s *postService) CreateComment(req dto.PostCommentCreateRequest) error {
	post, err := s.repo.FindPostByID(req.PostID)
	if err != nil {
		return err
	}

	comment := &domain.PostComment{
		PostID:  post.ID,
		UserID:  req.UserID,
		Content: req.Content,
	}
	return s.repo.CreateComment(comment)
}

func (s *postService) GetCommentsByPostID(postID uint) ([]domain.PostComment, error) {
	_, err := s.repo.FindPostByID(postID)
	if err != nil {
		return nil, err
	}

	return s.repo.FindCommentsByPostID(postID)
}

func (s *postService) UpdateComment(req dto.PostCommentUpdateRequest) error {
	comment, err := s.repo.FindCommentByID(req.ID)
	if err != nil {
		return err
	}

	if comment.UserID != req.UserID {
		return fiber.ErrUnauthorized
	}

	return s.repo.UpdateComment(&domain.PostComment{
		Model: gorm.Model{
			ID: req.ID,
		},
		Content: req.Content,
	})
}

func (s *postService) DeleteComment(req dto.PostCommentDeleteRequest) error {
	comment, err := s.repo.FindCommentByID(req.ID)
	if err != nil {
		return err
	}

	if comment.UserID != req.UserID {
		return fiber.ErrUnauthorized
	}

	return s.repo.DeleteComment(req.ID)
}

// Post Like Implementation
func (s *postService) LikePost(req dto.PostLikeRequest) error {
	if _, err := s.repo.FindPostByID(req.PostID); err != nil {
		return err
	}

	// Check if user already liked the post
	if existingLike, err := s.repo.FindLikeByUserAndPost(req.UserID, req.PostID); err == nil && existingLike != nil {
		return nil
	}

	like := &domain.PostLike{
		UserID: req.UserID,
		PostID: req.PostID,
	}
	return s.repo.CreateLike(like)
}

func (s *postService) UnlikePost(req dto.PostUnlikeRequest) error {
	if _, err := s.repo.FindPostByID(req.PostID); err != nil {
		return err
	}

	// Check if user has liked the post
	if _, err := s.repo.FindLikeByUserAndPost(req.UserID, req.PostID); err != nil {
		return err // Like not found
	}

	return s.repo.DeleteLike(req.UserID, req.PostID)
}
