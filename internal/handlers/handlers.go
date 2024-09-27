package handlers

import (
	"github.com/amledigital/arcxp-circulations/internal/config"
	"github.com/amledigital/arcxp-circulations/internal/database"
)

type HandlerRepo struct {
	App *config.AppConfig
	DB  database.DBRepo
}

var Repo *HandlerRepo

func NewHandelerRepo(a *config.AppConfig, conn database.DBRepo) *HandlerRepo {
	return &HandlerRepo{
		App: a,
		DB:  conn,
	}
}

func HandlerRepoInit(hr *HandlerRepo) {
	Repo = hr
}
