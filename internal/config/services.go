package config

import "github.com/bifidokk/go-api/internal/service/auth"

type Services struct {
	AuthService auth.Auth
}

func RegisterServices(conf *Config) {
	var authService = auth.NewAuth(
		conf.Repositories.UserRepository,
		conf.Env.JwtSecret,
		int(conf.Env.JwtTtl),
	)

	conf.Services = &Services{
		AuthService: authService,
	}
}
