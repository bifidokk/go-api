package config

import (
	"github.com/bifidokk/go-api/internal/validation"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidators(conf *Config) {
	userEmailUniqueValidator := validation.NewUserEmailUniqueValidator(conf.Repositories.UserRepository)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userEmailUnique", userEmailUniqueValidator.GetValidator())
	}
}
