package rule

import (
	"context"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Array() *arrayRule {
	return &arrayRule{}
}

type arrayRule struct {
	Bailer
}

func (r *arrayRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	if reflect.TypeOf(v).Kind() != reflect.Array {
		r.MarkBailed()

		return value, NewArrayValidationError()
	}

	return value, nil
}

func NewArrayValidationError() ArrayValidationError {
	return ArrayValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeArray,
		},
	}
}

type ArrayValidationError struct {
	ve.BasicValidationError
}

func (ArrayValidationError) Error() string {
	return "must be an array"
}
