package auth

import (
	"backend/src/common/models"
	"backend/src/common/types"
)

type DashboardUserService interface {
	FindOneWithPass(email string) (*models.DashboardUserWithPassword, error)
	FindOne(email string, db types.DB) (*models.DashboardUser, error)
	Create(name string, email string, hashedPassword string) (*models.DashboardUserWithPassword, error)
}

type Provider struct {
	DashboardUserService DashboardUserService
}
