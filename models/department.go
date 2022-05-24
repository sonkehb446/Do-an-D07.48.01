package models

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	Name string
}
