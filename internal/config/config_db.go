package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (c *Config) connectDb() error {
	db, err := gorm.Open(postgres.Open(c.Env.DbDsn), &gorm.Config{})

	if err != nil || db == nil {
		return err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	c.db = db

	return nil
}

func (c *Config) Db() *gorm.DB {
	if c.db == nil {
		log.Fatal("config: database not connected")
	}

	return c.db
}
