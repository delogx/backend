package controllers

import (
	"backend/db"
	"backend/src/auth/dtos"
	"backend/src/auth/services"
	"backend/src/dashboard_users/models"
	"backend/src/utils"
	"fmt"
	"net/http"
	"time"

	mail_service "backend/src/mail/services"

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
	dashboardUser, err := services.Register(registerData)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	type Result struct {
		ID              uint   `json:"id"`
		Name            string `json:"name"`
		Email           string `json:"email"`
		CreatedAt       string `json:"created_at"`
		VerifiedEmailAt string `json:"verified_email_a"`
	}
	result := Result{
		ID:              dashboardUser.ID,
		Name:            dashboardUser.Name,
		Email:           dashboardUser.Email,
		CreatedAt:       dashboardUser.CreatedAt.String(),
	}
	emailVerificationToken, err := services.GenerateVerificationEmailToken(dashboardUser.ID)
	if err == nil {
		fmt.Println("email_verification token", dashboardUser.ID, dashboardUser.Email, dashboardUser.Name, emailVerificationToken)
		go mail_service.SendMail(mail_service.Mail{
			To:      dashboardUser.Email,
			From:    "admin@delogx.com",
			Subject: "Email Verification",
			Content: fmt.Sprintf("Hi %s, this email is for verifying your identity.", dashboardUser.Name),
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
	now := time.Now()
	db.DB.Select("verified_email_at").Updates(&models.DashboardUser{
		ID:              uint(requestUser.ID),
		VerifiedEmailAt: &now,
	})
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func Me(ctx *gin.Context) {
	requestUser, ok := utils.GetRequestUser(ctx)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"dashboard_user": &requestUser})
}
