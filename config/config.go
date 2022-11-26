package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Database Database
}

type App struct {
	Env     string `mapstructure:"ENV"`
	GinMode string `mapstructure:"GIN_MODE" default:"release"`
	Port    string `mapstructure:"PORT"`
}

type Database struct {
	Hostname string `mapstructure:"MYSQL_HOSTNAME"`
	Port     string `mapstructure:"MYSQL_PORT"`
	Username string `mapstructure:"MYSQL_USERNAME"`
	Password string `mapstructure:"MYSQL_PASSWORD"`
	Database string `mapstructure:"MYSQL_DATABASE"`
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	var appConfig App
	var databaseConfig Database

	appConfig.Env = os.Getenv("ENV")
	appConfig.GinMode = os.Getenv("GIN_MODE")
	appConfig.Port = os.Getenv("PORT")

	databaseConfig.Hostname = os.Getenv("MYSQL_HOSTNAME")
	databaseConfig.Port = os.Getenv("MYSQL_PORT")
	databaseConfig.Username = os.Getenv("MYSQL_USERNAME")
	databaseConfig.Password = os.Getenv("MYSQL_PASSWORD")
	databaseConfig.Database = os.Getenv("MYSQL_DATABASE")

	return &Config{
		App:      appConfig,
		Database: databaseConfig,
	}
}
