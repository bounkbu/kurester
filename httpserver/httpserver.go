package httpserver

import (
	"github.com/BounkBU/kurester/config"
	_ "github.com/BounkBU/kurester/docs"
	"github.com/BounkBU/kurester/handler"
	"github.com/BounkBU/kurester/repository"
	"github.com/BounkBU/kurester/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title KU Rester API
// @version 1.0
// @description The KU Rester web API

// @contact.name KU Rester Support
// @contact.email thanathip.suw@gmail.com

// @license.name MIT License
// @license.url https://choosealicense.com/licenses/mit/

// @schemes https
// @host kurester.herokuapp.com
func (server *Server) SetUpRouter() {
	server.App.Use(cors.Default())
	server.App.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	restaurantRepository := repository.NewRestaurantRepository(server.Database)
	menuRepository := repository.NewMenuRepository(server.Database)
	ratioRepository := repository.NewRatioRepository(server.Database)
	formRepository := repository.NewFormRepository(server.Database)
	facultyRepository := repository.NewFacultyRepository(server.Database)

	restaurantService := service.NewRestaurantService(restaurantRepository)
	menuService := service.NewMenuService(menuRepository, restaurantRepository)
	ratioService := service.NewRatioService(ratioRepository)
	formService := service.NewFormService(formRepository)
	facultyService := service.NewFacultyService(facultyRepository)

	restaurantHandler := handler.NewRestaurantHandler(restaurantService)
	menuHandler := handler.NewMenuHandler(menuService)
	formHandler := handler.NewFormHandler(menuService, restaurantService, formService)
	ratioHandler := handler.NewRatioHandler(ratioService)
	facultyHandler := handler.NewFacultyHandler(facultyService)

	server.App.GET("/", handler.HealthCheckHandler)
	server.App.GET("/faculties", facultyHandler.GetAllFaculty)
	server.App.GET("/restaurants/popular", restaurantHandler.GetPopularRestaurant)
	server.App.POST("/restaurants/popularity/:restaurantId", restaurantHandler.CreateOrUpdateRestaurantPopularityHandler)
	server.App.POST("/restarants", restaurantHandler.CreateNewRestaurantHandler)
	server.App.POST("/menus", menuHandler.CreateNewMenuHandler)
	server.App.GET("/menus/type", menuHandler.GetAllFoodType)
	server.App.GET("/menus/type/min-price", menuHandler.GetMenuMinPrice)
	server.App.POST("/form", formHandler.SubmitFormHandler)
	server.App.GET("/ratio/spicyness", ratioHandler.GetSpicynessRatioHandler)
	server.App.GET("/ratio/price", ratioHandler.GetPriceRatioHandler)
	server.App.GET("/ratio/type", ratioHandler.GetFoodTypeRatioHandler)
	server.App.GET("/ratio/popularity", ratioHandler.GetPopularityFromAverageMenuPrice)
	server.App.GET("/ratio/popularity/average", ratioHandler.GetAveragePopularityFromPriceRange)
}

func (server *Server) Start() {
	server.SetUpRouter()

	port := server.Config.App.Port

	log.Infof("Server is starting on port: %s", port)
	server.App.Run(":" + port)
}
