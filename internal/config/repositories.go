package config

import "github.com/bifidokk/go-api/internal/repository"

type Repositories struct {
	UserRepository repository.UserRepository
	NoteRepository repository.NoteRepository
}

func RegisterRepositories(conf *Config) {
	var userRepository = repository.NewUserRepository(conf.Db())
	var noteRepository = repository.NewNoteRepository(conf.Db())

	conf.Repositories = &Repositories{
		UserRepository: userRepository,
		NoteRepository: noteRepository,
	}
}
