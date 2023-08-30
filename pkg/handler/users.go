package handler

import (
	avito "AvitoTask"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createUser(c *gin.Context) {
	var input avito.User
	fmt.Println("Я тута!!!")
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("И тута!!!")
	id, err := h.services.CreateUser(input) //!!!
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) createUserTTL(c *gin.Context) {

}

func (h *Handler) deleteUser(c *gin.Context) {

}
