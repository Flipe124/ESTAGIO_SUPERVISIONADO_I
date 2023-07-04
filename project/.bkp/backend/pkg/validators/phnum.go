package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Phnum is the validator function for validate if the field contain alnum and blank spaces.
var Phnum validator.Func = func(field validator.FieldLevel) bool {
	return regexp.MustCompile(
		`^([a-zA-ZÀ-ÿ0-9]+[[:blank:]]?)+[a-zA-ZÀ-ÿ0-9]{1}$`,
	).MatchString(
		field.Field().String(),
	)
}
