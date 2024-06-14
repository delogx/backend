package init

import (
	apps "backend/src/apps"
	appsModule "backend/src/apps/module"
	"backend/src/common/types"

	authMiddlewares "backend/src/auth/middlewares"

	"github.com/gin-gonic/gin"
)

type AuthService struct{}

func (AuthService) AuthMiddleware() types.Middleware {
	return authMiddlewares.AuthMiddleware()
}

func (AuthService) VerifiedDashboardUserMiddleware() types.Middleware {
	return authMiddlewares.AuthMiddleware()
}

func InitModule(r *gin.Engine) {
	appsModule.Init(r.Group("dashboard/apps"), apps.Provider{
		AuthService: AuthService{},
	})
}
