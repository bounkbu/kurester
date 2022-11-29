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

// SubmitForm godoc
// @summary Create form, analyze the appropriate menu from request and find nearest restaurants
// @tags Form
// @id SubmitForm
// @param data body model.Form true "request form data"
// @Success 200 {object} model.SubmitFormResponse
// @Failure 400
// @Failure 500
// @Router /form [post]
func (h *formHandler) SubmitFormHandler(c *gin.Context) {
	var req model.Form

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidRequestData))
		return
	}

	nearestRestaurants, err := h.restaurantService.GetNearestRestaurants(req.FacultyID)
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

	response := model.SubmitFormResponse{
		RecommendedMenu:   recommendedMenu,
		NearestRestaurant: nearestRestaurants,
	}
	c.JSON(http.StatusOK, response)
}
