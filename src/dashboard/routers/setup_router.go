package routers

import "github.com/gin-gonic/gin"
import users_router "backend/src/dashboard/users/routers"

func Setup(r *gin.RouterGroup) {
	users_router.Setup(r.Group("users"))
}
