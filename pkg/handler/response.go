package handler

import (
	avito "AvitoTask"
	"github.com/gin-gonic/gin"
	"log"
)

type customError struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type getAllUserOperationsResponse struct {
	Data []avito.Operation `json:"data"`
}

type getAllUserSegmentsResponse struct {
	Data []avito.Segment `json:"data"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, customError{message})
	log.Fatalf(message)
}
