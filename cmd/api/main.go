package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config    config
	logger    *log.Logger
	validator *validator.Validate
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	validator := validator.New()

	app := &application{
		config:    cfg,
		logger:    logger,
		validator: validator,
	}

	svr := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, svr.Addr)
	err := svr.ListenAndServe()
	logger.Fatal(err)
}
