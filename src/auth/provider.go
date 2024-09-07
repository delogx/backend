package auth

import (
	"backend/src/common/models"
	"backend/src/common/types"
)

type DashboardUserService interface {
	FindOneWithPass(email string, db types.DB) (*models.DashboardUserWithPassword, error)
	FindOne(db types.DB, user *models.DashboardUser) (*models.DashboardUser, error)
	Create(name string, email string, hashedPassword string, db types.DB) (*models.DashboardUserWithPassword, error)
}

type Provider struct {
	DashboardUserService DashboardUserService
	DB                   types.DB
}
