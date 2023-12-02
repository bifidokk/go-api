package entity

import (
	"time"
)

type Notes []Note

type Note struct {
	ID              uint      `gorm:"primary_key" json:"id"`
	NoteTitle       string    `gorm:"type:VARCHAR(255)" json:"title"`
	NoteDescription string    `gorm:"type:TEXT" json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	User            User      `json:"-"`
	UserID          uint      `gorm:"not null" json:"user_id"`
}

func (Note) TableName() string {
	return "notes"
}
