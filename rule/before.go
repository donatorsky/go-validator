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
	comparableObj, ok := value.(beforeComparable)
	if !ok {
		return value, NewBeforeValidationError(r.before.Format(time.RFC3339Nano))
	}

	if !comparableObj.Before(r.before) {
		return value, NewBeforeValidationError(r.before.Format(time.RFC3339Nano))
	}

	return value, nil
}

func NewBeforeValidationError(before string) BeforeValidationError {
	return BeforeValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeBefore,
		},
		Before: before,
	}
}

type BeforeValidationError struct {
	ve.BasicValidationError

	Before string `json:"before"`
}

func (e BeforeValidationError) Error() string {
	return fmt.Sprintf("beforeRule{After=%q}", e.Before)
}
