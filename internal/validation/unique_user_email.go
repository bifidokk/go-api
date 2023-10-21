package validation

import (
	"github.com/bifidokk/go-api/internal/repository"
	"github.com/go-playground/validator/v10"
)

type userEmailUniqueValidator struct {
	userRepository repository.UserRepository
}

type UserEmailUniqueValidator interface {
	GetValidator() validator.Func
}

func NewUserEmailUniqueValidator(userRepository repository.UserRepository) UserEmailUniqueValidator {
	return &userEmailUniqueValidator{
		userRepository: userRepository,
	}
}

func (v *userEmailUniqueValidator) GetValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		return false
	}
}
