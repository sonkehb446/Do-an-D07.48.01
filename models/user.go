package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	User struct {
		gorm.Model
		Email        string
		Password     string
		RoleUser     *bool
		BranchID     *int
		DepartmentID *int
		PositionID   *int
		ImageID      *int
		FullName     string
		Phone        *string
		Birthday     *string
		Description  *string
		Facebook     *string
		CreateBy     *int64
	}
	UserLogin struct {
		UserID   int
		FullName string
		Email    string
		Password string
		Url      string
		RoleUser *bool
	}

	UserDetail struct {
		UserID      int
		FullName    string
		Email       string
		ImageID     int
		RoleUser    *bool
		ImageUrl    string
		Branch      string
		Department  string
		Position    string
		Phone       string
		Birthday    string
		Description string
		Facebook    string
		CreatedAt   time.Time
	}

	UserEdit struct {
		FullName     string
		Email        string
		RoleUser     *bool
		ImageID      int64
		BranchID     int64
		DepartmentID int64
		PositionID   int64
		Phone        string
		Birthday     string
		Description  string
		Facebook     string
	}
)
