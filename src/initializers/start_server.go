package initializers

import (
	"backend/src/router"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	router.Setup(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	fmt.Print("port", port)
	r.Run(":" + port)
}
