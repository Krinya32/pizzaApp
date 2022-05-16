package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/krinya32/pizzaApp"
	"net/http"
	"strconv"
)

type GetAllPizzasResponse struct {
	Data []pizzaApp.PizzaStruct `json:"data"`
}

func (h *Handler) getAllPizzas(c *gin.Context) {
	var input pizzaApp.PizzaStruct
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	pizzas, err := h.services.Pizzas.GetAll()
	if err != nil {

	}

	c.JSON(http.StatusOK, GetAllPizzasResponse{
		Data: pizzas,
	})
}

func (h *Handler) getPizzaById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id parameter")
		return
	}

	pizza, err := h.services.Pizzas.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, pizza)
}

func (h *Handler) deletePizzaById(c *gin.Context) {

}

func (h *Handler) updatePizza(c *gin.Context) {

}
