package server

import (
	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/repository"
)

func RegisterRepositories(conf *config.Config) {
	var userRepository = repository.NewUserRepository(conf.Db())
	var noteRepository = repository.NewNoteRepository(conf.Db())

	conf.Repositories = &config.Repositories{
		UserRepository: userRepository,
		NoteRepository: noteRepository,
	}
}
