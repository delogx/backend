package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	IP            string `gorm:"nullable"`
	AppID         uint   `gorm:"not null"`
	UserIDFromApp string `gorm:"nullable"`
}
