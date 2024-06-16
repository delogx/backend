package services

import (
	auth "backend/src/auth"
	"backend/src/auth/dtos"
	auth_utils "backend/src/auth/utils"
	"backend/src/common/models"
	"backend/src/common/types"
	"backend/src/common/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func Login(dto dtos.LoginDTO, provider auth.Provider) (string, error) {
	var dashboardUser, err = provider.DashboardUserService.FindOneWithPass(dto.Email, provider.DB)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	if ok := auth_utils.CheckPasswordHash(dashboardUser.Password, dto.Password); !ok {
		return "", fmt.Errorf("invalid credentials")
	}
	return generateAccessToken(accessTokendashboardUser{
		ID:    dashboardUser.ID,
		Email: dashboardUser.Email,
		Name:  dashboardUser.Name,
	})
}

func AdminLogin(dto dtos.LoginDTO, provider auth.Provider) (string, error) {
	var dashboardUser, err = provider.DashboardUserService.FindOneWithPass(dto.Email, provider.DB)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	if ok := auth_utils.CheckPasswordHash(dashboardUser.Password, dto.Password); !ok {
		return "", fmt.Errorf("invalid credentials")
	}
	return generateAccessToken(accessTokendashboardUser{
		ID:      dashboardUser.ID,
		Email:   dashboardUser.Email,
		Name:    dashboardUser.Name,
		IsAdmin: false,
	})
}

type accessTokendashboardUser struct {
	ID      uint
	Email   string
	Name    string
	IsAdmin bool
}

func generateAccessToken(dashboardUser accessTokendashboardUser) (string, error) {
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

func Register(dto dtos.RegisterDTO, provider auth.Provider) (*models.DashboardUserWithPassword, error) {
	hashedPassword, err := auth_utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	return provider.DashboardUserService.Create(dto.Name, dto.Email, hashedPassword, provider.DB)
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
