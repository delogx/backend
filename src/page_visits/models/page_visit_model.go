package models

import "gorm.io/gorm"

type PageVisit struct {
	gorm.Model
	SessionID uint   `gorm:"not null"`
	URL       string `gorm:"not null"`
}
