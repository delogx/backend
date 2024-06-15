package routers

import (
	module "backend/src/apps"
	"backend/src/apps/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.RouterGroup, provider module.Provider) {
	r.POST(
		"",
		provider.AuthService.AuthMiddleware(),
		provider.AuthService.VerifiedDashboardUserMiddleware(provider.DB),
		func(ctx *gin.Context) { controllers.Create(ctx, provider) },
	)
}
