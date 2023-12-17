package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

type Env struct {
	DbDsn     string `mapstructure:"DB_DSN"`
	JwtSecret string `mapstructure:"JWT_SECRET"`
	JwtTtl    uint   `mapstructure:"JWT_TTL"`
	SentryDsn string `mapstructure:"SENTRY_DSN"`
}

func (c *Config) loadEnvironmentVariables() {
	viper.AutomaticEnv()

	jwtTtl, _ := strconv.Atoi(getEnv("JWT_TTL"))

	c.Env = &Env{
		DbDsn:     getEnv("DB_DSN"),
		JwtSecret: getEnv("JWT_SECRET"),
		JwtTtl:    uint(jwtTtl),
		SentryDsn: getEnv("SENTRY_DSN"),
	}
}

func getEnv(key string) string {
	value, ok := viper.Get(key).(string)

	if !ok {
		fmt.Printf("Invalid type assertion")
	}

	return value
}
