package models

import "gorm.io/gorm"

type AppUser struct {
	gorm.Model
	Email    string `gorm:"not null"`
	Password string `gorm:"->:false;<-:create"`
	Name     string `gorm:"not null"`
	IsAdmin  bool   `gorm:"default false"`
	AppID    uint   `gorm:"not null"`
}
