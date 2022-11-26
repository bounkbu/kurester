package handler

import (
	"database/sql"
	"errors"
	"net/http"

	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

func (h *restaurantHandler) GetPopularRestaurant(c *gin.Context) {
	out, err := h.restaurantService.GetPopularRestaurant()
	if err == nil {
		c.JSON(http.StatusOK, out)
		log.Info("return popular resturants")
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})
		log.Error(err)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})
	log.Error(err)
}
