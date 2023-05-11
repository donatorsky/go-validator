package rule

import (
	"context"
	"fmt"

	ve "github.com/donatorsky/go-validator/error"
)

func Float[Out floatType]() *floatRule[Out] {
	return &floatRule[Out]{}
}

type floatRule[Out floatType] struct {
	Bailer
}

func (r *floatRule[Out]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return (*Out)(nil), nil
	}

	if newValue, ok := v.(Out); !ok {
		r.MarkBailed()

		return nil, NewFloatValidationError(
			fmt.Sprintf("%T", newValue),
			fmt.Sprintf("%T", v),
		)
	}

	return value, nil
}

func NewFloatValidationError(expectedType, actualType string) FloatValidationError {
	return FloatValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeInt,
		},
		ExpectedType: expectedType,
		ActualType:   actualType,
	}
}

type FloatValidationError struct {
	ve.BasicValidationError

	ExpectedType string `json:"expected_type"`
	ActualType   string `json:"actual_type"`
}

func (e FloatValidationError) Error() string {
	return fmt.Sprintf(
		"must be a %s but is %s",
		e.ExpectedType,
		e.ActualType,
	)
}
