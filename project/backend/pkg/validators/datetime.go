package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// DateTime is the validator function for validate if the field contain datetime specification.
var DateTime validator.Func = func(field validator.FieldLevel) bool {
	if _, err := time.Parse("2006-01-02 15:04:05", field.Field().String()); err != nil {
		return false
	}
	return true
}
