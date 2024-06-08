package module

import (
	"backend/src/auth"
	"backend/src/auth/routers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup, provider auth.Provider) {
	routers.Setup(r, provider)
}
