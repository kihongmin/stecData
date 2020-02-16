package crawler

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

type Job struct {
	URL    string
	Title  string
	Origin string
}

type URLs struct {
	ID     int
	Title  string
	Origin string
	// start_date string
	// end_date string
	// position string
	URL string
	// basic string
	// advanced string
}

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

// errHandler is errHandler
func ErrHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
