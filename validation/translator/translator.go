package translator

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ITranslator interface {
	Build() locales.Translator
	RegisterTranslations(v *validator.Validate, translator ut.Translator) error
}
