package usecase

import (
	"github.com/jinzhu/gorm"

	"Hybrid-blog/helpers/database"
	"Hybrid-blog/models"
	repository "Hybrid-blog/repositories"
)

type UserCaseImage interface {
	CreateImage(image *models.Image) error
	GetImage(id int64) (error, models.Image)
}

type userCaseImage struct {
	db        *gorm.DB
	imageRepo repository.ImageInterface
}

func NewUserCaseImage() UserCaseImage {
	return &userCaseImage{
		db:        database.DB(),
		imageRepo: repository.NewImageRepo(),
	}
}

func (i *userCaseImage) CreateImage(image *models.Image) error {
	db := database.DB()
	err := i.imageRepo.CreateImage(image, db)
	if err != nil {
		return err
	}
	return nil
}

func (i *userCaseImage) GetImage(id int64) (error, models.Image) {
	return nil, models.Image{}
}
