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
}

func (*sliceRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	if value == nil {
		return value, nil
	}

	if kind := reflect.TypeOf(value); kind != nil && kind.Kind() == reflect.Slice {
		return value, nil
	}

	return value, NewSliceValidationError()
}

func NewSliceValidationError() SliceValidationError {
	return SliceValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeSlice,
		},
	}
}

type SliceValidationError struct {
	ve.BasicValidationError
}

func (SliceValidationError) Error() string {
	return "sliceRule{}"
}
