package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

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

// CreateNewRestaurant godoc
// @summary Create new restaurant
// @tags Restaurant
// @id CreateNewRestaurant
// @Success 200 {object} model.MessageResponse
// @Failure 400
// @Failure 500
// @Router /restaurants [post]
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

// GetPopularRestaurant godoc
// @summary Get popular restaurants
// @tags Restaurant
// @id GetPopularRestaurant
// @Success 200 {array} model.Restaurant
// @Failure 404
// @Failure 500
// @Router /restaurants/popular [get]
func (h *restaurantHandler) GetPopularRestaurant(c *gin.Context) {
	out, err := h.restaurantService.GetPopularRestaurant()

	if err == nil {
		c.JSON(http.StatusOK, out)
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		err = errors.New("not matching data")
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	c.JSON(http.StatusInternalServerError, errorResponse(err))
}

// CreateOrUpdateRestaurantPopularity godoc
// @summary Create new restaurant popularity, or update popularity if restaurant exists
// @tags Restaurant
// @id CreateOrUpdateRestaurantPopularity
// @Success 200 {object} model.MessageResponse
// @Failure 400
// @Failure 500
// @Router /restaurants/popularity/{restaurantId} [post]
func (h *restaurantHandler) CreateOrUpdateRestaurantPopularityHandler(c *gin.Context) {
	restaurantId, err := strconv.ParseInt(c.Param("restaurantId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidQueryParam))
		return
	}

	err = h.restaurantService.CreateOrUpdateRestaurantPopularity(restaurantId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := "Create restaurant popularity successfully"
	c.JSON(http.StatusOK, messageResponse(response))
}
