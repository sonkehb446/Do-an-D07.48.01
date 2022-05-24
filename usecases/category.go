package usecase

import (
	"Hybrid-blog/repositories/repoImlp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"Hybrid-blog/helpers/database"
	"Hybrid-blog/models"
	repository "Hybrid-blog/repositories"
)

var _ CategoryUseCase = (*categoryUseCase)(nil)

type CategoryUseCase interface {
	GetAllCategory(Name, Description, Create_From, Create_To string, offset, limit int64) ([]*models.Category, int64, error)
	GetAllCategoryNoPagination(c *gin.Context) ([]models.Category, error)
	CreateCategory(category models.Category) error
	UpdateCategory(category *models.Category, id int64) error
	DeleteCategoryByID(id int64) error

	//-------------------------------------------------------------------
	GetAllSubCategory(Name, Description, Create_From, Create_To, nameParent string, offset, limit int64) ([]*models.Category, int64, error)
	CreateSubCategory(category models.Category) error
	UpdateSubCategory(category *models.Category, id int64) error
	FindCategoryByID(CategoryID int64) ([]models.Category, error)

	//devSon
	FindSubCategoryByID(CategoryID int64) ([]models.Category, error)
	FindCategoryBySubCategory(SubCategoryID int64) (*models.Category, error)
}
type (
	categoryUseCase struct {
		db                 *gorm.DB
		categoryRepository repository.CategoryRepository
	}
)

func NewCategoryUseCase() CategoryUseCase {
	return &categoryUseCase{
		db:                 database.DB(),
		categoryRepository: repoimlp.NewCategoryRepository(),
	}
}
func (u *categoryUseCase) GetAllCategory(Name, Description, Create_From, Create_To string, offset, limit int64) ([]*models.Category, int64, error) {
	categories, countQuery, err := u.categoryRepository.FindAllCategory(u.db, Name, Description, Create_From, Create_To, offset, limit)
	if err != nil {
		return nil, -1, err
	}
	return categories, countQuery, nil
}
func (u *categoryUseCase) GetAllCategoryNoPagination(c *gin.Context) ([]models.Category, error) {
	categories, err := u.categoryRepository.FindAllCategoryNoPagination(u.db)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
func (u *categoryUseCase) CreateCategory(category models.Category) error {
	err := u.categoryRepository.CreateCategory(u.db, category)
	if err != nil {
		return err
	}
	return nil
}
func (u *categoryUseCase) UpdateCategory(category *models.Category, id int64) error {
	err := u.categoryRepository.EditCategory(u.db, id, category)
	if err != nil {
		return err
	}
	return nil
}
func (u *categoryUseCase) DeleteCategoryByID(id int64) error {
	err := u.categoryRepository.RemoveCategoryByID(u.db, id)
	if err != nil {
		return err
	}
	return nil
}

//------------------------------------------------------------------------------------
func (u *categoryUseCase) GetAllSubCategory(Name, Description, Create_From, Create_To, nameParent string, offset, limit int64) ([]*models.Category, int64, error) {
	sub_categories, countQuery, err := u.categoryRepository.FindAllSubCategory(u.db, Name, Description, Create_From, Create_To, nameParent, offset, limit)
	if err != nil {
		return nil, -1, err
	}
	return sub_categories, countQuery, nil
}
func (u *categoryUseCase) CreateSubCategory(category models.Category) error {
	err := u.categoryRepository.CreateSubCategory(u.db, category)
	if err != nil {
		return err
	}
	return nil
}
func (u *categoryUseCase) UpdateSubCategory(category *models.Category, id int64) error {
	err := u.categoryRepository.EditSubCategory(u.db, id, category)
	if err != nil {
		return err
	}
	return nil
}

func (u *categoryUseCase) FindCategoryByID(id int64) (category []models.Category, err error) {
	category, err = u.categoryRepository.FindCategoryByID(u.db, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

//devSon
func (u *categoryUseCase) FindSubCategoryByID(CategoryID int64) ([]models.Category, error) {
	id := strconv.Itoa(int(CategoryID))
	subcate, err := u.categoryRepository.FindSubCategoryByID(id, u.db)
	if err != nil {
		return nil, err
	}
	return subcate, nil
}

func (u *categoryUseCase) FindCategoryBySubCategory(SubCategoryID int64) (*models.Category, error) {
	id := strconv.Itoa(int(SubCategoryID))
	subcate, err := u.categoryRepository.FindCategoryBySubCategory(id, u.db)
	if err != nil {
		return nil, err
	}
	return subcate, nil
}
