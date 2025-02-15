package validation

import (
	"go-builder-plate/src/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateLogin(payload models.LoginRequest) error {
	return validate.Struct(payload)
}
