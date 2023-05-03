package rule

import (
	"context"
	"fmt"

	ve "github.com/donatorsky/go-validator/error"
)

type floatType interface {
	~float32 | ~float64
}

func Float[Out floatType]() *floatRule[Out] {
	return &floatRule[Out]{}
}

type floatRule[Out floatType] struct {
}

func (floatRule[Out]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return (*Out)(nil), nil
	}

	newValue, ok := v.(Out)
	if !ok {
		return value, NewFloatValidationError(
			fmt.Sprintf("%T", newValue),
			fmt.Sprintf("%T", v),
		)
	}

	return newValue, nil
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
		"floatRule{ExpectedType=%q, ActualType=%q}",
		e.ExpectedType,
		e.ActualType,
	)
}
