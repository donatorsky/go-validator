package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func ArrayOf[Out any]() *arrayOfRule[Out] {
	return &arrayOfRule[Out]{}
}

type arrayOfRule[Out any] struct {
	Bailer
}

func (r *arrayOfRule[Out]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return (*[0]Out)(nil), nil
	}

	var el Out

	if typeOf := reflect.TypeOf(v); typeOf.Kind() != reflect.Array || !typeOf.Elem().AssignableTo(reflect.TypeOf(el)) {
		r.MarkBailed()

		return nil, NewArrayOfValidationError(
			fmt.Sprintf("%T", el),
			fmt.Sprintf("%T", v),
		)
	}

	return value, nil
}

func NewArrayOfValidationError(expected, actual string) ArrayOfValidationError {
	return ArrayOfValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeArrayOf,
		},
		ExpectedType: expected,
		ActualType:   actual,
	}
}

type ArrayOfValidationError struct {
	ve.BasicValidationError

	ExpectedType string `json:"expected_type"`
	ActualType   string `json:"actual_type"`
}

func (e ArrayOfValidationError) Error() string {
	return fmt.Sprintf("must be an array of %q, but is %q", e.ExpectedType, e.ActualType)
}
