package server

import (
	"backend/pkg/validators"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if validator, isValidator := binding.Validator.Engine().(*validator.Validate); isValidator {
		validator.RegisterValidation("phrase", validators.Phrase)
		validator.RegisterValidation("phnum", validators.Phnum)
		validator.RegisterValidation("username", validators.Username)
		validator.RegisterValidation("length", validators.Length)
		validator.RegisterValidation("datetime", validators.DateTime)
	}
}
