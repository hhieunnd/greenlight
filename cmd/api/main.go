package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
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

	flag.StringVar(&cfg.db.dsn, "db-dns", "postgres://admin:It123456@@localhost:5432/greenlight?sslmode=disable", "PostgreSQL DNS")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	validator := validator.New()

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Printf("database connection pool established")

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
	err = svr.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
