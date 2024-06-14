package utils

import (
	"backend/src/common/types"

	"github.com/gin-gonic/gin"
)

func GetRequestUser(ctx *gin.Context) (*types.RequestDashboardUser, bool) {
	user, ok := ctx.Get("dashboard_user")
	if !ok {
		return nil, false
	}
	requestUser, ok := user.(types.RequestDashboardUser)
	if !ok {
		return nil, false
	}
	return &requestUser, true
}
