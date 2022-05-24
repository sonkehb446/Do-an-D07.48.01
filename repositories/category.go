package repository

import (
	"github.com/jinzhu/gorm"

	"Hybrid-blog/models"
)

type CategoryRepository interface {
	FindAllCategory(db *gorm.DB, Name, Description, Create_From, Create_To string, offset, limit int64) ([]*models.Category, int64, error)
	FindAllCategoryNoPagination(db *gorm.DB) ([]models.Category, error)
	CreateCategory(db *gorm.DB, category models.Category) error
	EditCategory(db *gorm.DB, id int64, category *models.Category) error
	RemoveCategoryByID(db *gorm.DB, CategoryID int64) error

	FindCategoryByID(db *gorm.DB, CategoryID int64) ([]models.Category, error)

	FindAllSubCategory(db *gorm.DB, Name, Description, Create_From, Create_To, nameParent string, offset, limit int64) ([]*models.Category, int64, error)
	CreateSubCategory(db *gorm.DB, category models.Category) error
	EditSubCategory(db *gorm.DB, id int64, category *models.Category) error

	//devSon
	FindSubCategoryByID(CategoryID string, db *gorm.DB) ([]models.Category, error)
	FindCategoryBySubCategory(SubCategoryID string, db *gorm.DB) (*models.Category, error)
}
