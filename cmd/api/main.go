package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/jtowe1/whats-on-tap-api/internal/data"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config    config
	BeerModel data.BeerModel
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "MySQL DSN")

	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Println("database connection pool open")

	app := &application{
		config:    cfg,
		BeerModel: data.BeerModel{DB: db},
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.db.dsn)
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
