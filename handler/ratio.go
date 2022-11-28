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

// GetSpicynessRatio godoc
// @summary Get the spicyness ratio
// @tags Visualization
// @id GetSpicynessRatio
// @Success 200 {array} model.SpicynessRatio
// @Failure 500
// @Router /ratio/spicyness [get]
func (h *ratioHandler) GetSpicynessRatioHandler(c *gin.Context) {
	spicynessRatio, err := h.ratioService.GetSpicynessRatio()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, spicynessRatio)
}

// GetPriceRatio godoc
// @summary Get the price ratio
// @tags Visualization
// @id GetPriceRatio
// @Success 200 {array} model.PriceRatio
// @Failure 500
// @Router /ratio/price [get]
func (h *ratioHandler) GetPriceRatioHandler(c *gin.Context) {
	priceRatio, err := h.ratioService.GetPriceRatio()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, priceRatio)
}

// GetFoodTypeRatio godoc
// @summary Get the type of food ratio
// @tags Visualization
// @id GetFoodTypeRatio
// @Success 200 {array} model.FoodTypeRatio
// @Failure 500
// @Router /ratio/type [get]
func (h *ratioHandler) GetFoodTypeRatioHandler(c *gin.Context) {
	foodTypeRatio, err := h.ratioService.GetFoodTypeRatio()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, foodTypeRatio)
}

// GetPopularityFromAverageMenuPrice godoc
// @summary Get the popularity compare with average menu price
// @tags Visualization
// @id GetPopularityFromAverageMenuPrice
// @Success 200 {array} model.PopularityFromAverageMenuPrice
// @Failure 500
// @Router /ratio/popularity [get]
func (h *ratioHandler) GetPopularityFromAverageMenuPrice(c *gin.Context) {
	popularity, err := h.ratioService.GetPopularityFromAverageMenuPrice()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, popularity)
}

// GetAveragePopularityFromPriceRange godoc
// @summary Get the average popularity compare with range of menu price
// @tags Visualization
// @id GetAveragePopularityFromPriceRange
// @Success 200 {array} model.PriceRatio
// @Failure 500
// @Router /ratio/popularity/average [get]
func (h *ratioHandler) GetAveragePopularityFromPriceRange(c *gin.Context) {
	averagePopularity, err := h.ratioService.GetAveragePopularityFromPriceRange()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, averagePopularity)
}
