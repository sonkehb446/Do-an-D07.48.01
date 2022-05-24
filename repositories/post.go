package repository

import (
	"github.com/jinzhu/gorm"

	"Hybrid-blog/models"
)

type PostInterface interface {
	CreatePost(post *models.Post, db *gorm.DB) error
	FindPostByCategoryID(id int, name string, offset, limit int64, db *gorm.DB) ([]*models.PostHome, int64, error)
	GetLimitPostByType(Type string, db *gorm.DB, offset, limit string) ([]*models.PostHome, int64, error)
	FindPostByID(id string, db *gorm.DB) (*models.PostHome, error)
	DeletePostByID(id string, db *gorm.DB) error
	EditPost(post *models.Post, db *gorm.DB) error
	GetAllPostByID(id, Title, Description, Create_By, Create_from, Create_To, offset, limit string, db *gorm.DB) ([]*models.PostHome, int64, error)
}
