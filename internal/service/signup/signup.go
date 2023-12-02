package signup

import (
	"github.com/bifidokk/go-api/internal/entity"
	"github.com/bifidokk/go-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type signup struct {
	userRepository repository.UserRepository
}

type Signup interface {
	CreateUser(request SignupRequest) (*entity.User, error)
}

func NewSignup(userRepository repository.UserRepository) Signup {
	return &signup{
		userRepository: userRepository,
	}
}

func (s *signup) CreateUser(request SignupRequest) (*entity.User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    request.Email,
		Password: string(encryptedPassword),
	}

	createdUser, err := s.userRepository.Create(user)

	return createdUser, err
}
