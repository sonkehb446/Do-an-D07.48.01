package repository

import (
	"github.com/jinzhu/gorm"

	"Hybrid-blog/models"
	model "Hybrid-blog/models"
)

type User interface {
	ForgotPassword(email string, db *gorm.DB) (error, bool)
	ResetPassword(email string, pass string, db *gorm.DB) (error, bool)
	IsLogin(email string, db *gorm.DB) (*model.UserLogin, error)
	GetInfoUserDetails(email string, db *gorm.DB) (*model.UserDetail, error)
	EditUserDetail(userid int64, user models.UserEdit, db *gorm.DB) error
	FindUserByEmail(email string, db *gorm.DB) (*models.User, error)
	CreateUser(user models.User, db *gorm.DB) (bool, error)
	GetAllUser(email, username, atFrom, atTo string, offset, limit int64, db *gorm.DB) ([]*model.UserDetail, int64, error)
	DeleteUserByID(userid string, db *gorm.DB) error
	EditUserByID(user *models.User, db *gorm.DB) error
	CreateFwPassword(fw *models.Reset_password, db *gorm.DB) error
	UpdateFwPassword(fw *models.Reset_password, db *gorm.DB) error
	DeleteFwPassword(id string, db *gorm.DB) error
	FindFwPassword(id string, db *gorm.DB) (*model.Reset_password, error)
}
