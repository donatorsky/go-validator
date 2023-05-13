package rule

import (
	"context"
	"regexp"

	ve "github.com/donatorsky/go-validator/error"
)

func Regex(regex *regexp.Regexp) *regexRule {
	return &regexRule{
		regex: regex,
	}
}

type regexRule struct {
	regex *regexp.Regexp
}

func (r *regexRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	newValue, ok := v.(string)
	if !ok {
		return value, NewRegexValidationError()
	}

	if !r.regex.MatchString(newValue) {
		return value, NewRegexValidationError()
	}

	return value, nil
}

func NewRegexValidationError() RegexValidationError {
	return RegexValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeRegex,
		},
	}
}

type RegexValidationError struct {
	ve.BasicValidationError
}

func (e RegexValidationError) Error() string {
	return "format is invalid"
}
