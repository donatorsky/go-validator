package rule

import (
	"context"
	"net/mail"

	ve "github.com/donatorsky/go-validator/error"
)

func EmailAddress() *emailAddressRule {
	return &emailAddressRule{}
}

type emailAddressRule struct {
}

func (*emailAddressRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	newValue, ok := v.(string)
	if !ok {
		return value, NewEmailValidationError()
	}

	email, err := mail.ParseAddress(newValue)
	if err != nil {
		return value, NewEmailValidationError()
	}

	return email.Address, nil
}
