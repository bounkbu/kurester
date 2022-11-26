package httpserver

import (
	"github.com/BounkBU/kurester/config"
	"github.com/BounkBU/kurester/handler"
	"github.com/BounkBU/kurester/repository"
	"github.com/BounkBU/kurester/service"
	"github.com/gin-contrib/cors"
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
	server.App.Use(cors.Default())
	server.App.GET("/", handler.HealthCheckHandler)

	restaurantRepository := repository.NewRestaurantRepository(server.Database)
	menuRepository := repository.NewMenuRepository(server.Database)

	restaurantService := service.NewRestaurantService(restaurantRepository)
	menuService := service.NewMenuService(menuRepository)
	// userService := service.NewUserService(restaurantRepository)

	restaurantHandler := handler.NewRestaurantHandler(restaurantService)
	menuHandler := handler.NewMenuHandler(menuService)

	server.App.POST("/restarants", restaurantHandler.CreateNewRestaurantHandler)
	server.App.POST("/menus", menuHandler.CreateNewMenuHandler)
}

func (server *Server) Start() {
	server.SetUpRouter()

	port := server.Config.App.Port

	log.Infof("Server is starting on port: %s", port)
	server.App.Run(":" + port)
}
