package test

import (
	"context"
	"os"

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

	var db *gorm.DB
	{
		var err error
		// should mock instead
		dsn := "user= password= dbname= port= sslmode=disable"

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(1)
		}
	}
	repository := usersvc.NewRepo(db, logger)
	service = usersvc.NewBasicService(repository, logger)
	return service, context.Background()
}
