package httpserver

import (
	"github.com/BounkBU/kurester/config"
	"github.com/BounkBU/kurester/handler"
	"github.com/BounkBU/kurester/repository"
	"github.com/BounkBU/kurester/service"
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

func (server *Server) SetUpRouter() {
	server.App.GET("/", handler.HealthCheckHandler)

	restaurantRepository := repository.NewRestaurantRepository(server.Database)
	restaurantService := service.NewRestaurantService(restaurantRepository)
	restaurantHandler := handler.NewRestaurantHandler(restaurantService)
	server.App.POST("/restarants", restaurantHandler.CreateNewRestaurantHandler)
}

func (server *Server) Start() {
	server.SetUpRouter()

	port := server.Config.App.Port

	log.Infof("Server is starting on port: %s", port)
	server.App.Run(":" + port)
}
