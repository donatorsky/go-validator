package rule

import (
	"context"
	"regexp"

	ve "github.com/donatorsky/go-validator/error"
)

func NotRegex(regex *regexp.Regexp) *notRegexRule {
	return &notRegexRule{
		notRegex: regex,
	}
}

type notRegexRule struct {
	notRegex *regexp.Regexp
}

func (r *notRegexRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	if newValue, ok := v.(string); !ok || r.notRegex.MatchString(newValue) {
		return value, NewNotRegexValidationError()
	}

	return value, nil
}

func NewNotRegexValidationError() NotRegexValidationError {
	return NotRegexValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeNotRegex,
		},
	}
}

type NotRegexValidationError struct {
	ve.BasicValidationError
}

func (e NotRegexValidationError) Error() string {
	return "format is invalid"
}
