package middlewares

import (
	"backend/db"
	"backend/src/auth/dtos"
	"backend/src/users/models"
	"backend/src/utils"

	"github.com/gin-gonic/gin"
)

type UserDTO struct {
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,min=11,max=11"`
}

func DuplicateUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userData dtos.RegisterDTO
		if ok := utils.ValidateRequestBody(c, &dtos.RegisterDTO{}); !ok {
			c.Abort()
			return
		}

		email := userData.Email

		var user models.User
		if err := db.DB.Where("email = ?", email).First(&user).Error; err == nil {
			c.JSON(400, gin.H{"duplicate_erros": []string{"email", "phone_number"}})
			c.Abort()
			return
		}

		c.Next()
	}
}
