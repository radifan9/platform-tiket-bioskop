package utils

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/radifan9/platform-tiket-bioskop/models"
)

func HandleResponse(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, models.SuccessResponse{
		Success: true,
		Status:  status,
		Data:    data,
	})
}

func HandleError(ctx *gin.Context, status int, message string, err string) {
	log.Printf("%s\nCause: %s\n", message, err)
	ctx.JSON(status, models.ErrorResponse{
		Success: false,
		Status:  status,
		Error:   err,
	})
}
