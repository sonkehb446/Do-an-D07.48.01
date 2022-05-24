package repoimlp

import (
	"Hybrid-blog/models"
	"Hybrid-blog/repositories"

	"github.com/jinzhu/gorm"
)

var _ repository.MenuRepository = (*menuRepository)(nil)

type menuRepository struct{}

func NewMenuRepository() repository.MenuRepository {
	return &menuRepository{}
}

func (r *menuRepository) FindAll(db *gorm.DB) (menus []models.Menu, err error) {
	if err := db.Find(&menus).Error; err != nil {
		return menus, err
	}
	return menus, nil
}
func (r *menuRepository) FindMenuByID(db *gorm.DB, id int) ([]models.Menu, error) {
	var menus []models.Menu
	result := db.Table("menus c").Select("*").
		Where("c.id = ?", id).
		Scan(&menus)
	if result.Error != nil {
		return nil, result.Error
	}
	return menus, nil
}
