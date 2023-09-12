package translator

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type ENTranslator struct {
}

func (t *ENTranslator) Build() locales.Translator {
	return en.New()
}

func (t *ENTranslator) RegisterTranslations(v *validator.Validate, translator ut.Translator) error {
	return en_translations.RegisterDefaultTranslations(v, translator)
}
