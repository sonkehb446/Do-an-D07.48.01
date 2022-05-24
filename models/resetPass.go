package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	Reset_password struct {
		gorm.Model
		UserID     int64
		Token      string
		Time       string
		Expired_at time.Time
	}
)
