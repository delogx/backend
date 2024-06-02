package models

import "gorm.io/gorm"

type App struct {
	gorm.Model
	HostName string `gorm:not null`
}
