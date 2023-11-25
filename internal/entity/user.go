package entity

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Email     string `gorm:"type:VARCHAR(255);uniqueIndex"`
	Password  string `gorm:"type:VARCHAR(255);"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Notes     []Note `gorm:"constraint:OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "users"
}
