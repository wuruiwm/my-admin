package util

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Trans ut.Translator
)

func ValidatorError(err error) (ret string) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			return e.Translate(Trans)
		}
	}
	return err.Error()
}
