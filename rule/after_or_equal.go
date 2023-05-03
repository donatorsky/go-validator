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
	comparableObj, ok := value.(afterOrEqualComparable)
	if !ok {
		return value, NewAfterOrEqualValidationError(r.afterOrEqual.Format(time.RFC3339Nano))
	}

	if comparableObj.Equal(r.afterOrEqual) || comparableObj.After(r.afterOrEqual) {
		return value, nil
	}

	return value, NewAfterOrEqualValidationError(r.afterOrEqual.Format(time.RFC3339Nano))
}

func NewAfterOrEqualValidationError(afterOrEqual string) AfterOrEqualValidationError {
	return AfterOrEqualValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeAfterOrEqual,
		},
		AfterOrEqual: afterOrEqual,
	}
}

type AfterOrEqualValidationError struct {
	ve.BasicValidationError

	AfterOrEqual string `json:"after_or_equal"`
}

func (e AfterOrEqualValidationError) Error() string {
	return fmt.Sprintf("afterOrEqualRule{After=%q}", e.AfterOrEqual)
}
