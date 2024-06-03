package routers

import (
	"backend/src/auth/controllers"
	"backend/src/auth/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.RouterGroup) {
	r.POST("login", controllers.Login)
	r.POST("register", middlewares.DuplicateUserMiddleware(), controllers.Register)
	r.POST("verify-email", middlewares.AuthMiddleware(), controllers.VerifyEmail)
	r.GET("me", middlewares.AuthMiddleware(), controllers.Me)
	r.POST("admin/login", controllers.AdminLogin)
}
