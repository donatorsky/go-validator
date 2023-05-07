package rule

import (
	"context"
	"time"

	ve "github.com/donatorsky/go-validator/error"
)

func Duration() *durationRule {
	return &durationRule{}
}

type durationRule struct {
}

func (*durationRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return (*time.Duration)(nil), nil
	}

	switch newValue := v.(type) {
	case time.Duration:
		return value, nil

	case string:
		parsedDuration, err := time.ParseDuration(newValue)
		if err != nil {
			return value, NewDurationValidationError()
		}

		return parsedDuration, nil

	default:
		return value, NewDurationValidationError()
	}
}

func NewDurationValidationError() DurationValidationError {
	return DurationValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeDuration,
		},
	}
}

type DurationValidationError struct {
	ve.BasicValidationError
}

func (e DurationValidationError) Error() string {
	return "must be a valid duration"
}
