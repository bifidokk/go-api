package entity

import (
	"time"

	"gorm.io/gorm"
)

type Notes []Note

type Note struct {
	gorm.Model
	ID              uint   `gorm:"primary_key"`
	NoteTitle       string `gorm:"type:VARCHAR(255);"`
	NoteDescription string `gorm:"type:TEXT;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
