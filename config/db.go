package config

import "fmt"

type DBConfig struct {
	Host     string `env:"HOST" envDefault:"localhost" required:"true"`
	Port     int    `env:"PORT" envDefault:"5432" required:"true"`
	DBName   string `env:"DB_NAME" envDefault:"rms" required:"true"`
	User     string `env:"USER" envDefault:"admin" required:"true"`
	Password string `env:"PASSWORD" envDefault:"admin" required:"true"`
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.Host, c.User, c.Password, c.DBName, c.Port)
}
