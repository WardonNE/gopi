package validation

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// Engine validation engine interface
type Engine interface {
	Translator(locale string) ut.Translator
	Struct(any) error
}

// IValidateForm validate form interface
type IValidateForm interface {
	// SetEngine set validation engine
	SetEngine(engine Engine)
	// Engine get validation engine
	Engine() Engine
	// AutoValidate should auth run validate when a request entered
	AutoValidate() bool
	// Validate run validate
	Validate(form IValidateForm)
	// BeforeValidate before validate event
	BeforeValidate() bool
	// AfterValidate after validate event
	AfterValidate()
	// SetLocale set locale
	SetLocale(locale string)
	// Locale get locale
	Locale() string
	// Fails returns whether the validation failed
	Fails() bool
	// Errors returns error messages
	Errors() map[string][]string
	// AddError add error message
	AddError(key, message string)
	// CustomValidations custom validations
	CustomValidations() []CustomValidation
}

type Form struct {
	messages map[string][]string

	locale string
	engine Engine
	form   IValidateForm
}

func (f *Form) SetEngine(engine Engine) {
	f.engine = engine
}

func (f *Form) Engine() Engine {
	return f.engine
}

func (f *Form) AutoValidate() bool {
	return true
}

func (f *Form) SetLocale(locale string) {
	f.locale = locale
}

func (f *Form) Locale() string {
	return f.locale
}

func (f *Form) BeforeValidate() bool {
	return true
}

func (f *Form) AfterValidate() {}

func (f *Form) Validate(form IValidateForm) {
	err := f.engine.Struct(form)
	if err == nil {
		return
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		panic(err)
	}
	translator := form.Engine().Translator(form.Locale())
	for _, err := range errs {
		form.AddError(err.Field(), err.Translate(translator))
	}
	customValidations := form.CustomValidations()
	for _, customValidation := range customValidations {
		customValidation(form)
	}
}

func (f *Form) Empty() bool {
	if len(f.messages) == 0 {
		return true
	}
	for _, msgs := range f.messages {
		if len(msgs) > 0 {
			return false
		}
	}
	return true
}

func (f *Form) Fails() bool {
	return !f.Empty()
}

func (f *Form) Errors() map[string][]string {
	return f.messages
}

func (f *Form) AddError(key, message string) {
	if f.messages == nil {
		f.messages = make(map[string][]string)
	}
	msgs := f.messages[key]
	if len(msgs) == 0 {
		f.messages[key] = append(f.messages[key], strings.TrimSpace(message))
	} else {
		for _, msg := range msgs {
			if msg == message {
				return
			}
		}
		f.messages[key] = append(f.messages[key], strings.TrimSpace(message))
	}
}

func (f *Form) CustomValidations() []CustomValidation {
	return []CustomValidation{}
}
