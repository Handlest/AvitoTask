package handler

import (
	avito "AvitoTask"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createSegment(c *gin.Context) {
	fmt.Println("Я тут 0")
	var input avito.Segment
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("Я тут 1")
	answ, err := h.services.CreateSegment(input)
	fmt.Println("Я тут 2")
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("Я тут 2")
	var response string
	if answ == true {
		response = "Added segment successfully!"
	} else {
		response = "ERROR occurred!"
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"response": response,
	})
}

func (h *Handler) deleteSegment(c *gin.Context) {

}

func (h *Handler) getAllSegments(c *gin.Context) {

}
