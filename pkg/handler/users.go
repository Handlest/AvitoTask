package handler

import (
	avito "AvitoTask"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createUser(c *gin.Context) {
	var users avito.UserList

	if err := c.BindJSON(&users); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input avito.User
	input.UserId = users.UserId
	input.Added = users.Added
	input.Expiry = users.Expiry

	for _, segmentName := range users.SegmentNamesAdd {
		input.SegmentName = segmentName
		err := h.services.CreateUser(input)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	for _, segmentName := range users.SegmentNamesRemove {
		err := h.services.DeleteUser(users.UserId, segmentName)
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "All operations success!",
	})
}

func (h *Handler) deleteUser(c *gin.Context) {
	userId, segmentName, err := getSegmentNameWithUserId(c)
	if err != nil {
		return
	}

	err = h.services.DeleteUser(userId, segmentName)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Delete operation success!",
	})
}

type getAllUserOperationsResponse struct {
	Data []avito.Operation `json:"data"`
}

func (h *Handler) getUserInfo(c *gin.Context) {
	var user avito.UserInfo
	if err := c.BindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	operations, err := h.services.GetOperations(user)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUserOperationsResponse{
		Data: operations,
	})
}
