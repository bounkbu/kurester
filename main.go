package main

import (
	"github.com/BounkBU/kurester/config"
	"github.com/BounkBU/kurester/httpserver"
	"github.com/BounkBU/kurester/pkg/database"
	"github.com/BounkBU/kurester/pkg/logger"
	log "github.com/sirupsen/logrus"
)

var appConfig *config.Config

func init() {
	appConfig = config.LoadConfig()

	logger.InitLogger(appConfig.App)
}

func main() {
	db, err := database.NewMySQLDatabaseConnection(appConfig)
	if err != nil {
		log.Fatalf("error, create mysql database connection, %s", err.Error())
	}

	server := httpserver.NewHTTPServer(appConfig, db)

	server.Start()
}
