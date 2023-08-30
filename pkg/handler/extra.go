package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUserId(c *gin.Context) (int, error) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "could not get id from request")
		return 0, errors.New("could not get information from request")
	}
	userId := data["userId"].(float64) // Приведение к нужному типу
	return int(userId), nil
}

func getSegmentName(c *gin.Context) (string, error) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "could not get segment name from request")
		return "", errors.New("could not get information from request")
	}
	segmentName := data["segmentName"].(string) // Приведение к нужному типу
	return segmentName, nil
}

func getSegmentNameWithUserId(c *gin.Context) (int, string, error) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "could not get information from request")
		return 0, "", errors.New("could not get information from request")
	}
	userId := data["userId"].(float64)          // Приведение к нужному типу
	segmentName := data["segmentName"].(string) // Приведение к нужному типу
	return int(userId), segmentName, nil
}
