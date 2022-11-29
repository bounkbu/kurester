package handler

import (
	"net/http"

	model "github.com/BounkBU/kurester/models"
	"github.com/BounkBU/kurester/service"
	"github.com/gin-gonic/gin"
)

type menuHandler struct {
	menuService service.MenuService
}

func NewMenuHandler(menuService service.MenuService) menuHandler {
	return menuHandler{
		menuService: menuService,
	}
}

// CreateNewMenu godoc
// @summary Create new menu
// @tags Menu
// @id CreateNewMenu
// @Success 200 {object} model.MessageResponse
// @Failure 400
// @Failure 500
// @Router /menus [post]
func (h *menuHandler) CreateNewMenuHandler(c *gin.Context) {
	var req model.Menu

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrInvalidRequestData))
		return
	}

	err := h.menuService.CreateNewMenu(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	message := "create new menu successfully"
	c.JSON(http.StatusOK, messageResponse(message))
}

// GetAllFoodType godoc
// @summary Get all the types of food
// @tags Menu
// @id GetAllFoodType
// @Success 200 {array} model.Menu
// @Failure 400
// @Failure 500
// @Router /menus/type [get]
func (h *menuHandler) GetAllFoodType(c *gin.Context) {
	foodTypes, err := h.menuService.GetAllFoodType()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, foodTypes)
}

// GetMenuMinPrice godoc
// @summary Get min price of each food type
// @tags Menu
// @id GetMenuMinPrice
// @Success 200 {object} map[string]float64
// @Failure 400
// @Failure 500
// @Router /menus/type/min-price [get]
func (h *menuHandler) GetMenuMinPrice(c *gin.Context) {
	menuMinPrice, err := h.menuService.GetMenuMinPrice()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, menuMinPrice)
}
