package models

import (
	"time"
)

type AppAdmin struct {
	AppID           uint       `gorm:"primaryKey;not null"`
	DashboardUserID uint       `gorm:"primaryKey;not null"`
	CreatedAt       time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time  `gorm:"default:current_timestamp on update current_timestamp"`
	DeletedAt       *time.Time `gorm:"nullable"`
}
