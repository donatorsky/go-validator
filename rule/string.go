package rule

import (
	"context"

	ve "github.com/donatorsky/go-validator/error"
)

func String() *stringRule {
	return &stringRule{}
}

type stringRule struct {
}

func (*stringRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return (*string)(nil), nil
	}

	newValue, ok := v.(string)
	if !ok {
		return value, NewStringValidationError()
	}

	return newValue, nil
}

func NewStringValidationError() StringValidationError {
	return StringValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeString,
		},
	}
}

type StringValidationError struct {
	ve.BasicValidationError
}

func (StringValidationError) Error() string {
	return "stringRule{}"
}
