package database

import (
	"os"

	"github.com/sirupsen/logrus"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const ConnectionDefault = "default"

func New(connection string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger:                   logger.Default.LogMode(getLoggerLevel()),
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"module": "gorm"}).Fatal(err)
	}

	return db
}

func getLoggerLevel() logger.LogLevel {
	loggerLevel := os.Getenv("DB_LOGGER_MODE")

	switch loggerLevel {
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	case "silent":
		return logger.Silent
	default:
		return logger.Silent
	}
}
