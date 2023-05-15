package rule

import (
	"context"
	"fmt"
	"time"

	ve "github.com/donatorsky/go-validator/error"
)

func Before(before time.Time) *beforeRule {
	return &beforeRule{
		before: before,
	}
}

type beforeRule struct {
	before time.Time
}

type beforeComparable interface {
	Before(time.Time) bool
}

func (r beforeRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	if _, isNil := Dereference(value); isNil {
		return value, nil
	}

	if comparableObj, ok := value.(beforeComparable); !ok || !comparableObj.Before(r.before) {
		return value, NewBeforeValidationError(r.before.Format(time.RFC3339Nano))
	}

	return value, nil
}

func NewBeforeValidationError(before string) BeforeValidationError {
	return BeforeValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleBefore,
		},
		Before: before,
	}
}

type BeforeValidationError struct {
	ve.BasicValidationError

	Before string `json:"before"`
}

func (e BeforeValidationError) Error() string {
	return fmt.Sprintf("must be a date before %s", e.Before)
}
