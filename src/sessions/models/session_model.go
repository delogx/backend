package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	IP              string `gorm:"nullable"`
	dashboardUserID uint   `gorm:"not null"`
}
