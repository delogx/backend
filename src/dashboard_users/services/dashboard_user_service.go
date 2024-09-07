package services

import (
	"backend/src/common/models"
	"backend/src/common/types"
	"fmt"
)

type Service struct{}

func (sc *Service) FindOneWithPass(email string, db types.DB) (*models.DashboardUserWithPassword, error) {
	var user models.DashboardUserWithPassword
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return nil, fmt.Errorf("dashboard user not found")
	}
	return &user, nil
}

func (sc *Service) FindOne(db types.DB, user *models.DashboardUser) (*models.DashboardUser, error) {
	db.Where(user).First(&user)
	if user.ID == 0 {
		return nil, fmt.Errorf("dashboard user not found")
	}
	return user, nil
}

func (sc *Service) Create(name string, email string, hashedPassword string, db types.DB) (*models.DashboardUserWithPassword, error) {
	dashboardUser := models.DashboardUserWithPassword{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		IsAdmin:  false,
	}
	result := db.Create(&dashboardUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dashboardUser, nil
}
