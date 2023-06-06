package validator

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *validator.Validate {
	vald := validator.New()
	vald.RegisterValidation("excluded_if", excludedIf)
	vald.RegisterValidation("alpha_with_space_and_dot", AlphaWithSpaceAndDot)
	vald.RegisterValidation("password_atleast_one_uppercase_lowercase_symbol_and_number", PasswordAtleastOneUppercaseLowercaseSymbolAndNumber)
	vald.RegisterValidation("not_xss_script", NotXssScript)
	return vald
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
