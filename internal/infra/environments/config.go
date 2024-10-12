package environments

import (
	godotEnv "github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiEnv           string `envconfig:"API_ENV"`
	ApiHost          string `envconfig:"API_HOST"`
	ApiPort          int    `envconfig:"API_PORT"`
	DatabaseHost     string `envconfig:"DATABASE_HOST"`
	DatabasePort     int    `envconfig:"DATABASE_PORT"`
	DatabaseUsername string `envconfig:"DATABASE_USERNAME"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD"`
	DatabaseDBName   string `envconfig:"DATABASE_DBNAME"`
	DatabaseDrive    string `envconfig:"DATABASE_DRIVE"`
	DatabaseSchema   string `envconfig:"DATABASE_SCHEMA"`
}

var Env Config

func StartConfig() error {
	if err := godotEnv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	return nil
}
