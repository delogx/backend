package controllers

import (
	"backend/db"
	"backend/src/auth/dtos"
	"backend/src/auth/services"
	mail_service "backend/src/mail/services"
	"backend/src/app_users/models"
	"backend/src/utils"
	"fmt"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var loginData dtos.LoginDTO
	if ok := utils.ValidateRequestBody(ctx, &loginData); !ok {
		return
	}
	accessToken, err := services.Login(loginData)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func Register(ctx *gin.Context) {
	var registerData dtos.RegisterDTO
	if ok := utils.ValidateRequestBody(ctx, &registerData); !ok {
		return
	}
	appUser, err := services.Register(registerData)
	if err != nil {
		ctx.JSON(400, "email or phone already exists")
		return
	}
	type Result struct {
		ID              uint   `json:"id"`
		Name            string `json:"name"`
		Email           string `json:"email"`
		CreatedAt       string `json:"created_at"`
		VerifiedEmailAt string `json:"verified_email_at"`
	}
	result := Result{
		ID:              appUser.ID,
		Name:            appUser.Name,
		Email:           appUser.Email,
		CreatedAt:       appUser.CreatedAt.String(),
		VerifiedEmailAt: appUser.VerifiedEmailAt.String(),
	}
	emailVerificationToken, err := services.GenerateVerificationEmailToken(appUser.ID)
	if err == nil {
		fmt.Println("email_verification token", appUser.ID, appUser.Email, appUser.Name, emailVerificationToken)
		go mail_service.SendMail(mail_service.Mail{
			To:      appUser.Email,
			From:    "admin@delogx.com",
			Subject: "Email Verification",
			Content: fmt.Sprintf("Hi %s, this email is for verifying your identity.", appUser.Name),
		})
	}
	ctx.JSON(http.StatusOK, &result)
}

func VerifyEmail(ctx *gin.Context) {
	var data dtos.VerifyEmailDto
	if ok := utils.ValidateRequestBody(ctx, &data); !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	requestUser, ok := utils.GetRequestUser(ctx)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if ok := services.VerifyVerificationToken(requestUser.ID, data.VerificationToken); !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	db.DB.Select("verified_email_at").Updates(&models.AppUser{
		ID:              uint(requestUser.ID),
		VerifiedEmailAt: time.Now(),
	})
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func Me(ctx *gin.Context) {
	requestUser, ok := utils.GetRequestUser(ctx)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"app_user": &requestUser})
}
