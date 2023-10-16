package auth

import (
	"github.com/bifidokk/go-api/internal/entity"
	"github.com/bifidokk/go-api/internal/repository"
)

type auth struct {
	userRepository repository.UserRepository
}

type Auth interface {
	Authenticate(request LoginRequest) (string, error)
	GetUserByEmail(email string) (entity.User, error)
}

func NewAuth(userRepository repository.UserRepository) Auth {
	return &auth{
		userRepository: userRepository,
	}
}

func (auth *auth) Authenticate(request LoginRequest) (string, error) {
	return "", nil
}

func (auth *auth) GetUserByEmail(email string) (entity.User, error) {
	return auth.userRepository.FindByEmail(email)
}