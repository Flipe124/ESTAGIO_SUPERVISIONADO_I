package api

import (
	"backend/internal/models"
	"backend/pkg/helpers/logger"

	"github.com/gin-gonic/gin"
)

// Return is the standard return of API and system log.
func Return(ctx *gin.Context, code int, message string) {
	ctx.JSON(
		code,
		&models.HTTP{
			Code:  code,
			Error: message,
		},
	)
	ctx.Abort()
}

// LogReturn is the standard return of API and system log.
func LogReturn(ctx *gin.Context, code int, message, log string) {
	ctx.JSON(
		code,
		&models.HTTP{
			Code:  code,
			Error: message,
		},
	)
	ctx.Abort()
	logger.Log.Error(log)
}
