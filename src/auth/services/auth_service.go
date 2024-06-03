package services

import (
	"backend/db"
	"backend/src/auth/dtos"
	auth_utils "backend/src/auth/utils"
	"backend/src/dashboard_users/models"
	"backend/src/types"
	"backend/src/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func Login(dto dtos.LoginDTO) (string, error) {
	var dashboardUser, err = GetdashboardUserWithPass(dto)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	if ok := auth_utils.CheckPasswordHash(dashboardUser.Password, dto.Password); !ok {
		return "", fmt.Errorf("invalid credentials")
	}
	return GenerateAccessToken(AccessTokendashboardUser{
		ID:    dashboardUser.ID,
		Email: dashboardUser.Email,
		Name:  dashboardUser.Name,
	})
}

func AdminLogin(dto dtos.LoginDTO) (string, error) {
	var dashboardUser, err = GetdashboardUserWithPass(dto)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	if ok := auth_utils.CheckPasswordHash(dashboardUser.Password, dto.Password); !ok {
		return "", fmt.Errorf("invalid credentials")
	}
	return GenerateAccessToken(AccessTokendashboardUser{
		ID:      dashboardUser.ID,
		Email:   dashboardUser.Email,
		Name:    dashboardUser.Name,
		IsAdmin: false,
	})
}

func GetdashboardUserWithPass(dto dtos.LoginDTO) (*models.DashboardUserWithPassword, error) {
	var user models.DashboardUserWithPassword
	db.DB.Where("email = ? OR phone_number = ?", dto.Username, dto.Username).First(&user)
	if user.ID == 0 {
		return nil, fmt.Errorf("app user not found")
	}
	return &user, nil
}

type AccessTokendashboardUser struct {
	ID      uint
	Email   string
	Name    string
	IsAdmin bool
}

func GenerateAccessToken(dashboardUser AccessTokendashboardUser) (string, error) {
	requestdashboardUser := types.RequestDashboardUser{
		ID:      float64(dashboardUser.ID),
		Email:   dashboardUser.Email,
		Name:    dashboardUser.Name,
		IsAdmin: dashboardUser.IsAdmin,
	}
	claims := jwt.MapClaims{
		"dashboard_user": requestdashboardUser,
	}
	accessToken, err := auth_utils.GenerateJWT(claims)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func Register(dto dtos.RegisterDTO) (*models.DashboardUserWithPassword, error) {
	hashedPassword, err := auth_utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	dashboardUser := models.DashboardUserWithPassword{
		Name:     dto.Name,
		Password: hashedPassword,
		IsAdmin:  false,
	}
	result := db.DB.Create(&dashboardUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dashboardUser, nil
}

func GenerateVerificationEmailToken(dashboardUserId uint) (string, error) {
	claims := jwt.MapClaims{
		"isEmailVerificationToken": true,
		"userId":                   dashboardUserId,
	}
	token, err := auth_utils.GenerateJWT(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyVerificationToken(dashboardUserId float64, verificationToken string) bool {
	claims, ok := utils.ParseJWT(verificationToken)
	if !ok || claims == nil {
		return false
	}
	isVerificationToken, ok := (*claims)["isEmailVerificationToken"].(bool)
	if !ok || !isVerificationToken {
		return false
	}
	tokenUserId, ok := (*claims)["userId"].(float64)
	if !ok || tokenUserId != dashboardUserId {
		return false
	}
	return true
}
