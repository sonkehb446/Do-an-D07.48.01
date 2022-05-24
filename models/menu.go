package models

import "github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	Name string
}
