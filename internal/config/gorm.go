package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}

func NewGORM(config *viper.Viper, log *logrus.Logger) *gorm.DB {
	var dsn string

	if isuri := config.GetBool("database.uri"); isuri {
		log.Info("Using URI")
		dsn = config.GetString("database.uri")
	} else {
		log.Info("Using individual config")
		username := config.GetString("database.username")
		password := config.GetString("database.password")
		host := config.GetString("database.host")
		port := config.GetInt("database.port")
		database := config.GetString("database.name")
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, password, host, port, database)
	}
	log.Tracef("DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})

	if err != nil {
		log.WithError(err).Fatal("failed to connect to database")
	}

	return db
}
