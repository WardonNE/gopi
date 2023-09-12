package validation

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/wardonne/gopi/validation/translator"
)

var v = New().WithTranslator(
	new(translator.ENTranslator),
	new(translator.ENTranslator),
	new(translator.ZHTranslator),
)

func Default() *Validator {
	return v
}

type Validator struct {
	*validator.Validate
	uni *ut.UniversalTranslator
}

func New() *Validator {
	return &Validator{
		Validate: validator.New(),
	}
}

func (v *Validator) WithTranslator(fallback translator.ITranslator, translators ...translator.ITranslator) *Validator {
	v.uni = ut.New(fallback.Build())
	fb := v.uni.GetFallback()
	if err := fallback.RegisterTranslations(v.Validate, fb); err != nil {
		panic(err)
	}
	for _, translator := range translators {
		t := translator.Build()
		if err := v.uni.AddTranslator(t, true); err != nil {
			panic(err)
		} else {
			unit, _ := v.uni.GetTranslator(t.Locale())
			if err := translator.RegisterTranslations(v.Validate, unit); err != nil {
				panic(err)
			}
		}
	}
	return v
}

func (v *Validator) WithCustomRules(rules ...ICustomValidateRule) *Validator {
	for _, rule := range rules {
		if err := v.RegisterValidation(rule.Tag(), rule.Validate, !rule.SkipNull()); err != nil {
			panic(err)
		}
	}
	return v
}

func (v *Validator) Translator(locale string) ut.Translator {
	translator, _ := v.uni.GetTranslator(locale)
	return translator
}
