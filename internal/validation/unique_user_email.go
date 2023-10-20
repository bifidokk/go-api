package validation

import "github.com/go-playground/validator/v10"

var UserEmailUnique validator.Func = func(fl validator.FieldLevel) bool {
	return false
}