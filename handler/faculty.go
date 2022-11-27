package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/BounkBU/kurester/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type facultyHandler struct {
	facultyService service.FacultyService
}

func NewFacultyHandler(facultyService service.FacultyService) facultyHandler {
	return facultyHandler{
		facultyService: facultyService,
	}
}

func (h *facultyHandler) GetAllFaculty(c *gin.Context) {
	out, err := h.facultyService.GetAllFaculty()
	if err == nil {
		c.JSON(http.StatusOK, out)
		log.Info("return popular resturants")
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		err = errors.New("not matching data")
		c.JSON(http.StatusNotFound, errorResponse(err))
		log.Error(err)
		return
	}

	err = errors.New("something went wrong")
	c.JSON(http.StatusInternalServerError, errorResponse(err))
	log.Error(err)
}
