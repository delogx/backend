package routers

import (
	"backend/src/ws/controllers"
	"backend/src/ws/types"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.RouterGroup) {
	rm := types.NewRoomManager()
	r.GET("/:userId", func(ctx *gin.Context) {
		controllers.WSHandler(ctx, rm)
	})
}
