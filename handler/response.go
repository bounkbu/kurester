package handler

import (
	"errors"

	model "github.com/BounkBU/kurester/models"
	"github.com/gin-gonic/gin"
)

var ErrInvalidRequestData = errors.New("invalid request data")
var ErrInvalidQueryParam = errors.New("invalid query parameter")

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func messageResponse(meesage string) model.MessageResponse {
	return model.MessageResponse{
		Message: meesage,
	}
}
