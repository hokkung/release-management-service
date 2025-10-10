package config

type GitHubConfig struct {
	Token string `env:"TOKEN" envDefault:"localhost" required:"true"`
	Owner string `env:"OWNER" envDefault:"hok" required:"true"`
}
