package middlewares

import (
	"backend/src/auth"
	"backend/src/common/models"
	"backend/src/common/types"
	"backend/src/common/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Make sure AuthMiddleware is called before this
func VerifiedDashboardUserMiddleware(dashboardUserService auth.DashboardUserService, db types.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestUser, ok := utils.GetRequestUser(ctx)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, err := dashboardUserService.FindOne(db, &models.DashboardUser{Email: requestUser.Email})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		if user.VerifiedEmailAt == nil {
			ctx.AbortWithStatusJSON(403, map[string]string{"message": "dashboard_user_email_not_verified"})
		}
		ctx.Next()
	}
}
