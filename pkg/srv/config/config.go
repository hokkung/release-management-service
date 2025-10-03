package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"go.uber.org/zap"
)

// Config ...
type Config struct {
	AppName string `env:"SRV_APP_NAME" required:"true"`
	Port    int    `env:"SRV_SERVER_PORT" required:"true"`
}

// NewConfig ...
func NewConfig() *Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Panic("unable to bind env", zap.Error(err))
	}

	return &cfg
}
