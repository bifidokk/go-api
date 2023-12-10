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
	UpdateNote(request UpdateRequest, note *entity.Note) (*entity.Note, error)
	DeleteNote(note *entity.Note) error
}

func NewNoteService(noteRepository repository.NoteRepository) NoteService {
	return &noteService{
		noteRepository: noteRepository,
	}
}

func (noteService *noteService) CreateNote(request CreateRequest, user *entity.User) (*entity.Note, error) {
	note := &entity.Note{
		NoteTitle:       request.Title,
		NoteDescription: request.Description,
		UserID:          user.ID,
	}

	createdNote, err := noteService.noteRepository.Save(note)
	return createdNote, err
}

func (noteService *noteService) UpdateNote(request UpdateRequest, note *entity.Note) (*entity.Note, error) {
	note.NoteTitle = request.Title
	note.NoteDescription = request.Description

	createdNote, err := noteService.noteRepository.Save(note)
	return createdNote, err
}

func (noteService *noteService) DeleteNote(note *entity.Note) error {
	err := noteService.noteRepository.Delete(note)
	return err
}
