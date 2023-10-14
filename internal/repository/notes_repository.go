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

func (tr *noteRepository) FindAll() (results entity.Notes, err error) {
	err = tr.database.Find(&results).Error

	return results, err
}
