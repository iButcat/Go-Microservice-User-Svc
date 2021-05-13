package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// internal pkg
	config2 "usersvc/config"
	"usersvc/usersvc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := config2.LoadConfig("./config")
	if err != nil {
		fmt.Print(err)
	}

	var (
		httpAddr = flag.String("http.addr", config.Port, "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
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

	var service usersvc.Service
	{
		repository := usersvc.NewRepo(db, logger)
		service = usersvc.NewBasicService(repository, logger)
	}

	var h http.Handler
	{
		h = usersvc.MakeHTTPHandler(service, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
