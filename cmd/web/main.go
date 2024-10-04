package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/amledigital/arcxp-circulations/internal/config"
	"github.com/amledigital/arcxp-circulations/internal/database/sqldbrepo"
	"github.com/amledigital/arcxp-circulations/internal/handlers"
	"github.com/amledigital/arcxp-circulations/internal/repository/sqlrepo"
)

const (
	APP_IDLE_TIMEOUT        = time.Second * 90
	APP_READ_TIMEOUT        = time.Second * 90
	APP_WRITE_TIMEOUT       = time.Second * 90
	APP_READ_HEADER_TIMEOUT = time.Second * 10
)

var app config.AppConfig

func main() {

	parseFlags(&app)

	startServer()
}

func startServer() {

	app.CTX = context.Background()
	app.WG = &sync.WaitGroup{}
	app.MU = &sync.Mutex{}
	db, err := sqldbrepo.NewSQLConn(&app)
	if err != nil {
		log.Fatalln(err)
	}

	dbRepo := sqlrepo.NewSQLRepo(&app, db.Conn)

	handlerRepo := handlers.NewHandelerRepo(&app, dbRepo)

	handlers.HandlerRepoInit(handlerRepo)
	fmt.Printf("Starting server on %s", app.Port)

	srv := &http.Server{
		Addr:              app.Port,
		Handler:           routes(),
		IdleTimeout:       APP_IDLE_TIMEOUT,
		ReadTimeout:       APP_READ_TIMEOUT,
		WriteTimeout:      APP_WRITE_TIMEOUT,
		ReadHeaderTimeout: APP_READ_TIMEOUT,
		MaxHeaderBytes:    2 >> 20,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func parseFlags(app *config.AppConfig) {
	flag.StringVar(&app.Port, "port", ":8080", "The port the application runs on")
	flag.StringVar(&app.Version, "version", "0.0.1", "The app version")
	flag.StringVar(&app.ArcContentBase, "contentbase", "", "arcxp content base (sandbox|staging|production)")
	flag.StringVar(&app.ArcAccessToken, "accesstoken", "", "arcxp accesstoken (sandbox|staging|production)")
	flag.StringVar(&app.ArcWebsite, "arcwebsite", "910news", "arcxp website id (910news|mynorth)")
	flag.StringVar(&app.DSN, "DSN", "", "the database connection string for the sql driver")

	flag.Parse()

	fmt.Printf("%+v", app)

}
