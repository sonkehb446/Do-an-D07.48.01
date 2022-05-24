package repoimlp

import (
	"Hybrid-blog/repositories"
	"fmt"

	"github.com/jinzhu/gorm"

	"Hybrid-blog/models"
)

var _ repository.CategoryRepository = (*categoryRepository)(nil)

type categoryRepository struct {
}

func NewCategoryRepository() repository.CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) FindAllCategory(db *gorm.DB, Name, Description, Create_From, Create_To string, offset, limit int64) ([]*models.Category, int64, error) {
	var categories []*models.Category
	var countQuery int64
	result := db.Table("categories c").
		Select("c.*")
	if Name != "" {
		result = result.Where("c.category_name like ?", "%"+Name+"%")
	}

	if Description != "" {
		result = result.Where("c.description like ?", "%"+Description+"%")
	}

	if Create_From != "" {
		timefrom := fmt.Sprintf("%s 00:00:01", Create_From)
		result = result.Where("c.created_at >= ?  ", timefrom)
	}
	if Create_To != "" {
		timeto := fmt.Sprintf("%s 23:59:59", Create_To)
		result = result.Where("c.created_at <=  ?", timeto)
	}

	result = result.Where("c.parent_id ISNULL and c.deleted_at ISNULL")
	queryRow := result.Count(&countQuery)
	if queryRow.Error != nil {
		return nil, -1, result.Error
	}
	result.Order("c.id  Desc").
		Limit(limit).
		Offset(offset).
		Scan(&categories)

	return categories, countQuery, nil
}
func (r *categoryRepository) FindAllCategoryNoPagination(db *gorm.DB) (categories []models.Category, err error) {

	result := db.Table("categories c").
		Select("c.*").
		Where("c.parent_id ISNULL and c.deleted_at ISNULL").
		Order("c.id  Desc")

	if err := result.Scan(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
func (r *categoryRepository) CreateCategory(db *gorm.DB, category models.Category) error {
	result := db.Create(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *categoryRepository) EditCategory(db *gorm.DB, id int64, category *models.Category) error {
	result := db.Exec("UPDATE public.categories SET menu_id = ?, category_name = ? ,description = ? Where id = ?", category.MenuID, category.CategoryName, category.Description, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *categoryRepository) RemoveCategoryByID(db *gorm.DB, categoryID int64) error {
	result := db.Delete(&models.Category{}, categoryID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//--------------------------------------------------------------------

func (r *categoryRepository) FindAllSubCategory(db *gorm.DB, Name, Description, Create_From, Create_To, nameParent string, offset, limit int64) ([]*models.Category, int64, error) {
	var categories []*models.Category
	var countQuery int64
	result := db.Table("categories c").
		Select("c.*,s.category_name as parent_name").
		Joins("inner JOIN categories as s on s.id= c.parent_id")
	if Name != "" {
		result = result.Where("c.category_name like ?", "%"+Name+"%")
	}

	if Description != "" {
		result = result.Where("c.description like ?", "%"+Description+"%")
	}
	if nameParent != "" {
		result = result.Where("s.category_name like ?", "%"+nameParent+"%")
	}
	if Create_From != "" {
		timefrom := fmt.Sprintf("%s 00:00:01", Create_From)
		result = result.Where("c.created_at >= ?  ", timefrom)
	}
	if Create_To != "" {
		timeto := fmt.Sprintf("%s 23:59:59", Create_To)
		result = result.Where("c.created_at <=  ?", timeto)
	}

	result = result.Where("c.parent_id NOTNULL and c.deleted_at ISNULL and s.deleted_at ISNULL")
	queryRow := result.Count(&countQuery)
	if queryRow.Error != nil {
		return nil, -1, result.Error
	}
	result.Order("c.id  Desc").
		Limit(limit).
		Offset(offset).
		Scan(&categories)

	return categories, countQuery, nil
}
func (r *categoryRepository) CreateSubCategory(db *gorm.DB, category models.Category) error {
	result := db.Create(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *categoryRepository) EditSubCategory(db *gorm.DB, id int64, category *models.Category) error {
	result := db.Exec("UPDATE public.categories SET parent_id = ?, category_name = ? ,description = ? Where id = ?", category.ParentID, category.CategoryName, category.Description, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *categoryRepository) FindCategoryByID(db *gorm.DB, categoryID int64) ([]models.Category, error) {
	var category []models.Category
	result := db.Table("categories c").Select("*").
		Where("c.id = ?", categoryID).
		Scan(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (r *categoryRepository) FindSubCategoryByID(CategoryID string, db *gorm.DB) ([]models.Category, error) {
	var subCategories []models.Category

	result := db.Table("categories c").Select("*").
		Where("c.parent_id NOTNULL and c.deleted_at ISNULL").
		Where("c.parent_id = ?", CategoryID).Scan(&subCategories)
	if result.Error != nil {
		return nil, result.Error
	}
	return subCategories, nil
}

func (r *categoryRepository) FindCategoryBySubCategory(SubCategoryID string, db *gorm.DB) (*models.Category, error) {
	var category []*models.Category
	result := db.Table("categories c").Select("c.*").
		Where("c.deleted_at ISNULL").
		Where("c.id = ?", SubCategoryID).
		Scan(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(category) > 0 {
		return category[0], nil
	}
	return nil, nil
}
