package utils

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/radifan9/platform-tiket-bioskop/models"
)

func HandleError(ctx *gin.Context, status int, err string, logMsg string) {
	log.Printf("%s\nCause: %s\n", logMsg, err)
	ctx.JSON(status, models.ErrorResponse{
		Success: false,
		Status:  status,
		Error:   err,
	})
}

func HandleMiddlewareError(ctx *gin.Context, status int, err string, logMsg string) {
	log.Printf("%s\nCause: %s\n", logMsg, err)
	ctx.AbortWithStatusJSON(status, models.ErrorResponse{
		Success: false,
		Status:  status,
		Error:   err,
	})
}
