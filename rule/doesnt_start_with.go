package rule

import (
	"context"
	"fmt"
	"strings"

	ve "github.com/donatorsky/go-validator/error"
)

func DoesntStartWith(prefix string, prefixes ...string) *doesntStartWithRule {
	rule := doesntStartWithRule{
		prefixes: make([]string, 0, len(prefixes)+1),
	}

	rule.prefixes = append(rule.prefixes, prefix)
	rule.prefixes = append(rule.prefixes, prefixes...)

	return &rule
}

type doesntStartWithRule struct {
	prefixes []string
}

func (r *doesntStartWithRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	stringValue, ok := v.(string)
	if !ok {
		return value, NewDoesntStartWithValidationError(r.prefixes)
	}

	for _, prefix := range r.prefixes {
		if strings.HasPrefix(stringValue, prefix) {
			return value, NewDoesntStartWithValidationError(r.prefixes)
		}
	}

	return value, nil
}

func NewDoesntStartWithValidationError(prefixes []string) DoesntStartWithValidationError {
	return DoesntStartWithValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeDoesntStartWith,
		},
		Prefixes: prefixes,
	}
}

type DoesntStartWithValidationError struct {
	ve.BasicValidationError

	Prefixes []string `json:"prefixes"`
}

func (e DoesntStartWithValidationError) Error() string {
	return fmt.Sprintf("must not start with any of the following: %s", strings.Join(e.Prefixes, "; "))
}
