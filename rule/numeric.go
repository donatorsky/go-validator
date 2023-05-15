package rule

import (
	"context"
	"strconv"

	ve "github.com/donatorsky/go-validator/error"
)

func Numeric() *numericRule {
	return &numericRule{}
}

type numericRule struct {
}

func (*numericRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	switch newValue := v.(type) {
	case string:
		if newValue, err := strconv.ParseInt(newValue, 10, 64); err == nil {
			return newValue, nil
		}

		if newValue, err := strconv.ParseUint(newValue, 10, 64); err == nil {
			return newValue, nil
		}

		if newValue, err := strconv.ParseFloat(newValue, 64); err == nil {
			return newValue, nil
		}

		if newValue, err := strconv.ParseComplex(newValue, 128); err == nil {
			return newValue, nil
		}

	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128:
		return value, nil
	}

	return value, NewNumericValidationError()
}

func NewNumericValidationError() NumericValidationError {
	return NumericValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeNumeric,
		},
	}
}

type NumericValidationError struct {
	ve.BasicValidationError
}

func (e NumericValidationError) Error() string {
	return "must be a number"
}
