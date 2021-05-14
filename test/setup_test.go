package test

import (
	"context"
	"os"
	config2 "usersvc/config"
	"usersvc/usersvc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setup our service
func setup() (service usersvc.Service, ctx context.Context) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	config, err := config2.LoadConfig("../config")
	if err != nil {
		level.Error(logger).Log("exit", err)
	}

	var db *gorm.DB
	{
		var err error

		db, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(1)
		}
	}
	repository := usersvc.NewRepo(db, logger)
	service = usersvc.NewBasicService(repository, logger)
	return service, context.Background()
}
