package config

import (
	"github.com/bifidokk/go-api/internal/service/auth"
	"github.com/bifidokk/go-api/internal/service/note"
)

type Services struct {
	AuthService auth.Auth
	NoteService note.NoteService
}

func RegisterServices(conf *Config) {
	var authService = auth.NewAuth(
		conf.Repositories.UserRepository,
		conf.Env.JwtSecret,
		int(conf.Env.JwtTtl),
	)

	var noteService = note.NewNoteService(
		conf.Repositories.NoteRepository,
	)

	conf.Services = &Services{
		AuthService: authService,
		NoteService: noteService,
	}
}
