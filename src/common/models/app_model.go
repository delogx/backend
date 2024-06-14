package models

import "gorm.io/gorm"

type App struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	HostName string `gorm:"not null"`

	DashboardUsers []*DashboardUser `gorm:"many2many:app_dashboard_users"`
}
