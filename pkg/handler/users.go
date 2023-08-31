package handler

import (
	avito "AvitoTask"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary createUser
// @Tags users
// @Description Добавление и удаление пользователя в указанные сегменты
// @ID create-user
// @Accept json
// @Produce json
// @Param input body avito.User true "user data"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} customError
// @Failure 500 {object} customError
// @Failure default {object} customError
// @Router /api/users [post]
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

// @Summary deleteUser
// @Tags users
// @Description Удаление пользователя из сегмента
// @ID delete-user
// @Accept json
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} customError
// @Failure 500 {object} customError
// @Failure default {object} customError
// @Router /api/users [delete]
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

// @Summary getUserInfo
// @Tags users
// @Description Получение информации о добавлении и удалении пользователя в сегменты во временном промежутке
// @ID get-user-info
// @Accept json
// @Produce json
// @Param user body avito.UserInfo true "user data"
// @Success 200 {object} getAllUserOperationsResponse
// @Failure 400,404 {object} customError
// @Failure 500 {object} customError
// @Failure default {object} customError
// @Router /api/userInfo [post]
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
