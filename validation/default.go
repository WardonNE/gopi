package validation

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/wardonne/gopi/validation/translator"
)

var v = New().WithTranslator(
	new(translator.ENTranslator),
	new(translator.ENTranslator),
	new(translator.ZHTranslator),
).RegisterTagNameFunc(func(field reflect.StructField) string {
	label := field.Tag.Get("label")
	if label == "" {
		return field.Name
	}
	return label
})

func Default() *Validator {
	return v
}

func SetDefault(instance *Validator) {
	v = instance
}

func WithCustomRules(rules ...ICustomValidateRule) *Validator {
	return v.WithCustomRules(rules...)
}

func Translator(locale string) ut.Translator {
	return v.Translator(locale)
}

func Struct(s any) error {
	return v.Struct(s)
}

func StructFiltered(s any, fn validator.FilterFunc) error {
	return v.StructFiltered(s, fn)
}

func StructPartial(s any, fields ...string) error {
	return v.StructPartial(s, fields...)
}

func StructExcept(s any, fields ...string) error {
	return v.StructExcept(s, fields...)
}

func Var(field any, tag string) error {
	return v.Var(field, tag)
}

func VarWithValue(field, other any, tag string) error {
	return v.VarWithValue(field, other, tag)
}

func ValidateMap(data map[string]any, rules map[string]any) map[string]any {
	return v.ValidateMap(data, rules)
}

func RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return v.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

func RegisterAlias(alias, tags string) {
	v.RegisterAlias(alias, tags)
}

func RegisterStructValidation(fn validator.StructLevelFunc, types ...any) {
	v.RegisterStructValidation(fn, types...)
}

func RegisterStructValidationMapRules(rules map[string]string, types ...any) {
	v.RegisterStructValidationMapRules(rules, types...)
}

func RegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...any) {
	v.RegisterCustomTypeFunc(fn, types...)
}

func RegisterTranslation(tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc, translationFn validator.TranslationFunc) (err error) {
	return v.RegisterTranslation(tag, trans, registerFn, translationFn)
}
