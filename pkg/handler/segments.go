package handler

import (
	avito "AvitoTask"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createSegment(c *gin.Context) {
	var input avito.Segment
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	answ, err := h.services.CreateSegment(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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

type getAllUserSegmentsResponse struct {
	Data []avito.Segment `json:"data"`
}

func (h *Handler) getUserSegments(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	segments, err := h.services.GetSegments(userId)

	c.JSON(http.StatusOK, getAllUserSegmentsResponse{
		Data: segments,
	})
}

func (h *Handler) deleteSegment(c *gin.Context) {
	segmentName, err := getSegmentName(c)
	if err != nil {
		return
	}
	err = h.services.DeleteSegment(segmentName)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Delete operation success!",
	})
}