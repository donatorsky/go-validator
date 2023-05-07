package rule

import (
	"context"
	"fmt"
	"time"

	ve "github.com/donatorsky/go-validator/error"
)

func After(after time.Time) *afterRule {
	return &afterRule{
		after: after,
	}
}

type afterRule struct {
	after time.Time
}

type afterComparable interface {
	After(time.Time) bool
}

func (r afterRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	if _, isNil := Dereference(value); isNil {
		return value, nil
	}

	if comparableObj, ok := value.(afterComparable); !ok || !comparableObj.After(r.after) {
		return value, NewAfterValidationError(r.after.Format(time.RFC3339Nano))
	}

	return value, nil
}

func NewAfterValidationError(after string) AfterValidationError {
	return AfterValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeAfter,
		},
		After: after,
	}
}

type AfterValidationError struct {
	ve.BasicValidationError

	After string `json:"after"`
}

func (e AfterValidationError) Error() string {
	return fmt.Sprintf("must be a date after %s", e.After)
}
