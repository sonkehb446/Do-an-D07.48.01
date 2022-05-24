package repository

import (
	"github.com/jinzhu/gorm"

	"Hybrid-blog/models"
)

type imageRepo struct{}

type ImageInterface interface {
	CreateImage(image *models.Image, db *gorm.DB) error
}

func NewImageRepo() ImageInterface {
	return &imageRepo{}
}

func (r *imageRepo) CreateImage(image *models.Image, db *gorm.DB) error {
	// if *image.Base64 == "" {
	// 	*image.Base64 = "null"
	// }
	result := db.Create(image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
