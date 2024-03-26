package exception

import "fmt"

// InvalidParamTypeErr invalid arg type error
type InvalidParamTypeErr struct {
	token string
	arg   any
}

func (e *InvalidParamTypeErr) Error() string {
	return fmt.Sprintf("invalid args type (%T)%v for \"%s\"", e.arg, e.arg, e.token)
}

// NewInvalidParamTypeErr create a new [InvalidParamTypeErr]
func NewInvalidParamTypeErr(token string, arg any) *InvalidParamTypeErr {
	return &InvalidParamTypeErr{
		token: token,
		arg:   arg,
	}
}

// ThrowInvalidParamTypeErr create a new [InvalidParamTypeErr] and panic
func ThrowInvalidParamTypeErr(token string, arg any) {
	panic(NewInvalidParamTypeErr(token, arg))
}
