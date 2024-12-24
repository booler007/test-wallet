package controller

import (
	"errors"
	"log"
	"net/http"

	"wallet/internal/service"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if err := ctx.Errors.Last(); err != nil {
			switch {
			case err.Type == gin.ErrorTypeBind:
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": gin.H{err.Error(): err.Meta}})
			case errors.Is(err, service.ErrInsufficientBalance):
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			case errors.Is(err, service.ErrWalletNotFound):
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			default:
				log.Print()
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
		}

	}
}
