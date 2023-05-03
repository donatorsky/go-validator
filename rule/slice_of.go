package rule

import (
	"context"
	"fmt"

	ve "github.com/donatorsky/go-validator/error"
)

func SliceOf[T any]() *sliceOfRule[T] {
	return &sliceOfRule[T]{}
}

type sliceOfRule[T any] struct {
}

func (*sliceOfRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	if value == nil {
		return value, nil
	}

	if newValue, ok := value.([]T); ok {
		return newValue, nil
	}

	var el T
	return value, NewSliceOfValidationError(fmt.Sprintf("%T", el), fmt.Sprintf("%T", value))
}

func NewSliceOfValidationError(expected, actual string) SliceOfValidationError {
	return SliceOfValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeSliceOf,
		},
		Expected: expected,
		Actual:   actual,
	}
}

type SliceOfValidationError struct {
	ve.BasicValidationError

	Expected string `json:"expected"`
	Actual   string `json:"actual"`
}

func (e SliceOfValidationError) Error() string {
	return fmt.Sprintf("sliceOfRule{Expected=%q, Actual=%q}", e.Expected, e.Actual)
}
