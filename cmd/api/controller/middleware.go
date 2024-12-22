package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if err := ctx.Errors.Last(); err != nil {
			switch {
			case err.Type == gin.ErrorTypeBind:
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": gin.H{err.Error(): err.Meta}})
			default:
				log.Print()
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
		}

	}
}
