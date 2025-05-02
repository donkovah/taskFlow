package authController

import (
	validation "be/src/app/controllers/Validation"
	"be/src/domain/models"
)

type AuthValidator struct {
	validator *validation.Validator
}

func NewAuthValidator() *AuthValidator {
	return &AuthValidator{
		validator: validation.NewValidator(),
	}
}

func (v *AuthValidator) ValidateRegister(req models.RegisterRequest) []validation.ValidationError {
	return v.validator.ValidateStruct(req)
}

func (v *AuthValidator) ValidateLogin(req models.LoginRequest) []validation.ValidationError {
	return v.validator.ValidateStruct(req)
}
