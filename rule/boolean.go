package rule

import (
	"context"
	"fmt"
	"strconv"

	ve "github.com/donatorsky/go-validator/error"
)

func Boolean() *booleanRule {
	return &booleanRule{}
}

type booleanRule struct {
}

func (*booleanRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	switch newValue := value.(type) {
	case bool:
		return newValue, nil

	case string:
		value, err := strconv.ParseBool(newValue)
		if err != nil {
			return newValue, NewBooleanValidationError(fmt.Sprintf("%T", value))
		}

		return value, nil

	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		if newValue != 0 && newValue != 1 {
			return value, NewBooleanValidationError(fmt.Sprintf("%T", value))
		}

		return newValue == 1, nil

	case float32, float64:
		if newValue != 0.0 && newValue != 1.0 {
			return value, NewBooleanValidationError(fmt.Sprintf("%T", value))
		}

		return newValue == 1.0, nil

	default:
		return value, NewBooleanValidationError(fmt.Sprintf("%T", value))
	}
}

func NewBooleanValidationError(actual string) BooleanValidationError {
	return BooleanValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeBoolean,
		},
		Actual: actual,
	}
}

type BooleanValidationError struct {
	ve.BasicValidationError

	Actual string `json:"actual"`
}

func (e BooleanValidationError) Error() string {
	return fmt.Sprintf("booleanRule{Actual=%q}", e.Actual)
}
