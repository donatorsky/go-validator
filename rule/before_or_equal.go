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
	comparableObj, ok := value.(beforeOrEqualComparable)
	if !ok {
		return value, NewBeforeOrEqualValidationError(r.beforeOrEqual.Format(time.RFC3339Nano))
	}

	if comparableObj.Equal(r.beforeOrEqual) || comparableObj.Before(r.beforeOrEqual) {
		return value, nil
	}

	return value, NewBeforeOrEqualValidationError(r.beforeOrEqual.Format(time.RFC3339Nano))
}

func NewBeforeOrEqualValidationError(beforeOrEqual string) BeforeOrEqualValidationError {
	return BeforeOrEqualValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeBeforeOrEqual,
		},
		BeforeOrEqual: beforeOrEqual,
	}
}

type BeforeOrEqualValidationError struct {
	ve.BasicValidationError

	BeforeOrEqual string `json:"after_or_equal"`
}

func (e BeforeOrEqualValidationError) Error() string {
	return fmt.Sprintf("beforeOrEqualRule{After=%q}", e.BeforeOrEqual)
}
