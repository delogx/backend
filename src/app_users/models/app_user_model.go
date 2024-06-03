package models

import (
	"time"

	"gorm.io/gorm"
)

type AppUser struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey"`
	Email           string    `gorm:"not null"`
	Password        string    `gorm:"->:false;<-:create"`
	Name            string    `gorm:"not null"`
	IsAdmin         bool      `gorm:"default false"`
	AppID           uint      `gorm:"nullable"`
	CreatedAt       time.Time `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time `gorm:"default:current_timestamp on update current_timestamp"`
	VerifiedEmailAt time.Time `gorm:"nullable"`
}

type AppUserWithPassword struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey"`
	Email           string    `gorm:"not null"`
	Password        string    `gorm:"not null"`
	Name            string    `gorm:"not null"`
	IsAdmin         bool      `gorm:"default false"`
	AppID           uint      `gorm:"nullable"`
	CreatedAt       time.Time `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time `gorm:"default:current_timestamp on update current_timestamp"`
	VerifiedEmailAt time.Time `gorm:"nullable"`
}

func (table AppUserWithPassword) TableName() string {
	return "users"
}
