package models

import (
	"backend/src/types"
	"gorm.io/gorm"
)

type Component struct {
	gorm.Model
	TagName     string      `gorm:"not null"`
	Attributes  types.JSONB `gorm:"type:json;nullable"`
	TextContent string      `gorm:"nullable"`
	Children    types.JSONB `gorm:"type:json;nullable"`
	AppID       uint        `gorm:"not nullable"`
}
