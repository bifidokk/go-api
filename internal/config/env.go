package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DbDsn string `mapstructure:"DB_DSN"`
}

func (c *Config) loadEnvironmentVariables() {
	viper.SetConfigFile("./.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&c.Env); err != nil {
		log.Fatal(err)
	}
}
