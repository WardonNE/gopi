package validation

import "strings"

type IValidateForm interface {
	Fails() bool
	Errors() map[string][]string
	AddError(key, message string)
	CustomValidations() []CustomValidation
}

type ValidateForm struct {
	messages map[string][]string
}

func (form *ValidateForm) Empty() bool {
	if len(form.messages) == 0 {
		return true
	}
	for _, msgs := range form.messages {
		if len(msgs) > 0 {
			return false
		}
	}
	return true
}

func (form *ValidateForm) Fails() bool {
	return !form.Empty()
}

func (form *ValidateForm) Errors() map[string][]string {
	return form.messages
}

func (form *ValidateForm) AddError(key, message string) {
	if form.messages == nil {
		form.messages = make(map[string][]string)
	}
	msgs := form.messages[key]
	if len(msgs) == 0 {
		form.messages[key] = append(form.messages[key], strings.TrimSpace(message))
	} else {
		for _, msg := range msgs {
			if msg == message {
				return
			}
		}
		form.messages[key] = append(form.messages[key], strings.TrimSpace(message))
	}
}

func (form *ValidateForm) CustomValidations() []CustomValidation {
	return []CustomValidation{}
}
