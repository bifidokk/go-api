package auth

import (
	"errors"

	"github.com/bifidokk/go-api/internal/entity"
	"github.com/bifidokk/go-api/internal/repository"
	"github.com/bifidokk/go-api/internal/service/token"
	"golang.org/x/crypto/bcrypt"
)

type auth struct {
	userRepository repository.UserRepository
	jwtSecret      string
	jwtTtl         int
}

type Auth interface {
	Authenticate(request LoginRequest) (string, error)
	GetUserByEmail(email string) (entity.User, error)
}

func NewAuth(
	userRepository repository.UserRepository,
	jwtSecret string,
	jwtTtl int,
) Auth {
	return &auth{
		userRepository: userRepository,
		jwtSecret:      jwtSecret,
		jwtTtl:         jwtTtl,
	}
}

func (auth *auth) Authenticate(request LoginRequest) (string, error) {
	user, err := auth.GetUserByEmail(request.Email)

	if err != nil {
		return "", errors.New("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return "", errors.New("invalid credentials")
	}

	accessToken, err := token.CreatAccessToken(
		&user,
		auth.jwtSecret,
		auth.jwtTtl,
	)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (auth *auth) GetUserByEmail(email string) (entity.User, error) {
	return auth.userRepository.FindByEmail(email)
}
