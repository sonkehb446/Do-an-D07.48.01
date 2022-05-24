package usecase

import (
	"strconv"

	"github.com/jinzhu/gorm"

	"Hybrid-blog/helpers/database"
	"Hybrid-blog/models"
	repository "Hybrid-blog/repositories"
	repo "Hybrid-blog/repositories/repoImlp"
)

type UserCasePost interface {
	CreatePost(Post *models.Post) error
	GetPostByType(Type int, offset, limit int) ([]*models.PostHome, int64, error)
	FindPostByID(id int) (*models.PostHome, error)
	DeletePostByID(id int) error
	EditPost(*models.Post) error
	GetAllPostByID(IDUser, Title, Description, Create_By, Create_from, Create_To string, offset, limit int) ([]*models.PostHome, int64, error)
	FindPostByCategoryID(id int, name string, offset, limit int64) ([]*models.PostHome, int64, error)
}

type (
	userCasePost struct {
		db         *gorm.DB
		repository repository.PostInterface
	}
)

func NewPostUseCase() UserCasePost {
	return &userCasePost{
		db:         database.DB(),
		repository: repo.NewPostRepo(),
	}
}

func (u *userCasePost) CreatePost(Post *models.Post) error {
	if err := u.repository.CreatePost(Post, u.db); err != nil {
		return err
	}
	return nil
}

func (u *userCasePost) EditPost(post *models.Post) error {
	if err := u.repository.EditPost(post, u.db); err != nil {
		return err
	}
	return nil
}

func (u *userCasePost) GetAllPostByID(IDUser, Title, Description, Create_By, Create_from, Create_To string, offset, limit int) ([]*models.PostHome, int64, error) {
	offSet := strconv.Itoa(offset)
	limit1 := strconv.Itoa(limit)
	post, countQuery, err := u.repository.GetAllPostByID(IDUser, Title, Description, Create_By, Create_from, Create_To, offSet, limit1, u.db)
	if err != nil {
		return nil, -1, err
	}
	return post, countQuery, nil
}

func (u *userCasePost) GetPostByType(Type int, offset, limit int) ([]*models.PostHome, int64, error) {
	typeI := strconv.Itoa(Type)
	offSet := strconv.Itoa(offset)
	limit1 := strconv.Itoa(limit)
	posts, countQuery, err := u.repository.GetLimitPostByType(typeI, u.db, offSet, limit1)
	if err != nil {
		return nil, -1, err
	}
	return posts, countQuery, nil
}

func (u *userCasePost) FindPostByID(id int) (*models.PostHome, error) {
	ID := strconv.Itoa(id)
	post, err := u.repository.FindPostByID(ID, u.db)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *userCasePost) DeletePostByID(id int) error {
	ID := strconv.Itoa(id)
	err := u.repository.DeletePostByID(ID, u.db)
	if err != nil {
		return err
	}
	return nil
}

func (u *userCasePost) FindPostByCategoryID(id int, name string, offset, limit int64) ([]*models.PostHome, int64, error) {
	post, count, err := u.repository.FindPostByCategoryID(id, name, offset, limit, u.db)
	if err != nil {
		return post, count, err
	}
	return post, count, err
}
