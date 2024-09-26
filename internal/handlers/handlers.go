package handlers

import "github.com/amledigital/arcxp-circulations/internal/config"

type HandlerRepo struct {
	app *config.AppConfig
	db  database.dbrepo
}

var app *HandlerRepo

func NewHandelerRepo(a *config.AppConfig) {
	app = a
}
