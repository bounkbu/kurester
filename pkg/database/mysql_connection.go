package database

import (
	"github.com/BounkBU/kurester/config"
	"github.com/BounkBU/kurester/pkg/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func NewMySQLDatabaseConnection(config *config.Config) (*sqlx.DB, error) {
	mysqlUrl := util.NewConnectionUrlBuilder("dns", config.Database)
	db, err := sqlx.Connect("mysql", mysqlUrl)
	if err != nil {
		log.Errorf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Errorf("error, can't ping the database, %s", err.Error())
		return nil, err
	}

	log.Info("Connected to mysql database successfully")
	return db, nil
}
