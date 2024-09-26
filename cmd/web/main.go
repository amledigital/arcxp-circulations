package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/amledigital/arcxp-circulations/internal/config"
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

	flag.Parse()

}
