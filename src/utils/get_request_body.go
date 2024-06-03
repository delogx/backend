package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRequestBody[T interface{}](ctx *gin.Context, v *T) bool {
	requestBody, ok := ctx.Get("RequestBody")
	if ok {
		*v = requestBody.(T)
	} else if err := ctx.ShouldBindJSON(&v); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	ctx.Set("RequestBody", *v)
	return true
}
