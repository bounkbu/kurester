package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrInvalidRequestData = errors.New("invalid request data")

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func messageResponse(meesage string) gin.H {
	return gin.H{
		"message": meesage,
	}
}
