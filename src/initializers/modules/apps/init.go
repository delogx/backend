package init

import (
	"backend/db"
	apps "backend/src/apps"
	appsModule "backend/src/apps/module"
	"backend/src/common/types"

	authMiddlewares "backend/src/auth/middlewares"

	dashboardUserService "backend/src/dashboard_users/services"

	"github.com/gin-gonic/gin"
)

type AuthService struct{}

func (AuthService) AuthMiddleware() types.Middleware {
	return authMiddlewares.AuthMiddleware()
}

func (AuthService) VerifiedDashboardUserMiddleware(db types.DB) types.Middleware {
	return authMiddlewares.VerifiedDashboardUserMiddleware(&dashboardUserService.Service{}, db)
}

func InitModule(r *gin.Engine) {
	appsModule.Init(r.Group("dashboard/apps"), apps.Provider{
		AuthService: AuthService{},
		DB:          db.DB,
	})
}
