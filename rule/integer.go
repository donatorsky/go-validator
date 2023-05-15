package rule

import (
	"context"
	"fmt"

	ve "github.com/donatorsky/go-validator/error"
)

func Integer[Out integerType]() *integerRule[Out] {
	return &integerRule[Out]{}
}

type integerRule[Out integerType] struct {
	Bailer
}

func (r *integerRule[Out]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return (*Out)(nil), nil
	}

	if newValue, ok := v.(Out); !ok {
		r.MarkBailed()

		return nil, NewIntegerValidationError(
			fmt.Sprintf("%T", newValue),
			fmt.Sprintf("%T", v),
		)
	}

	return value, nil
}

func NewIntegerValidationError(expectedType, actualType string) IntegerValidationError {
	return IntegerValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleInt,
		},
		ExpectedType: expectedType,
		ActualType:   actualType,
	}
}

type IntegerValidationError struct {
	ve.BasicValidationError

	ExpectedType string `json:"expected_type"`
	ActualType   string `json:"actual_type"`
}

func (e IntegerValidationError) Error() string {
	return fmt.Sprintf(
		"must be an %s but is %s",
		e.ExpectedType,
		e.ActualType,
	)
}
