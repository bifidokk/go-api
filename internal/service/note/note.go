package note

import (
	"github.com/bifidokk/go-api/internal/entity"
	"github.com/bifidokk/go-api/internal/repository"
)

type noteService struct {
	noteRepository repository.NoteRepository
}

type NoteService interface {
	CreateNote(request CreateRequest, user *entity.User) (*entity.Note, error)
}

func NewNoteService(noteRepository repository.NoteRepository) NoteService {
	return &noteService{
		noteRepository: noteRepository,
	}
}

func (noteService *noteService) CreateNote(request CreateRequest, user *entity.User) (*entity.Note, error) {
	return nil, nil
}
