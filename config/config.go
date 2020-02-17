package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

// Configuration is
type Configuration struct {
	Host      string `env:"HOST" envDefault:"localhost"`
	Port      string `env:"PORT" envDefault:"80"`
	MySQLHost string `env:"MySQL_HOST" envDefault:"127.0.0.1"`
	MySQLPort string `env:"MySQL_PORT" envDefault:"3306"`
	MySQLUser string `env:"MySQL_USER" envDefault:"root"`
	MySQLPW   string `env:"MySQL_PW"`
	MySQLDB   string `env:"MySQL_DB" envDefault:"stec"`
}

// GetConfig is
func GetConfig() Configuration {
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return cfg
}
