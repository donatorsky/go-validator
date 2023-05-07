package rule

import (
	"context"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Struct() *structRule {
	return &structRule{}
}

type structRule struct {
	Bailer
}

func (r *structRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return nil, nil
	}

	if reflect.TypeOf(v).Kind() != reflect.Struct {
		r.MarkBailed()

		return value, NewStructValidationError()
	}

	return value, nil
}

func NewStructValidationError() StructValidationError {
	return StructValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeStruct,
		},
	}
}

type StructValidationError struct {
	ve.BasicValidationError
}

func (StructValidationError) Error() string {
	return "must be a struct"
}
