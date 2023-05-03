package rule

import (
	"context"
	"fmt"

	ve "github.com/donatorsky/go-validator/error"
)

type integerType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Integer[Out integerType]() *integerRule[Out] {
	return &integerRule[Out]{}
}

type integerRule[Out integerType] struct {
}

func (integerRule[Out]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return (*Out)(nil), nil
	}

	newValue, ok := v.(Out)
	if !ok {
		return value, NewIntegerValidationError(
			fmt.Sprintf("%T", newValue),
			fmt.Sprintf("%T", v),
		)
	}

	return newValue, nil
}

func NewIntegerValidationError(expectedType, actualType string) IntegerValidationError {
	return IntegerValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeInt,
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
		"integerRule{ExpectedType=%q, ActualType=%q}",
		e.ExpectedType,
		e.ActualType,
	)
}
