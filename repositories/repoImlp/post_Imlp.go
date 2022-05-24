package repoimlp

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"Hybrid-blog/models"
	model "Hybrid-blog/models"
	repo "Hybrid-blog/repositories"
)

type postRepo struct {
}

func NewPostRepo() repo.PostInterface {
	return &postRepo{}
}

func (r *postRepo) CreatePost(post *models.Post, db *gorm.DB) error {
	result := db.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *postRepo) GetAllPostByID(id, Title, Description, Create_By, Create_from, Create_To, offset, limit string, db *gorm.DB) ([]*models.PostHome, int64, error) {
	var users []*model.PostHome
	var countQuery int64
	result := db.Table("public.posts as p").
		Select("p.id as post_id, u.full_name as author ,m.name as item , p.title, c.category_name as category, p.description, i.url as image_url, p.content, p.created_at").
		Joins("LEFT JOIN public.images i on i.id = p.image_id").
		Joins("LEFT JOin public.users u on u.id = p.user_id").
		Joins("LEFT JOIN public.categories c ON c.id = p.category_id").
		Joins("LEFT JOIN public.menus m ON m.id = c.menu_id")

	if id != "" {
		result = result.Where("p.user_id = ?", id)
	}

	if Title != "" {
		result = result.Where("p.title like ?", "%"+Title+"%")
	}

	if Create_By != "" {
		result = result.Where("u.full_name like ?", "%"+Create_By+"%")
	}

	if Description != "" {
		result = result.Where("p.description like ?", "%"+Description+"%")
	}

	if Create_from != "" {
		timefrom := fmt.Sprintf("%s 00:00:01", Create_from)
		result = result.Where("p.created_at >= ?  ", timefrom)
	}
	if Create_To != "" {
		timeto := fmt.Sprintf("%s 23:59:59", Create_To)
		result = result.Where("p.created_at <=  ?", timeto)
	}
	result = result.Where("p.deleted_at ISNULL")
	queryRow := result.Count(&countQuery)
	result.Order("u.id  Desc").Offset(offset).Limit(limit).Scan(&users)
	if queryRow.Error != nil {
		return nil, -1, result.Error
	}
	log.Println(len(users))
	return users, countQuery, nil
}

func (r *postRepo) GetLimitPostByType(Type string, db *gorm.DB, offset, limit string) ([]*models.PostHome, int64, error) {
	var posts []*models.PostHome
	var countQuery int64
	result := db.Table("public.posts as p").
		Select("p.id as post_id, u.full_name as author ,m.name as item , p.title, c.category_name as category, p.description, i.url as image_url, p.content, p.created_at").
		Joins("LEFT JOIN public.images i on i.id = p.image_id").
		Joins("LEFT JOin public.users u on u.id = p.user_id").
		Joins("LEFT JOIN public.categories c ON c.id = p.category_id").
		Joins("LEFT JOIN public.menus m ON m.id = c.menu_id").
		Where("c.menu_id = ? and p.deleted_at isnull", Type)
	queryRow := result.Count(&countQuery)
	result.Order("p.id  Desc").Offset(offset).Limit(limit).Scan(&posts)
	if queryRow.Error != nil {
		return nil, -1, result.Error
	}
	return posts, countQuery, nil
}

func (r *postRepo) FindPostByID(id string, db *gorm.DB) (*models.PostHome, error) {
	var post models.PostHome
	result := db.Table("public.posts as p").
		Select("p.id as post_id, p.user_id, u.full_name as author ,m.name as item , p.image_id, p.title,p.category_id as category_id ,c.category_name as category, p.description, i.url as image_url, p.content, p.created_at").
		Joins("LEFT JOIN public.images i on i.id = p.image_id").
		Joins("LEFT JOin public.users u on u.id = p.user_id").
		Joins("LEFT JOIN public.categories c ON c.id = p.category_id").
		Joins("LEFT JOIN public.menus m ON m.id = c.menu_id").
		Where("p.id = ? and p.deleted_at isnull", id).
		Order("p.id DESC").Limit(1)
	result.Scan(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil

}

func (r *postRepo) DeletePostByID(id string, db *gorm.DB) error {
	result := db.Delete(&models.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *postRepo) EditPost(post *models.Post, db *gorm.DB) error {
	result := db.Save(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *postRepo) FindPostByCategoryID(id int, name string, offset, limit int64, db *gorm.DB) ([]*models.PostHome, int64, error) {
	var posts []*models.PostHome
	var countQuery int64
	result := db.Table("public.posts as p").
		Select("p.id as post_id, u.full_name as author ,m.name as item , p.title, c.category_name as category, p.description, i.url as image_url, p.content, p.created_at").
		Joins("inner join public.categories as sub on sub.id = p.category_id").
		Joins("LEFT JOIN public.images i on i.id = p.image_id").
		Joins("LEFT JOin public.users u on u.id = p.user_id").
		Joins("LEFT JOIN public.categories c ON c.id = p.category_id").
		Joins("LEFT JOIN public.menus m ON m.id = c.menu_id").Where("p.deleted_at isnull").Order("p.id Desc")
	if name != "" {
		result = result.Where("sub.parent_name like ? and sub.deleted_at isnull", "%"+name+"%")
	} else {
		result = result.Where("p.category_id = ?", id)
	}
	queryRow := result.Count(&countQuery)
	result = result.Offset(offset).Limit(limit)
	result.Scan(&posts)
	if queryRow.Error != nil {
		return nil, countQuery, result.Error
	}
	return posts, countQuery, nil
}
