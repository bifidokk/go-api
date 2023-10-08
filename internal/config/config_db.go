package config

import (
	"time"

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

	if (err != nil) {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	c.db = db

	return nil
}