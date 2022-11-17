package handlers

import (
	"github.com/No1ball/avitoTask/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRout() *gin.Engine {
	router := gin.New()
	router.GET("/:id", h.getUserBalance)
	router.POST("/:id", h.addCashToUser)
	router.POST("/reserve/:id", h.reserveCash)
	router.PUT("/confirm/:id", h.revenueСonfirmation)
	return router
}
