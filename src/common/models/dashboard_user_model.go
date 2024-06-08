package models

import (
	"time"

	"gorm.io/gorm"
)

type DashboardUser struct {
	gorm.Model
	ID              uint       `gorm:"primaryKey"`
	Email           string     `gorm:"not null"`
	Password        string     `gorm:"->:false;<-:create"`
	Name            string     `gorm:"not null"`
	IsAdmin         bool       `gorm:"default false"`
	CreatedAt       time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time  `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt       *time.Time `gorm:"nullable"`
	VerifiedEmailAt *time.Time `gorm:"nullable"`

	Apps []*App `gorm:"many2many:app_dashboard_users"`
}

type DashboardUserWithPassword struct {
	gorm.Model
	ID              uint       `gorm:"primaryKey"`
	Email           string     `gorm:"not null"`
	Password        string     `gorm:"not null"`
	Name            string     `gorm:"not null"`
	IsAdmin         bool       `gorm:"default false"`
	CreatedAt       time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time  `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt       *time.Time `gorm:"nullable"`
	VerifiedEmailAt *time.Time `gorm:"nullable"`
}

func (table DashboardUserWithPassword) TableName() string {
	return "dashboard_users"
}
