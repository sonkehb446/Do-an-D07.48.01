package repoimlp

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"Hybrid-blog/models"
	model "Hybrid-blog/models"
	repo "Hybrid-blog/repositories"
)

type userRepo struct{}

func NewUserRepo() repo.User {
	return &userRepo{}
}

func (u *userRepo) SelectUsers(db *gorm.DB) (*model.User, error) {
	var user *model.User
	result := db.Raw("SELECT * FROM public.users WHERE  public.users.deleted_at ISNULL").Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *userRepo) IsLogin(email string, db *gorm.DB) (*model.UserLogin, error) {
	var list []*model.UserLogin
	var result *gorm.DB
	var user *model.UserLogin
	var err error
	result = db.Table("public.users").
		Select(" users.id as user_id ,users.email, users.password , users.full_name , images.url ,users.image_id , images.id , users.role_user").
		Joins("LEFT JOIN public.images ON images.id = users.image_id").
		Where("users.email = ?  and users.deleted_at ISNULL", email)
	result.Scan(&list)
	if len(list) > 0 {
		user = list[0]
		err = nil
	}
	if result.Error != nil {
		user = nil
		err = result.Error
	}
	return user, err
}

func (u *userRepo) ForgotPassword(email string, db *gorm.DB) (error, bool) {
	var err error
	var checkFounds = false
	var list []*model.User

	var result *gorm.DB
	result = db.Table("public.users").Select("email").Where("email = ?  and users.deleted_at ISNULL", email).Limit(1).Scan(&list)

	if len(list) > 0 {
		checkFounds = true
		err = nil
	} else if result.Error != nil {
		err = result.Error
	}
	return err, checkFounds
}
func (u *userRepo) ResetPassword(id string, pass string, db *gorm.DB) (error, bool) {
	var checkBool = true
	rback := db.Begin()
	err := db.Table("public.users").Where("id = ?", id).Update("password", pass)
	if err.Error != nil {
		checkBool = false
		log.Println("RoolBack")
		rback.Rollback()
		return err.Error, checkBool
	}
	return nil, checkBool
}

func (u *userRepo) GetInfoUserDetails(email string, db *gorm.DB) (*model.UserDetail, error) {
	var list []*model.UserDetail
	var result *gorm.DB
	var user *model.UserDetail
	var err error
	result = db.Table("public.users as u").
		Select(" u.id as user_id, u.email as email, u.full_name as full_name, i.url  as image_url, u.image_id, b.name as branch, d.name as department, po.name as position, u.phone  as  phone, u.birthday  as birthday, u.description as description , u.facebook  as facebook,u.role_user").
		Joins("LEFT JOIN public.departments as d on u.department_id = d.id").
		Joins("LEFT JOIN public.positions as po on u.position_id = po.id").
		Joins("LEFT JOIN public.images as i on u.image_id = i.id").
		Joins("LEFT JOIN public.branches as b on u.branch_id = b.id").
		Where("u.email = ?  and u.deleted_at ISNULL", email)
	result.Scan(&list)
	if len(list) > 0 {
		user = list[0]
		err = nil
	} else {
		result = db.Table("public.users").
			Select("users.id, users.email, users.password").
			Where("users.email = ?  and users.deleted_at ISNULL", email)
		result.Scan(&list)
		if len(list) > 0 {
			user = list[0]
			err = nil
		}
	}
	if result.Error != nil {
		user = nil
		err = result.Error
	}
	return user, err
}

func (u *userRepo) EditUserDetail(userid int64, user models.UserEdit, db *gorm.DB) error {
	result := db.Exec("UPDATE public.users SET  email=?,branch_id=?,department_id=?,position_id=?,image_id=?,full_name=?,phone=?,birthday=?,description=?,facebook=? Where ID = ?",
		user.Email, user.BranchID, user.DepartmentID, user.PositionID, user.ImageID, user.FullName, user.Phone, user.Birthday, user.Description, user.Facebook, userid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepo) FindUserByEmail(email string, db *gorm.DB) (*models.User, error) {
	var user []*model.User
	result := db.Table("public.users").
		Select("users.id, users.email, users.password").
		Where("users.email = ? and users.deleted_at ISNULL", email).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(user) > 0 {
		return user[0], nil
	}
	return nil, nil
}

func (u *userRepo) CreateUser(user models.User, db *gorm.DB) (bool, error) {
	rx := db.Begin()
	if err := rx.Create(&user).Error; err != nil {
		rx.Rollback()
		return false, err
	}
	rx.Commit()
	return true, nil
}

///////////////////////////////////////////////////
func (u *userRepo) GetAllUser(email, username, atFrom, atTo string, offset, limit int64, db *gorm.DB) ([]*model.UserDetail, int64, error) {
	var users []*model.UserDetail
	var countQuery int64
	var result = db.Table("public.users as u").
		Select("u.id as user_id, u.email as email, u.full_name as full_name, u.created_at ,u.role_user as role_user")
	if username != "" {
		result = result.Where("u.full_name like ?", "%"+username+"%")
	}

	if email != "" {
		result = result.Where("u.email like ?", "%"+email+"%")
	}
	if atFrom != "" {
		timefrom := fmt.Sprintf("%s 00:00:01", atFrom)
		result = result.Where("u.created_at >= ?  ", timefrom)
	}
	if atTo != "" {
		timeto := fmt.Sprintf("%s 23:59:59", atTo)
		result = result.Where("u.created_at <=  ?", timeto)
	}
	queryRow := result.Count(&countQuery)
	result = result.Where("u.deleted_at ISNULL")
	result.Order("u.id  Desc").Offset(offset).Limit(limit).Scan(&users)
	if queryRow.Error != nil {
		return nil, -1, result.Error
	}
	log.Println("Tong", countQuery)
	return users, countQuery, nil
}

func (u *userRepo) EditUserByID(user *models.User, db *gorm.DB) error {
	result := db.Model(&user).Update(models.User{FullName: user.FullName, Email: user.Email, RoleUser: user.RoleUser})
	//Updates(User{Name: "hello", Age: 18, Active: false})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepo) DeleteUserByID(userid string, db *gorm.DB) error {
	result := db.Delete(&models.User{}, userid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepo) CreateFwPassword(fw *models.Reset_password, db *gorm.DB) error {
	result := db.Create(fw)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (u *userRepo) UpdateFwPassword(fw *models.Reset_password, db *gorm.DB) error {
	result := db.Save(&fw)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *userRepo) DeleteFwPassword(id string, db *gorm.DB) error {
	result := db.Delete(&models.Reset_password{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepo) FindFwPassword(email string, db *gorm.DB) (*model.Reset_password, error) {
	var model []*model.Reset_password
	result := db.Where("user_id = ?", email).First(&models.Reset_password{}).Scan(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(model) > 0 {
		return model[0], nil
	}
	return nil, nil
}
