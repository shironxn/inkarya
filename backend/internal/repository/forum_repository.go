package repository

import (
	"github.com/shironxn/inkarya/internal/domain"
	"gorm.io/gorm"
)

// Forum Repository Interface
type ForumRepository interface {
	// Forum
	CreateForum(forum *domain.Forum) error
	FindAllForums() ([]domain.Forum, error)
	FindForumByID(id uint) (*domain.Forum, error)
	UpdateForum(forum *domain.Forum) error
	DeleteForum(id uint) error

	// Forum Category
	FindAllCategories() ([]domain.ForumCategory, error)
	FindCategoryByID(id uint) (*domain.ForumCategory, error)

	// Forum Comment
	CreateComment(comment *domain.ForumComment) error
	FindCommentsByForumID(forumID uint) ([]domain.ForumComment, error)
	FindCommentByID(id uint) (*domain.ForumComment, error)
	UpdateComment(comment *domain.ForumComment) error
	DeleteComment(id uint) error
}

type forumRepository struct {
	DB *gorm.DB
}

func NewForumRepository(db *gorm.DB) ForumRepository {
	return &forumRepository{
		DB: db,
	}
}

// Forum Implementation
func (r *forumRepository) CreateForum(forum *domain.Forum) error {
	return r.DB.Create(forum).Error
}

func (r *forumRepository) FindAllForums() ([]domain.Forum, error) {
	var forums []domain.Forum
	if err := r.DB.Preload("Category").Find(&forums).Error; err != nil {
		return nil, err
	}
	return forums, nil
}

func (r *forumRepository) FindForumByID(id uint) (*domain.Forum, error) {
	var forum domain.Forum
	if err := r.DB.Preload("Category").First(&forum, id).Error; err != nil {
		return nil, err
	}
	return &forum, nil
}

func (r *forumRepository) UpdateForum(forum *domain.Forum) error {
	return r.DB.Model(&domain.Forum{}).Where("id = ?", forum.ID).Updates(forum).Error
}

func (r *forumRepository) DeleteForum(id uint) error {
	return r.DB.Delete(&domain.Forum{}, id).Error
}

// Forum Category Implementation
func (r *forumRepository) FindAllCategories() ([]domain.ForumCategory, error) {
	var categories []domain.ForumCategory
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *forumRepository) FindCategoryByID(id uint) (*domain.ForumCategory, error) {
	var category domain.ForumCategory
	if err := r.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// Forum Comment Implementation
func (r *forumRepository) CreateComment(comment *domain.ForumComment) error {
	return r.DB.Create(comment).Error
}

func (r *forumRepository) FindCommentsByForumID(forumID uint) ([]domain.ForumComment, error) {
	var comments []domain.ForumComment
	if err := r.DB.Where("forum_id = ?", forumID).Preload("User").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *forumRepository) FindCommentByID(id uint) (*domain.ForumComment, error) {
	var comment domain.ForumComment
	if err := r.DB.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *forumRepository) UpdateComment(comment *domain.ForumComment) error {
	return r.DB.Model(&domain.ForumComment{}).Where("id = ?", comment.ID).Updates(comment).Error
}

func (r *forumRepository) DeleteComment(id uint) error {
	return r.DB.Delete(&domain.ForumComment{}, id).Error
}
