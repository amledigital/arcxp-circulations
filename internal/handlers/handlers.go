package handlers

import (
	"github.com/amledigital/arcxp-circulations/internal/config"
	"github.com/amledigital/arcxp-circulations/internal/repository"
)

type HandlerRepo struct {
	App *config.AppConfig
	DB  repository.DBRepository
}

var Repo *HandlerRepo

func NewHandelerRepo(a *config.AppConfig, conn repository.DBRepository) *HandlerRepo {
	return &HandlerRepo{
		App: a,
		DB:  conn,
	}
}

func HandlerRepoInit(hr *HandlerRepo) {
	Repo = hr
}

func (hr *HandlerRepo) ConstructEnvelope(key string, payload any) map[string]any {
	envelope := make(map[string]any)

	envelope[key] = payload

	return envelope
}
