package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey"`
	FirstName       string    `gorm:"not null"`
	LastName        string    `gorm:"not null"`
	Email           string    `gorm:"not null;unique"`
	Password        string    `gorm:"-"`
	PhoneNumber     string    `gorm:"not null"`
	IsAdmin         bool      `gorm:"default:false"`
	VerifiedEmailAt time.Time `gorm:"default:null"`
	CreatedAt       time.Time `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt       time.Time `gorm:"default:null"`
}

type UserWithPassword struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey"`
	FirstName       string    `gorm:"not null"`
	LastName        string    `gorm:"not null"`
	Email           string    `gorm:"not null;unique"`
	Password        string    `gorm:"default:null"`
	PhoneNumber     string    `gorm:"not null"`
	IsAdmin         bool      `gorm:"default:false"`
	VerifiedEmailAt time.Time `gorm:"default:null"`
	CreatedAt       time.Time `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt       time.Time `gorm:"default:null"`
}

func (table UserWithPassword) TableName() string {
	return "users"
}
