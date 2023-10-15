package entity

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Email     string `gorm:"type:VARCHAR(255);"`
	Password  string `gorm:"type:VARCHAR(255);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
