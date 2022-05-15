package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/krinya32/pizzaApp/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	pizzas := router.Group("/pizza")
	{
		pizzas.GET("/", h.getAllPizzas)
		pizzas.GET("/:id", h.getPizzaById)
		pizzas.DELETE("/:id", h.deletePizzaById)
		pizzas.PUT("/:id", h.updatePizza)
	}
	return router
}
