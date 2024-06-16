package services

import (
	"backend/src/common/models"
	"backend/src/common/types"
	"fmt"
)

type Service struct{}

func (*Service) Create(Name string, HostName string, UserID uint, db types.DB) (*models.App, error) {
	app := models.App{
		Name:     Name,
		HostName: HostName,
	}
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, nil
	}
	if err := tx.Create(&app).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Model(&app).Association("Admins").Append(&models.AppAdmin{
		AppID:           app.ID,
		DashboardUserID: UserID,
	}); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &app, nil
}

func (*Service) FindOne(db types.DB) (*models.App, error) {
	var app models.App
	db.First(&app)
	if app.ID == 0 {
		return nil, fmt.Errorf("app not found")
	}
	return &app, nil
}
