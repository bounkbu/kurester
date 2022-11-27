package handler

import (
	"net/http"

	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/service"
	"github.com/gin-gonic/gin"
)

type formHandler struct {
	menuService       service.MenuService
	restaurantService service.RestaurantService
	formService       service.FormService
}

func NewFormHandler(menuService service.MenuService, restaurantService service.RestaurantService, formService service.FormService) formHandler {
	return formHandler{
		menuService:       menuService,
		restaurantService: restaurantService,
		formService:       formService,
	}
}

func (h *formHandler) SubmitFormHandler(c *gin.Context) {
	var req model.Form

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidRequestData))
		return
	}

	nearestRestaurants, err := h.restaurantService.GetNearestRestaurants(req.FacaltyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	recommendedMenu, err := h.menuService.GetRecommendedMenu(req.Type, req.IsSpicy, req.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = h.formService.CreateNewForm(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := gin.H{
		"recommended_menu":    recommendedMenu,
		"nearest_restaurants": nearestRestaurants,
	}
	c.JSON(http.StatusOK, response)
}
