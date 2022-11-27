package handler

import (
	"net/http"
	"time"

	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/service"
	"github.com/gin-gonic/gin"
)

type formHandler struct {
	menuService       service.MenuService
	restaurantService service.RestaurantService
}

func NewFormHandler(menuService service.MenuService, restaurantService service.RestaurantService) formHandler {
	return formHandler{
		menuService:       menuService,
		restaurantService: restaurantService,
	}
}

func (h *formHandler) SubmitFormHandler(c *gin.Context) {
	var req model.Form

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidRequestData))
		return
	}

	response := model.Form{
		ID:        1,
		FacaltyID: req.FacaltyID,
		Type:      req.Type,
		Price:     req.Price,
		IsSpicy:   req.IsSpicy,
		CreatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, response)
}
