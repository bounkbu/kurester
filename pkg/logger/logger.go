package logger

import (
	"github.com/BounkBU/kurester/config"
	"github.com/sirupsen/logrus"
)

func InitLogger(appConfig config.App) {
	if appConfig.Env == "production" {
		timeFormatLayout := "2006-01-02T15:04:05.000Z"
		logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyLevel: "log_level",
			},
			TimestampFormat: timeFormatLayout,
		})
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
