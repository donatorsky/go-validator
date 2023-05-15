package rule

import (
	"context"
	"fmt"
	"strings"

	ve "github.com/donatorsky/go-validator/error"
)

func DoesntEndWith(suffix string, suffixes ...string) *doesntEndWithRule {
	rule := doesntEndWithRule{
		suffixes: make([]string, 0, len(suffixes)+1),
	}

	rule.suffixes = append(rule.suffixes, suffix)
	rule.suffixes = append(rule.suffixes, suffixes...)

	return &rule
}

type doesntEndWithRule struct {
	suffixes []string
}

func (r *doesntEndWithRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	stringValue, ok := v.(string)
	if !ok {
		return value, NewDoesntEndWithValidationError(r.suffixes)
	}

	for _, suffix := range r.suffixes {
		if strings.HasSuffix(stringValue, suffix) {
			return value, NewDoesntEndWithValidationError(r.suffixes)
		}
	}

	return value, nil
}

func NewDoesntEndWithValidationError(suffixes []string) DoesntEndWithValidationError {
	return DoesntEndWithValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleDoesntEndWith,
		},
		Suffixes: suffixes,
	}
}

type DoesntEndWithValidationError struct {
	ve.BasicValidationError

	Suffixes []string `json:"suffixes"`
}

func (e DoesntEndWithValidationError) Error() string {
	return fmt.Sprintf("must not end with any of the following: %s", strings.Join(e.Suffixes, "; "))
}
