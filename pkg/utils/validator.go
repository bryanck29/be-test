package utils

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var v *validator.Validate

func InitValidator() {
	v = validator.New()
}

// ParseParameter will parse request to struct
func ParseParameter(ctx echo.Context, i interface{}) error {
	return ctx.Bind(i)
}

// ValidateParameter will validate request
func ValidateParameter(ctx echo.Context, i interface{}) (err error) {
	err = v.Struct(i)
	return
}

// RegisterStructValidator registers a custom struct level validation to the validator
func RegisterStructValidator(fn func(validator.StructLevel)) {
	v.RegisterStructValidation(fn)
}
