package init

import (
	"backend/db"
	auth "backend/src/auth"
	authModule "backend/src/auth/module"
	DashboardUsersService "backend/src/dashboard_users/services"

	"github.com/gin-gonic/gin"
)

func InitModule(r *gin.Engine) {
	authModule.Init(r.Group("dashboard/auth"), auth.Provider{
		DashboardUserService: &DashboardUsersService.Service{},
		DB:                   db.DB,
	})
}
