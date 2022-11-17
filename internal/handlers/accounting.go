package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) revenueСonfirmation(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input dataStruct
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := h.services.Accounting.RevenueConfirmation(userId, input.OrderId, input.ServiceId, input.Cost); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"Status": "ok",
	})
}
