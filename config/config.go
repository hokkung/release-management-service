package config

import "github.com/caarlos0/env/v11"

type Configuration struct {
	GitHub GitHubConfig `envPrefix:"GITHUB_" envSeparator:"_" required:"true"`  
	DB DBConfig `envPrefix:"DB_" envSeparator:"_" required:"true"`
}

func New() *Configuration {
	var cfg Configuration
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
