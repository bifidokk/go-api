package config

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Config struct {
	db           *gorm.DB
	Env          *Env
	Repositories *Repositories
	Services     *Services
}

func NewConfig() *Config {
	config := &Config{}

	return config
}

func Init(c *Config) error {
	start := time.Now()

	c.loadEnvironmentVariables()

	if err := c.connectDb(); err != nil {
		return err
	}

	log.Printf("config: successfully initialized [%s]", time.Since(start))

	return nil
}

func InitTest(c *Config) error {
	c.loadEnvironmentVariables()

	if err := c.connectDb(); err != nil {
		return err
	}

	c.InitTestDb()

	return nil
}
