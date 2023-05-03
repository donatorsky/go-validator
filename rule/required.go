package rule

import (
	"context"

	ve "github.com/donatorsky/go-validator/error"
)

func Required() *requiredRule {
	return &requiredRule{}
}

type requiredRule struct {
	Bailer
}

func (r *requiredRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	if value == nil {
		r.bailed = true

		return nil, NewRequiredValidationError()
	}

	return value, nil
}

func NewRequiredValidationError() RequiredValidationError {
	return RequiredValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeRequired,
		},
	}
}

type RequiredValidationError struct {
	ve.BasicValidationError
}

func (RequiredValidationError) Error() string {
	return "requiredRule{}"
}
