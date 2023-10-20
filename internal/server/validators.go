package server

import (
	"github.com/bifidokk/go-api/internal/validation"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userEmailUnique", validation.UserEmailUnique)
	}
}