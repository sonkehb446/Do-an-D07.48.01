package models

import "github.com/jinzhu/gorm"

type Position struct {
	gorm.Model
	Name string
}
