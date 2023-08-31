package handler

import (
	avito "AvitoTask"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary		createUser
// @Tags			user
// @Description	Добавление и удаление пользователя в указанные сегменты
// @ID				create-user
// @Accept			json
// @Produce		json
// @Param			input	body		avito.UserList	true	"Поля added и expiry являются необязательными и имеют формат ДД.ММ.ГГГГ. segment_names_add - список сегментов, в которые нужно добавить пользователя. segment_names_remove - список сегментов, из которых нужно удалить пользователя. Названия сегментов в списках перечисляются через запятую"
// @Success		200		{object}	statusResponse
// @Failure		400,404	{object}	customError
// @Failure		500		{object}	customError
// @Failure		default	{object}	customError
// @Router			/api/users/ [post]
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

// @Summary		deleteUser
// @Tags			user
// @Description	Удаление пользователя из сегмента
// @ID				delete-user
// @Accept			json
// @Produce		json
// @Param			user		body		avito.UserSegmentWithId	true	"id пользователя и название сегмента, из которого его нужно удалить"
// @Success		200		{object}	statusResponse
// @Failure		400,404	{object}	customError
// @Failure		500		{object}	customError
// @Failure		default	{object}	customError
// @Router			/api/users/ [delete]
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

// @Summary		getUserInfo
// @Tags			user
// @Description	Получение информации о добавлении и удалении пользователя в сегменты во временном промежутке
// @ID				get-user-info
// @Accept			json
// @Produce		json
// @Param			user		body		avito.UserInfo	true	"id пользователя и временной промежуток. Поля end и start являются строками и имеют формат записи ДД.ММ.ГГГГ"
// @Success		200		{object}	getAllUserOperationsResponse
// @Failure		400,404	{object}	customError
// @Failure		500		{object}	customError
// @Failure		default	{object}	customError
// @Router			/api/userInfo/ [post]
func (h *Handler) getUserInfo(c *gin.Context) {
	var usr avito.UserInfo
	if err := c.BindJSON(&usr); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	operations, err := h.services.GetOperations(usr)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUserOperationsResponse{
		Data: operations,
	})
}
