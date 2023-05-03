package rule

import (
	"context"
	"net/mail"

	ve "github.com/donatorsky/go-validator/error"
)

func Email() *emailRule {
	return &emailRule{}
}

type emailRule struct {
}

func (*emailRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	newValue, ok := v.(string)
	if !ok {
		return value, NewEmailValidationError()
	}

	_, err := mail.ParseAddress(newValue)
	if err != nil {
		return value, NewEmailValidationError()
	}

	return value, nil
}

func NewEmailValidationError() EmailValidationError {
	return EmailValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeEmail,
		},
	}
}

type EmailValidationError struct {
	ve.BasicValidationError
}

func (e EmailValidationError) Error() string {
	return "emailRule{}"
}
