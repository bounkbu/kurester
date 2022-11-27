package main

import (
	"github.com/BounkBU/kurester/config"
	"github.com/BounkBU/kurester/httpserver"
	"github.com/BounkBU/kurester/pkg/database"
	"github.com/BounkBU/kurester/pkg/logger"
	log "github.com/sirupsen/logrus"
)

var serverConfig *config.Config

func init() {
	serverConfig = config.LoadConfig()

	logger.InitLogger(serverConfig.App)
}

func main() {
	db, err := database.NewMySQLDatabaseConnection(serverConfig)
	if err != nil {
		log.Fatalf("error, create mysql database connection, %s", err.Error())
	}

	server := httpserver.NewHTTPServer(serverConfig, db)

	server.Start()
}
