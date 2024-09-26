package config

import "sync"

type AppConfig struct {
	Port           string
	ArcAccessToken string
	ArcContentBase string
	Version        string
	DSN            string
	IsProduction   bool
	WG             *sync.WaitGroup
	MU             *sync.Mutex
}
