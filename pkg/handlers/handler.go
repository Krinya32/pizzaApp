package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
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
