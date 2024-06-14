package initializers

import (
	apps "backend/src/initializers/modules/apps"
	auth "backend/src/initializers/modules/auth"

	"github.com/gin-gonic/gin"
)

func InitModules(r *gin.Engine) {
	auth.InitModule(r)
	apps.InitModule(r)
}
