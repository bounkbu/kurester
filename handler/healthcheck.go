package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealthcheck godoc
// @summary Healthcheck
// @tags healthcheck
// @id GetHealthcheck
// @Success 200
// @router / [get]
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"server": "running",
	})
}
