package rule

import (
	"context"
	"fmt"

	ve "github.com/donatorsky/go-validator/error"
)

type customValidator[In any, Out any] func(ctx context.Context, value In, data any) (newValue Out, err error)

func Custom[In any, Out any](validator customValidator[In, Out]) *customRule[In, Out] {
	return &customRule[In, Out]{
		validator: validator,
	}
}

type customRule[In any, Out any] struct {
	validator customValidator[In, Out]
}

func (r customRule[In, Out]) Apply(ctx context.Context, value any, data any) (any, ve.ValidationError) {
	inValue, ok := value.(In)
	if !ok {
		return value, NewCustomValidationError(fmt.Errorf("invalid data type provided: %T, expected %T", value, inValue))
	}

	outValue, err := r.validator(ctx, inValue, data)
	if err != nil {
		if vErr, ok := err.(ve.ValidationError); ok {
			return value, vErr
		}

		return value, NewCustomValidationError(err)
	}

	return outValue, nil
}

func NewCustomValidationError(err error) CustomValidationError {
	return CustomValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleCustom,
		},
		Err: err.Error(),
	}
}

type CustomValidationError struct {
	ve.BasicValidationError

	Err string `json:"error"`
}

func (e CustomValidationError) Error() string {
	return e.Err
}
