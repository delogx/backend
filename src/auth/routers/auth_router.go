package routers

import (
	module "backend/src/auth"
	"backend/src/auth/controllers"
	"backend/src/auth/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.RouterGroup, provider module.Provider) {
	r.POST("login", func(ctx *gin.Context) { controllers.Login(ctx, provider) })
	r.POST("register", middlewares.DuplicateUserMiddleware(), func(ctx *gin.Context) { controllers.Register(ctx, provider) })
	r.POST("verify-email", middlewares.AuthMiddleware(), controllers.VerifyEmail)
	r.GET("me", middlewares.AuthMiddleware(), controllers.Me)
	r.POST("admin/login", func(ctx *gin.Context) { controllers.AdminLogin(ctx, provider) })
}
