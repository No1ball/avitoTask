package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type dataStruct struct {
	OrderId   int `json:"orderId"`
	ServiceId int `json:"serviceId"`
	Cost      int `json:"cost"`
}

func (h *Handler) getUserBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	balance, err := h.services.User.GetUserBalance(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": id,
		"balance": balance,
	})
}

func (h *Handler) reserveCash(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input dataStruct
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.User.ReserveCash(userId, input.OrderId, input.ServiceId, input.Cost, "", ""); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"Status": "ok",
	})

}

func (h *Handler) addCashToUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input dataStruct
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("%d - cost", input.Cost)
	if err := h.services.User.AddCashToUser(userId, input.Cost); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"Status": "ok",
	})
}
