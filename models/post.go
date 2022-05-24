package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	Post struct {
		gorm.Model
		UserID      int64
		Title       string
		Description string
		CategoryID  int64
		ImageID     *int64
		Content     string
	}
	PostHome struct {
		PostID      int64
		Title       string
		Author      string
		UserID      *int64
		Item        string
		Category    string
		CategoryID  int64
		Description string
		ImageUrl    string
		ImageID     *int64
		Content     string
		Created_at  time.Time
	}
)
