package config

import (
	"log"
	"time"

	"github.com/bifidokk/go-api/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (c *Config) connectDb() error {
	dsn := "host=localhost user=postgres password=postgres dbname=api port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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

	entity.SetDbProvider(c)

	return nil
}

func (c *Config) Db() *gorm.DB {
	if c.db == nil {
		log.Fatal("config: database not connected")
	}

	return c.db
}
