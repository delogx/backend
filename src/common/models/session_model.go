package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	IP              string `gorm:"nullable"`
	DashboardUserID uint   `gorm:"not null"`
}
