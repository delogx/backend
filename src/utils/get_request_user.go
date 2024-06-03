package utils

import (
	"backend/src/types"

	"github.com/gin-gonic/gin"
)

func GetRequestUser(ctx *gin.Context) (*types.RequestAppUser, bool) {
	user, ok := ctx.Get("app_user")
	if !ok {
		return nil, false
	}
	requestUser, ok := user.(types.RequestAppUser)
	if !ok {
		return nil, false
	}
	return &requestUser, true
}
