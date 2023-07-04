package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Phrase is the validator function for validate if the field contain alpha and blank spaces.
var Phrase validator.Func = func(field validator.FieldLevel) bool {
	return regexp.MustCompile(
		`^([a-zA-ZÀ-ÿ]+[[:blank:]]?)+[a-zA-ZÀ-ÿ]{1}$`,
	).MatchString(
		field.Field().String(),
	)
}
