package rule

import (
	"context"
	"net/url"

	ve "github.com/donatorsky/go-validator/error"
)

func URL() *urlRule {
	return &urlRule{}
}

type urlRule struct {
}

func (*urlRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	stringValue, ok := v.(string)
	if !ok {
		return value, NewUrlValidationError()
	}

	if _, err := url.ParseRequestURI(stringValue); err != nil {
		return value, NewUrlValidationError()
	}

	return value, nil
}

func NewUrlValidationError() UrlValidationError {
	return UrlValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeURL,
		},
	}
}

type UrlValidationError struct {
	ve.BasicValidationError
}

func (e UrlValidationError) Error() string {
	return "must be a valid URL format"
}
