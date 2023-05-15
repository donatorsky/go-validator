package rule

import (
	"context"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Filled() *filledRule {
	return &filledRule{}
}

type filledRule struct {
}

func (r *filledRule) Apply(ctx context.Context, value any, data any) (any, ve.ValidationError) {
	valueOf := reflect.ValueOf(value)
	if !valueOf.IsValid() {
		return value, nil
	}

	if valueOf.Kind() == reflect.Pointer {
		v, _ := Dereference(value)
		_, validationError := r.Apply(ctx, v, data)

		return value, validationError
	}

	if valueOf.IsZero() {
		return value, NewFilledValidationError()
	}

	return value, nil
}

func NewFilledValidationError() FilledValidationError {
	return FilledValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeFilled,
		},
	}
}

type FilledValidationError struct {
	ve.BasicValidationError
}

func (FilledValidationError) Error() string {
	return "must not be empty"
}
