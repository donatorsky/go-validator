package rule

import (
	"context"
	"regexp"

	ve "github.com/donatorsky/go-validator/error"
)

func NotRegex(regex *regexp.Regexp) *notRegexRule {
	return &notRegexRule{
		regex: regex,
	}
}

type notRegexRule struct {
	regex *regexp.Regexp
}

func (r *notRegexRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	if stringValue, ok := v.(string); !ok || r.regex.MatchString(stringValue) {
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
