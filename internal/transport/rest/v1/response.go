package v1

import (
	"github.com/Smirnov-O/noter/pkg/logger"
	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
