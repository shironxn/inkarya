package repository

import (
	"github.com/google/uuid"
	"github.com/shironxn/inkarya/internal/domain"
	"gorm.io/gorm"
)

type PostRepository interface {
	// Post
	CreatePost(post *domain.Post) error
	FindAllPosts() ([]domain.Post, error)
	FindPostByID(id uint) (*domain.Post, error)
	UpdatePost(post *domain.Post) error
	DeletePost(id uint) error

	// Post Comment
	CreateComment(comment *domain.PostComment) error
	FindCommentsByPostID(postID uint) ([]domain.PostComment, error)
	FindCommentByID(id uint) (*domain.PostComment, error)
	UpdateComment(comment *domain.PostComment) error
	DeleteComment(id uint) error

	// Post Like
	CreateLike(like *domain.PostLike) error
	DeleteLike(userID uuid.UUID, postID uint) error
	FindLikeByUserAndPost(userID uuid.UUID, postID uint) (*domain.PostLike, error)
}

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		DB: db,
	}
}

// Post Implementation
func (r *postRepository) CreatePost(post *domain.Post) error {
	return r.DB.Create(post).Error
}

func (r *postRepository) FindAllPosts() ([]domain.Post, error) {
	var posts []domain.Post
	if err := r.DB.Preload("User").Preload("Likes").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) FindPostByID(id uint) (*domain.Post, error) {
	var post domain.Post
	if err := r.DB.Preload("User").Preload("Comments.User").Preload("Likes").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) UpdatePost(post *domain.Post) error {
	return r.DB.Model(&domain.Post{}).Where("id = ?", post.ID).Updates(post).Error
}

func (r *postRepository) DeletePost(id uint) error {
	return r.DB.Delete(&domain.Post{}, id).Error
}

// Post Comment Implementation
func (r *postRepository) CreateComment(comment *domain.PostComment) error {
	return r.DB.Create(comment).Error
}

func (r *postRepository) FindCommentsByPostID(postID uint) ([]domain.PostComment, error) {
	var comments []domain.PostComment
	if err := r.DB.Where("post_id = ?", postID).Preload("User").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *postRepository) FindCommentByID(id uint) (*domain.PostComment, error) {
	var comment domain.PostComment
	if err := r.DB.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *postRepository) UpdateComment(comment *domain.PostComment) error {
	return r.DB.Model(&domain.PostComment{}).Where("id = ?", comment.ID).Updates(comment).Error
}

func (r *postRepository) DeleteComment(id uint) error {
	return r.DB.Delete(&domain.PostComment{}, id).Error
}

// Post Like Implementation
func (r *postRepository) CreateLike(like *domain.PostLike) error {
	return r.DB.Create(like).Error
}

func (r *postRepository) DeleteLike(userID uuid.UUID, postID uint) error {
	return r.DB.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&domain.PostLike{}).Error
}

func (r *postRepository) FindLikeByUserAndPost(userID uuid.UUID, postID uint) (*domain.PostLike, error) {
	var like domain.PostLike
	if err := r.DB.Where("user_id = ? AND post_id = ?", userID, postID).First(&like).Error; err != nil {
		return nil, err
	}
	return &like, nil
}
