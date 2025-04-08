package domain

import (
	"github.com/shironxn/inkarya/internal/domain/disability"
	"github.com/shironxn/inkarya/internal/domain/post"
	"github.com/shironxn/inkarya/internal/domain/skill"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        *string `gorm:"unique"`
	AvatarURL    *string
	Bio          *string
	Interest     string
	DOB          string
	Phone        *string `gorm:"unique"`
	Location     string
	Status       *string
	Availability *string
	ResumeURL    *string
	Skills       []skill.Skill           `gorm:"many2many:users_skills;"`
	Disabilities []disability.Disability `gorm:"many2many:users_disabilities;"`
	Posts        []post.Post
	PostComments []post.PostComments
	PostLikes    []post.PostLikes
}
