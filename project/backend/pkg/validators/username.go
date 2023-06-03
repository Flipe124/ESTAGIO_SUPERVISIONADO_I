package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Username is the validator function for validate if the field contain alpha and blank spaces.
var Username validator.Func = func(field validator.FieldLevel) bool {
	return regexp.MustCompile(
		`^[[:alnum:]]+[-_]?[[:alnum:]]+$`,
	).MatchString(
		field.Field().String(),
	)
}
