package repository

import "github.com/sirupsen/logrus"

func generateLogger(functionName string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"Module":  "Repository",
		"Funtion": functionName,
	})
}
