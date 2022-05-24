package usecase

import (
	"Hybrid-blog/constant"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	"Hybrid-blog/helpers/database"
	"Hybrid-blog/helpers/utility"
	"Hybrid-blog/models"
	model "Hybrid-blog/models"
	repository "Hybrid-blog/repositories"
	repo "Hybrid-blog/repositories/repoImlp"
)

type UserCaseUser interface {
	IsLogin(email string, password string) (*model.UserLogin, bool)
	ResetPassword(userid int, password string) bool
	GetInfoUserDetails(email string) (*model.UserDetail, error)
	EditUserDetail(userid int64, user models.UserEdit) error
	FindUserByEmail(email string) (*models.User, error)
	CreateUser(user models.User) (bool, error)
	GetAllUser(email, username, atFrom, atTo string, offset, limit int64) ([]*model.UserDetail, int64, error)
	EditUserByID(userid int64, email, username string, role string) error
	DeletebyID(id int) error
	CreateFwPassword(fw *models.Reset_password) error
	UpdateFwPassword(Fwpass *models.Reset_password, time string) error
	DeleteFwPassword(id int) error
	FindFwPassword(id int) (*model.Reset_password, error)
}

type (
	userCaseUser struct {
		db             *gorm.DB
		userrepository repository.User
	}
)

func NewuserCase() UserCaseUser {
	return &userCaseUser{
		db:             database.DB(),
		userrepository: repo.NewUserRepo(),
	}
}

func (u *userCaseUser) IsLogin(email string, password string) (*model.UserLogin, bool) {
	user, ok := u.userrepository.IsLogin(email, u.db)
	var checkpass bool
	if user != nil {
		checkpass = utility.CheckPasswordHash(password, user.Password)
		if ok != nil {
			log.Println(ok)
			checkpass = false
			return nil, checkpass
		}
	}
	return user, checkpass
}

func (u *userCaseUser) ForgotPassword(email string) bool {
	err, ok := u.userrepository.ForgotPassword(email, u.db)
	if err != nil {
		log.Println(err)
	}
	if ok == true {
		log.Println(ok)
		return true
	}
	return false
}

func (u *userCaseUser) ResetPassword(userid int, password string) bool {
	id := strconv.Itoa(userid)
	err, ok := u.userrepository.ResetPassword(id, password, u.db)
	if err != nil {
		log.Println(err)
	}
	if ok == true {
		return true
	}
	return false
}

func (u *userCaseUser) GetInfoUserDetails(email string) (*model.UserDetail, error) {
	user, err := u.userrepository.GetInfoUserDetails(email, u.db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userCaseUser) EditUserDetail(userid int64, user models.UserEdit) error {
	err := u.userrepository.EditUserDetail(userid, user, u.db)
	if err != nil {
		return err
	}
	return nil
}

func (u *userCaseUser) FindUserByEmail(email string) (*models.User, error) {
	user, err := u.userrepository.FindUserByEmail(email, u.db)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userCaseUser) CreateUser(user models.User) (bool, error) {
	ok, err := u.userrepository.CreateUser(user, u.db)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (u *userCaseUser) GetAllUser(email, username, atFrom, atTo string, offset, limit int64) ([]*model.UserDetail, int64, error) {
	user, countQuery, err := u.userrepository.GetAllUser(email, username, atFrom, atTo, offset, limit, u.db)
	if err != nil {
		return nil, -1, err
	}
	return user, countQuery, nil
}

func (u *userCaseUser) EditUserByID(userid int64, email, username string, role string) error {
	userID := fmt.Sprintf("%d", userid)
	id, _ := strconv.ParseUint(userID, 0, 32)
	var RoleUser = false
	if role == constant.DefaultRoleAdmin {
		RoleUser = true
	}
	Id := uint(id)
	var user = &models.User{
		Model: gorm.Model{
			ID: Id,
		},
		Email:    email,
		FullName: username,
		RoleUser: &RoleUser,
	}

	err := u.userrepository.EditUserByID(user, u.db)
	if err != nil {
		return err
	}
	return nil
}

func (u *userCaseUser) DeletebyID(id int) error {
	userID := strconv.Itoa(id)
	err := u.userrepository.DeleteUserByID(userID, u.db)
	if err != nil {
		return err
	}
	return nil
}

func (u *userCaseUser) CreateFwPassword(fw *models.Reset_password) error {
	err := u.userrepository.CreateFwPassword(fw, u.db)
	if err != nil {
		return err
	}
	return nil
}

func (u *userCaseUser) UpdateFwPassword(Fwpass *models.Reset_password, Time string) error {
	Fwpass.Time = Time
	now := time.Now().Local().Add(time.Hour * time.Duration(24))
	log.Println(now)
	Fwpass.Expired_at = now
	err := u.userrepository.UpdateFwPassword(Fwpass, u.db)
	if err != nil {
		return err
	}
	return nil
}

func (u *userCaseUser) DeleteFwPassword(id int) error {
	userID := strconv.Itoa(id)
	err := u.userrepository.DeleteFwPassword(userID, u.db)
	if err != nil {
		return err
	}
	return nil
}

func (u *userCaseUser) FindFwPassword(id int) (*model.Reset_password, error) {
	userID := strconv.Itoa(id)
	rs, err := u.userrepository.FindFwPassword(userID, u.db)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
