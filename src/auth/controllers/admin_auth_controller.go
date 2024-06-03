package controllers

import (
	"backend/src/auth/dtos"
	"backend/src/auth/services"
	"backend/src/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminLogin(ctx *gin.Context) {
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

