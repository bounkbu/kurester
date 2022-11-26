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

func (h *menuHandler) GetSpicynessRatioHandler(c *gin.Context) {
	spicynessRatio, err := h.menuService.GetSpicynessRatio()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, spicynessRatio)
}
