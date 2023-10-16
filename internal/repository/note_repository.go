package repository

import (
	"github.com/bifidokk/go-api/internal/entity"
	"gorm.io/gorm"
)

type noteRepository struct {
	database *gorm.DB
}

type NoteRepository interface {
	FindAll() (entity.Notes, error)
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{
		database: db,
	}
}

func (nr *noteRepository) FindAll() (results entity.Notes, err error) {
	err = nr.database.Find(&results).Error

	return results, err
}
