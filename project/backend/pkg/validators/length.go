package validators

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

// Length is the validator function for validate integer length fields.
var Length validator.Func = func(field validator.FieldLevel) bool {
	size, err := strconv.Atoi(field.Param())
	if err != nil {
		panic("validator: wrong data type in tag value")
	}
	return len(strconv.Itoa(int(field.Field().Int()))) == size
}
