package models

import "github.com/jinzhu/gorm"

type (
	Category struct {
		gorm.Model
		MenuID       int
		CategoryName string
		ParentName   *string
		Description  string
		ParentID     *int64
	}
)
