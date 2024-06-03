package routers

import "github.com/gin-gonic/gin"
import auth_router "backend/src/auth/routers"

func Setup(r *gin.RouterGroup) {
	auth_router.Setup(r.Group("auth"))
}
