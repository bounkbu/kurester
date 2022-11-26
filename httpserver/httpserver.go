package httpserver

import (
	"github.com/BounkBU/kurester/config"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	App      *gin.Engine
	Database *sqlx.DB
	Config   *config.Config
}

func NewHTTPServer(config *config.Config, db *sqlx.DB) *Server {
	gin.SetMode(config.App.GinMode)
	app := gin.Default()
	return &Server{
		App:      app,
		Database: db,
		Config:   config,
	}
}

func (server *Server) Start() {
	port := server.Config.App.Port

	log.Infof("Server is starting on port : %s", port)
	server.App.Run(":" + port)
}
