package handler

import (
	"net/http"

	"github.com/BounkBU/kurester/service"
	"github.com/gin-gonic/gin"
)

type ratioHandler struct {
	ratioService service.RatioService
}

func NewRatioHandler(ratioService service.RatioService) ratioHandler {
	return ratioHandler{
		ratioService: ratioService,
	}
}

func (h *ratioHandler) GetSpicynessRatioHandler(c *gin.Context) {
	spicynessRatio, err := h.ratioService.GetSpicynessRatio()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, spicynessRatio)
}

func (h *ratioHandler) GetPriceRatioHandler(c *gin.Context) {
	priceRatio, err := h.ratioService.GetPriceRatio()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, priceRatio.Results)
}

func (h *ratioHandler) GetFoodTypeRatioHandler(c *gin.Context) {
	foodTypeRatio, err := h.ratioService.GetFoodTypeRatio()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, foodTypeRatio)
}

func (h *ratioHandler) GetPopularityFromAverageMenuPrice(c *gin.Context) {
	popularity, err := h.ratioService.GetPopularityFromAverageMenuPrice()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, popularity)
}

func (h *ratioHandler) GetAveragePopularityFromPriceRange(c *gin.Context) {
	averagePopularity, err := h.ratioService.GetAveragePopularityFromPriceRange()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, averagePopularity.Results)
}
