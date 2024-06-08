package initializers

import (
	auth "backend/src/auth"
	authModule "backend/src/auth/module"
	DashboardUsersService "backend/src/dashboard_users/services"

	"github.com/gin-gonic/gin"
)

func InitModules(r *gin.Engine) {
	InitAuthModule(r)
}

func InitAuthModule(r *gin.Engine) {
	authModule.Init(r.Group("dashboard/auth"), auth.Provider{
		DashboardUserService: &DashboardUsersService.Service{},
	})
}
