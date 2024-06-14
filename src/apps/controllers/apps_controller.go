package controllers

import (
	"backend/src/apps"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context, provider apps.Provider) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello apps service"})
}