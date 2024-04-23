package router

import "github.com/gin-gonic/gin"
import ws_router "backend/src/ws/routers"

func Setup(r *gin.Engine) {
	ws_router.Setup(r.Group("ws"))
}
