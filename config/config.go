package config

import "github.com/caarlos0/env/v11"

type Configuration struct {
	DB DBConfig `envPrefix:"DB_" envSeparator:"_" required:"true"`
}

func New() *Configuration {
	var cfg Configuration
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
