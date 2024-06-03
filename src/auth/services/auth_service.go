package services

import (
	"backend/db"
	"backend/src/app_users/models"
	"backend/src/auth/dtos"
	auth_utils "backend/src/auth/utils"
	"backend/src/types"
	"backend/src/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func Login(dto dtos.LoginDTO) (string, error) {
	var appUser, err = GetAppUserWithPass(dto)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	if ok := auth_utils.CheckPasswordHash(appUser.Password, dto.Password); !ok {
		return "", fmt.Errorf("invalid credentials")
	}
	return GenerateAccessToken(AccessTokenAppUser{
		ID:    appUser.ID,
		Email: appUser.Email,
		Name:  appUser.Name,
	})
}

func AdminLogin(dto dtos.LoginDTO) (string, error) {
	var appUser, err = GetAppUserWithPass(dto)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	if ok := auth_utils.CheckPasswordHash(appUser.Password, dto.Password); !ok {
		return "", fmt.Errorf("invalid credentials")
	}
	return GenerateAccessToken(AccessTokenAppUser{
		ID:      appUser.ID,
		Email:   appUser.Email,
		Name:    appUser.Name,
		IsAdmin: false,
	})
}

func GetAppUserWithPass(dto dtos.LoginDTO) (*models.AppUserWithPassword, error) {
	var user models.AppUserWithPassword
	db.DB.Where("email = ? OR phone_number = ?", dto.Username, dto.Username).First(&user)
	if user.ID == 0 {
		return nil, fmt.Errorf("app user not found")
	}
	return &user, nil
}

type AccessTokenAppUser struct {
	ID      uint
	Email   string
	Name    string
	IsAdmin bool
}

func GenerateAccessToken(appUser AccessTokenAppUser) (string, error) {
	requestAppUser := types.RequestAppUser{
		ID:      float64(appUser.ID),
		Email:   appUser.Email,
		Name:    appUser.Name,
		IsAdmin: appUser.IsAdmin,
	}
	claims := jwt.MapClaims{
		"app_user": requestAppUser,
	}
	accessToken, err := auth_utils.GenerateJWT(claims)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func Register(dto dtos.RegisterDTO) (*models.AppUserWithPassword, error) {
	hashedPassword, err := auth_utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	appUser := models.AppUserWithPassword{
		Name:     dto.Name,
		Password: hashedPassword,
		IsAdmin:  false,
	}
	result := db.DB.Create(&appUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &appUser, nil
}

func GenerateVerificationEmailToken(appUserId uint) (string, error) {
	claims := jwt.MapClaims{
		"isEmailVerificationToken": true,
		"userId":                  appUserId,
	}
	token, err := auth_utils.GenerateJWT(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyVerificationToken(appUserId float64, verificationToken string) bool {
	claims, ok := utils.ParseJWT(verificationToken)
	if !ok || claims == nil {
		return false
	}
	isVerificationToken, ok := (*claims)["isEmailVerificationToken"].(bool)
	if !ok || !isVerificationToken {
		return false
	}
	tokenUserId, ok := (*claims)["userId"].(float64)
	if !ok || tokenUserId != appUserId {
		return false
	}
	return true
}
