package auth

import (
	"backend/src/common/models"
)

type DashboardUserService interface {
	FindOneWithPass(email string) (*models.DashboardUserWithPassword, error)
	Create(name string, email string, hashedPassword string) (*models.DashboardUserWithPassword, error)
}

type Provider struct {
	DashboardUserService DashboardUserService
}
