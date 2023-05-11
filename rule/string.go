package rule

import (
	"context"

	ve "github.com/donatorsky/go-validator/error"
)

func String() *stringRule {
	return &stringRule{}
}

type stringRule struct {
	Bailer
}

func (r *stringRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return (*string)(nil), nil
	}

	if _, ok := v.(string); !ok {
		r.MarkBailed()

		return nil, NewStringValidationError()
	}

	return value, nil
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
	return "must be a string"
}
