package config

import (
	"context"
	"sync"
)

type AppConfig struct {
	Port           string
	ArcAccessToken string
	ArcContentBase string
	ArcWebsite     string
	Version        string
	DSN            string
	IsProduction   bool
	CTX            context.Context
	WG             *sync.WaitGroup
	MU             *sync.Mutex
}
