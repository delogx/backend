package controllers

import (
	"backend/src/apps"
	"backend/src/apps/dtos"
	"backend/src/apps/services"
	"backend/src/common/models"
	"backend/src/common/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context, provider apps.Provider) {
	requestUser, ok := utils.GetRequestUser(ctx)
	appService := &services.Service{}
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var dto dtos.CreateAppDto
	if ok := utils.ValidateRequestBody(ctx, &dto); !ok {
		return
	}
	if _, err := appService.FindOne(provider.DB, &models.App{HostName: dto.HostName}); err == nil {
		ctx.AbortWithStatusJSON(422, gin.H{"errors": gin.H{"host_name": "host_name already exists"}})
		return
	}
	if app, err := appService.Create(dto.Name, dto.HostName, uint(requestUser.ID), provider.DB); err == nil {
		type Result struct {
			ID       uint   `json:"id"`
			HostName string `json:"host_name"`
			Name     string `json:"name"`
		}
		result := Result{
			ID:       app.ID,
			HostName: app.HostName,
			Name:     app.Name,
		}
		ctx.JSON(http.StatusCreated, gin.H{"data": &result})
	} else {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"message": err})
	}
}
