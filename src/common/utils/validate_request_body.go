package utils

import (
	"backend/src/common/utils/validators"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateRequestBody[T interface{}](ctx *gin.Context, v *T) bool {
	if ctx.Request.ContentLength == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Request body is required"})
		return false
	}

	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})
	validate.RegisterValidation("hostname", validators.IsValidHostName)

	if ok := GetRequestBody(ctx, v); !ok {
		return false
	}

	if err := validate.Struct(*v); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return false
		}

		errors := err.(validator.ValidationErrors)
		errorMap := make(map[string]string, len(errors))

		for _, e := range errors {
			fieldName := e.Field()
			if len(e.Param()) > 0 {
				errorMap[fieldName] = fmt.Sprintf("%s: %s", e.Tag(), e.Param())
			} else {
				errorMap[fieldName] = e.Tag()
			}
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errorMap})
		return false
	}

	return true
}
