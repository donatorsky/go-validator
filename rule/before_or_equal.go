package rule

import (
	"context"
	"fmt"
	"time"

	ve "github.com/donatorsky/go-validator/error"
)

func BeforeOrEqual(beforeOrEqual time.Time) *beforeOrEqualRule {
	return &beforeOrEqualRule{
		beforeOrEqual: beforeOrEqual,
	}
}

type beforeOrEqualRule struct {
	beforeOrEqual time.Time
}

type beforeOrEqualComparable interface {
	Equal(time.Time) bool
	Before(time.Time) bool
}

func (r beforeOrEqualRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	if _, isNil := Dereference(value); isNil {
		return value, nil
	}

	if comparableObj, ok := value.(beforeOrEqualComparable); !ok || !comparableObj.Equal(r.beforeOrEqual) && !comparableObj.Before(r.beforeOrEqual) {
		return value, NewBeforeOrEqualValidationError(r.beforeOrEqual.Format(time.RFC3339Nano))
	}

	return value, nil
}

func NewBeforeOrEqualValidationError(beforeOrEqual string) BeforeOrEqualValidationError {
	return BeforeOrEqualValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleBeforeOrEqual,
		},
		BeforeOrEqual: beforeOrEqual,
	}
}

type BeforeOrEqualValidationError struct {
	ve.BasicValidationError

	BeforeOrEqual string `json:"after_or_equal"`
}

func (e BeforeOrEqualValidationError) Error() string {
	return fmt.Sprintf("must be a date before or equal to %s", e.BeforeOrEqual)
}
