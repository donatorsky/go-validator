package rule

import (
	"context"
	"fmt"
	"strings"

	ve "github.com/donatorsky/go-validator/error"
)

func StartsWith(prefix string, prefixes ...string) *startsWithRule {
	rule := startsWithRule{
		prefixes: make([]string, 0, len(prefixes)+1),
	}

	rule.prefixes = append(rule.prefixes, prefix)
	rule.prefixes = append(rule.prefixes, prefixes...)

	return &rule
}

type startsWithRule struct {
	prefixes []string
}

func (r *startsWithRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	stringValue, ok := v.(string)
	if !ok {
		return value, NewStartsWithValidationError(r.prefixes)
	}

	for _, prefix := range r.prefixes {
		if strings.HasPrefix(stringValue, prefix) {
			return value, nil
		}
	}

	return value, NewStartsWithValidationError(r.prefixes)
}

func NewStartsWithValidationError(prefixes []string) StartsWithValidationError {
	return StartsWithValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleStartsWith,
		},
		Prefixes: prefixes,
	}
}

type StartsWithValidationError struct {
	ve.BasicValidationError

	Prefixes []string `json:"prefixes"`
}

func (e StartsWithValidationError) Error() string {
	return fmt.Sprintf("must start with one of the following: %s", strings.Join(e.Prefixes, "; "))
}
