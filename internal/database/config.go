// Package database.
// DBConfig is responsible for Database config.
package database

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

const MaxConn = 10

// DBConfig is a struct for env vars.
type DBConfig struct {
	DBHost     string `required:"true"`
	DBUserName string `required:"true"`
	DBPassword string `required:"true"`
	DBPort     string `required:"true"`
	DBName     string `required:"true"`
	Timeout    int    `required:"true"`
}

// NewDBConfig returns database config struct.
func NewDBConfig() *DBConfig {
	var cfg DBConfig

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	return &cfg
}
