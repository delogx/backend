package models

import "gorm.io/gorm"

type StringComponent struct {
	gorm.Model
	Value string `gorm:"not null"`
}
