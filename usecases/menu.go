package usecase

import (
	"Hybrid-blog/helpers/database"
	"Hybrid-blog/models"
	repository "Hybrid-blog/repositories"
	"Hybrid-blog/repositories/repoImlp"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var _ MenuUseCase = (*menuUseCase)(nil)

type MenuUseCase interface {
	GetAllMenu(c *gin.Context) ([]models.Menu, error)
	FindMenuByID(id int) ([]models.Menu, error)
}
type (
	menuUseCase struct {
		db             *gorm.DB
		menuRepository repository.MenuRepository
	}
)

func NewMenuUseCase(c *gin.Context) MenuUseCase {
	return &menuUseCase{
		db:             database.DB(),
		menuRepository: repoimlp.NewMenuRepository(),
	}
}
func (u *menuUseCase) GetAllMenu(c *gin.Context) ([]models.Menu, error) {

	menus, err := u.menuRepository.FindAll(u.db)
	if err != nil {
		return nil, err
	}
	return menus, nil
}
func (u *menuUseCase) FindMenuByID(id int) ([]models.Menu, error) {

	menus, err := u.menuRepository.FindMenuByID(u.db, id)
	if err != nil {
		return nil, err
	}
	return menus, nil
}
