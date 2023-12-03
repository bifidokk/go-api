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
	FindByUser(user *entity.User) (entity.Notes, error)
	FindUserNoteById(user *entity.User, noteId int) (*entity.Note, error)
	Save(note *entity.Note) (*entity.Note, error)
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

func (nr *noteRepository) FindByUser(user *entity.User) (results entity.Notes, err error) {
	err = nr.database.Find(&results, entity.Note{UserID: user.ID}).Error

	return results, err
}

func (nr *noteRepository) Save(note *entity.Note) (*entity.Note, error) {
	result := nr.database.Save(&note)

	return note, result.Error
}

func (nr *noteRepository) FindUserNoteById(user *entity.User, noteId int) (*entity.Note, error) {
	var note entity.Note

	result := nr.database.First(
		&note, entity.Note{
			ID:     uint(noteId),
			UserID: user.ID,
		},
	)

	return &note, result.Error
}
