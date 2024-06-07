package middlewares

import (
	"backend/db"
	"backend/src/auth/dtos"
	"backend/src/dashboard_users/models"
	"backend/src/utils"

	"github.com/gin-gonic/gin"
)

type UserDTO struct {
	Email string `json:"email" validate:"required,email"`
}

func DuplicateUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userData dtos.RegisterDTO
		if ok := utils.ValidateRequestBody(c, &userData); !ok {
			c.Abort()
			return
		}

		email := userData.Email

		var user models.DashboardUser
		if err := db.DB.Where("email = ?", email).First(&user).Error; err == nil {
			c.JSON(400, gin.H{"duplicate_erros": []string{"email"}})
			c.Abort()
			return
		}

		c.Next()
	}
}
