package rule

import (
	"context"
	"fmt"
	"time"

	ve "github.com/donatorsky/go-validator/error"
)

func AfterOrEqual(afterOrEqual time.Time) *afterOrEqualRule {
	return &afterOrEqualRule{
		afterOrEqual: afterOrEqual,
	}
}

type afterOrEqualRule struct {
	afterOrEqual time.Time
}

type afterOrEqualComparable interface {
	Equal(time.Time) bool
	After(time.Time) bool
}

func (r afterOrEqualRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	if _, isNil := Dereference(value); isNil {
		return value, nil
	}

	if comparableObj, ok := value.(afterOrEqualComparable); !ok || !comparableObj.Equal(r.afterOrEqual) && !comparableObj.After(r.afterOrEqual) {
		return value, NewAfterOrEqualValidationError(r.afterOrEqual.Format(time.RFC3339Nano))
	}

	return value, nil
}

func NewAfterOrEqualValidationError(afterOrEqual string) AfterOrEqualValidationError {
	return AfterOrEqualValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleAfterOrEqual,
		},
		AfterOrEqual: afterOrEqual,
	}
}

type AfterOrEqualValidationError struct {
	ve.BasicValidationError

	AfterOrEqual string `json:"after_or_equal"`
}

func (e AfterOrEqualValidationError) Error() string {
	return fmt.Sprintf("must be a date after or equal to %s", e.AfterOrEqual)
}
