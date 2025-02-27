package customvalidator

import (
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	v10nonstandard "github.com/go-playground/validator/v10/non-standard/validators"
)

// reference:
// https://github.com/kittipat1413/go-common/blob/main/framework/validator/custom_validator.go
// https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Non_standard_validators

type NotBlankValidator struct{}

// Tag returns the tag identifier used in struct field validation tags.
func (*NotBlankValidator) Tag() string {
	return "notblank"
}

// Func returns the validator.Func that performs the validation logic.
func (*NotBlankValidator) Func() v10.Func {
	return v10nonstandard.NotBlank
}

// Translation returns the translation text and an custom translation function for the custom validator.
func (*NotBlankValidator) Translation() (string, v10.TranslationFunc) {
	translationText := "{0} cannot be blank"

	customTransFunc := func(ut ut.Translator, fe v10.FieldError) string {
		// {0} will be replaced with fe.Field()
		t, _ := ut.T(fe.Tag(), fe.Field())
		return t
	}

	return translationText, customTransFunc
}
