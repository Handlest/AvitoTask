package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type customError struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, customError{message})
	log.Fatalf(message)
}
