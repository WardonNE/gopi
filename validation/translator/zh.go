package translator

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type ZHTranslator struct {
}

func (t *ZHTranslator) Build() locales.Translator {
	return zh.New()
}

func (t *ZHTranslator) RegisterTranslations(v *validator.Validate, translator ut.Translator) error {
	return zh_translations.RegisterDefaultTranslations(v, translator)
}
