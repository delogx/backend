package module

import (
	"backend/src/apps"
	"backend/src/apps/routers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup, provider apps.Provider) {
	routers.Setup(r, provider)
}
