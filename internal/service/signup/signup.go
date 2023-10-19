package signup

import (
	"github.com/bifidokk/go-api/internal/entity"
	"github.com/bifidokk/go-api/internal/repository"
	"github.com/bifidokk/go-api/internal/service/auth"
	"golang.org/x/crypto/bcrypt"
)

type signup struct {
	userRepository repository.UserRepository
	auth           auth.Auth
}

type Signup interface {
	CreateUser(request SignupRequest) (entity.User, error)
}

func NewSignup(userRepository repository.UserRepository, auth auth.Auth) Signup {
	return &signup{
		userRepository: userRepository,
		auth:           auth,
	}
}

func (s *signup) CreateUser(request SignupRequest) (entity.User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Email:    request.Email,
		Password: string(encryptedPassword),
	}

	user, err = s.userRepository.Create(user)

	return user, err
}
