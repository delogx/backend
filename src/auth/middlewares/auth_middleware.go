package middlewares

import (
	"backend/src/common/types"
	"backend/src/common/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		access_token := strings.TrimPrefix(strings.TrimSpace(ctx.GetHeader("Authorization")), "Bearer ")
		if len(access_token) == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := utils.ParseJWT(access_token)
		if !ok || claims == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user := (*claims)["dashboard_user"].(map[string]any)
		id, ok := user["ID"].(float64)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		name, ok := user["Name"].(string)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		email, ok := user["Email"].(string)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		is_admin, ok := user["IsAdmin"].(bool)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		requestUser := types.RequestDashboardUser{
			ID:      id,
			Email:   email,
			Name:    name,
			IsAdmin: is_admin,
		}
		ctx.Set("dashboard_user", requestUser)
		ctx.Next()
	}
}
