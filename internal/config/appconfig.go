package config

import "sync"

type AppConfig struct {
	Port         string
	Version      string
	DSN          string
	IsProduction bool
	WG           *sync.WaitGroup
	MU           *sync.Mutex
}
