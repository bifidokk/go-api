package entity

import (
	"time"
)

type Notes []Note

type Note struct {
	ID              uint   `gorm:"primary_key"`
	NoteTitle       string `gorm:"type:VARCHAR(255);"`
	NoteDescription string `gorm:"type:TEXT;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
