package handler

import (
	"net/http"

	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/service"
	"github.com/gin-gonic/gin"
)

type restaurantHandler struct {
	restaurantService service.RestaurantService
}

func NewRestaurantHandler(restaurantService service.RestaurantService) restaurantHandler {
	return restaurantHandler{
		restaurantService: restaurantService,
	}
}

func (h *restaurantHandler) CreateNewRestaurantHandler(c *gin.Context) {
	var req model.Restaurant

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidRequestData))
		return
	}

	err := h.restaurantService.CreateNewRestaurant(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	message := "create new restaurant successfully"
	c.JSON(http.StatusOK, messageResponse(message))
}
