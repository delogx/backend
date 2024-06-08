package services

import (
	"backend/db"
	"backend/src/common/models"
	"fmt"
)

type Service struct{}

func (sc *Service) FindOneWithPass(email string) (*models.DashboardUserWithPassword, error) {
	var user models.DashboardUserWithPassword
	db.DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return nil, fmt.Errorf("dashboard user not found")
	}
	return &user, nil
}

func (sc *Service) Create(name string, email string, hashedPassword string) (*models.DashboardUserWithPassword, error) {
	dashboardUser := models.DashboardUserWithPassword{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		IsAdmin:  false,
	}
	result := db.DB.Create(&dashboardUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dashboardUser, nil
}
