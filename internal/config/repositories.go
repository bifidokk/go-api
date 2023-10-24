package config

import "github.com/bifidokk/go-api/internal/repository"

type Repositories struct {
	UserRepository repository.UserRepository
	NoteRepository repository.NoteRepository
}
