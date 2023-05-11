package rule

import (
	"context"
	"fmt"

	ve "github.com/donatorsky/go-validator/error"
)

func SliceOf[Out any]() *sliceOfRule[Out] {
	return &sliceOfRule[Out]{}
}

type sliceOfRule[Out any] struct {
	Bailer
}

func (r *sliceOfRule[Out]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return ([]Out)(nil), nil
	}

	if _, ok := v.([]Out); !ok {
		r.MarkBailed()

		var el Out
		return nil, NewSliceOfValidationError(
			fmt.Sprintf("%T", el),
			fmt.Sprintf("%T", v),
		)
	}

	return value, nil
}

func NewSliceOfValidationError(expected, actual string) SliceOfValidationError {
	return SliceOfValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeSliceOf,
		},
		ExpectedType: expected,
		ActualType:   actual,
	}
}

type SliceOfValidationError struct {
	ve.BasicValidationError

	ExpectedType string `json:"expected_type"`
	ActualType   string `json:"actual_type"`
}

func (e SliceOfValidationError) Error() string {
	return fmt.Sprintf("must be a slice of %q, but is %q", e.ExpectedType, e.ActualType)
}
