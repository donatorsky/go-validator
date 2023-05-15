package rule

import (
	"context"
	"fmt"
	"time"

	ve "github.com/donatorsky/go-validator/error"
)

func DateFormat(format string) *dateFormatRule {
	return &dateFormatRule{
		format: format,
	}
}

type dateFormatRule struct {
	format string
}

func (r dateFormatRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return (*time.Time)(nil), nil
	}

	switch newValue := v.(type) {
	case time.Time:
		return value, nil

	case string:
		parsedDateFormat, err := time.Parse(r.format, newValue)
		if err != nil {
			return value, NewDateFormatValidationError(r.format)
		}

		return parsedDateFormat, nil

	default:
		return value, NewDateFormatValidationError(r.format)
	}
}

func NewDateFormatValidationError(format string) DateFormatValidationError {
	return DateFormatValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleDateFormat,
		},
		Format: format,
	}
}

type DateFormatValidationError struct {
	ve.BasicValidationError

	Format string `json:"format"`
}

func (e DateFormatValidationError) Error() string {
	return fmt.Sprintf("does not match the date format %s", e.Format)
}
