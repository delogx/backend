package initializers

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	r := gin.Default()
	InitModules(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	fmt.Print("port", port)
	r.Run(":" + port)
}
