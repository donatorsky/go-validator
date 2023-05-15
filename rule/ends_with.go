package rule

import (
	"context"
	"fmt"
	"strings"

	ve "github.com/donatorsky/go-validator/error"
)

func EndsWith(suffix string, suffixes ...string) *endsWithRule {
	rule := endsWithRule{
		suffixes: make([]string, 0, len(suffixes)+1),
	}

	rule.suffixes = append(rule.suffixes, suffix)
	rule.suffixes = append(rule.suffixes, suffixes...)

	return &rule
}

type endsWithRule struct {
	suffixes []string
}

func (r *endsWithRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	stringValue, ok := v.(string)
	if !ok {
		return value, NewEndsWithValidationError(r.suffixes)
	}

	for _, suffix := range r.suffixes {
		if strings.HasSuffix(stringValue, suffix) {
			return value, nil
		}
	}

	return value, NewEndsWithValidationError(r.suffixes)
}

func NewEndsWithValidationError(suffixes []string) EndsWithValidationError {
	return EndsWithValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleEndsWith,
		},
		Suffixes: suffixes,
	}
}

type EndsWithValidationError struct {
	ve.BasicValidationError

	Suffixes []string `json:"suffixes"`
}

func (e EndsWithValidationError) Error() string {
	return fmt.Sprintf("must end with one of the following: %s", strings.Join(e.Suffixes, "; "))
}
