package validation

import (
	"be/src/domain/models"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) ValidateStruct(obj interface{}) []ValidationError {
	var errors []ValidationError
	err := v.validate.Struct(obj)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.Field = err.Field()
			element.Message = v.getMessage(err)
			errors = append(errors, element)
		}
	}

	return errors
}

func (v *Validator) getMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Minimum length is %s", err.Param())
	case "max":
		return fmt.Sprintf("Maximum length is %s", err.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", err.Param())
	default:
		return "Invalid value"
	}
}

func (v *Validator) ValidateLogin(req models.LoginRequest) []ValidationError {
	return v.ValidateStruct(req)
}

func (v *Validator) ValidateRegister(req models.RegisterRequest) []ValidationError {
	return v.ValidateStruct(req)
}
