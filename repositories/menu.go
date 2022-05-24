package repository

import (
	"Hybrid-blog/models"

	"github.com/jinzhu/gorm"
)

type MenuRepository interface {
	FindAll(db *gorm.DB) ([]models.Menu, error)
	FindMenuByID(db *gorm.DB, id int) ([]models.Menu, error)
}
