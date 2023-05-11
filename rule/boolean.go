package rule

import (
	"context"
	"strconv"

	ve "github.com/donatorsky/go-validator/error"
)

func Boolean() *booleanRule {
	return &booleanRule{}
}

type booleanRule struct {
	Bailer
}

func (r *booleanRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return (*bool)(nil), nil
	}

	switch newValue := v.(type) {
	case bool:
		return value, nil

	case string:
		parsedValue, err := strconv.ParseBool(newValue)
		if err != nil {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return parsedValue, nil

	case int:
		if newValue != 0 && newValue != 1 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1, nil

	case int8:
		if newValue != 0 && newValue != 1 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1, nil

	case int16:
		if newValue != 0 && newValue != 1 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1, nil

	case int32:
		if newValue != 0 && newValue != 1 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1, nil

	case int64:
		if newValue != 0 && newValue != 1 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1, nil

	case uint:
		if newValue != 0 && newValue != 1 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1, nil

	case uint8:
		if newValue != 0 && newValue != 1 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1, nil

	case uint16:
		if newValue != 0 && newValue != 1 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1, nil

	case uint32:
		if newValue != 0 && newValue != 1 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1, nil

	case uint64:
		if newValue != 0 && newValue != 1 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1, nil

	case float32:
		if newValue != 0.0 && newValue != 1.0 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1.0, nil

	case float64:
		if newValue != 0.0 && newValue != 1.0 {
			r.MarkBailed()

			return nil, NewBooleanValidationError()
		}

		return newValue == 1.0, nil

	default:
		r.MarkBailed()

		return nil, NewBooleanValidationError()
	}
}

func NewBooleanValidationError() BooleanValidationError {
	return BooleanValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeBoolean,
		},
	}
}

type BooleanValidationError struct {
	ve.BasicValidationError
}

func (e BooleanValidationError) Error() string {
	return "must be true or false"
}
