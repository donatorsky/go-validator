package rule

import (
	"context"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Slice() *sliceRule {
	return &sliceRule{}
}

type sliceRule struct {
	Bailer
}

func (r *sliceRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	if reflect.TypeOf(v).Kind() != reflect.Slice {
		r.MarkBailed()

		return nil, NewSliceValidationError()
	}

	return value, nil
}

func NewSliceValidationError() SliceValidationError {
	return SliceValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleSlice,
		},
	}
}

type SliceValidationError struct {
	ve.BasicValidationError
}

func (SliceValidationError) Error() string {
	return "must be a slice"
}
