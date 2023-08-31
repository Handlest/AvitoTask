package handler

import (
	avito "AvitoTask"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary createSegment
// @Tags segment
// @Description Создание нового сегмента
// @ID create-segment
// @Accept json
// @Produce json
// @Param input body avito.Segment true "segment data"
// @Success 200 {string} string "response"
// @Failure 400,404 {object} customError
// @Failure 500 {object} customError
// @Failure default {object} customError
// @Router /api/getUserInfo [post]
// @Router /api/segments [post]
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

// @Summary getUserSegments
// @Tags segment
// @Description Получение всех пользователя
// @ID get-user-segments
// @Accept json
// @Produce json
// @Success 200 {object} getAllUserSegmentsResponse
// @Failure 400,404 {object} customError
// @Failure 500 {object} customError
// @Failure default {object} customError
// @Router /api/getUserSegments [post]
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

// @Summary deleteSegment
// @Tags segment
// @Description Удаление сегмента
// @ID delete-segment
// @Accept json
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} customError
// @Failure 500 {object} customError
// @Failure default {object} customError
// @Router /api/segments [delete]
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
