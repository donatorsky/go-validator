package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Between[T numberType](min, max T) *betweenRule[T] {
	return &betweenRule[T]{
		min:       min,
		max:       max,
		inclusive: true,
	}
}

type betweenRule[T numberType] struct {
	min       T
	max       T
	inclusive bool
}

func (r *betweenRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	switch v := v.(type) {
	case string:
		if isBetween(len(v), r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeString, r.min, r.max, r.inclusive)
		}

	case int:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case int8:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case int16:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case int32:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case int64:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case uint:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case uint8:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case uint16:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case uint32:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case uint64:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case float32:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case float64:
		if isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice, reflect.Array:
			if isBetween(valueOf.Len(), r.min, r.max, r.inclusive) {
				return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeSlice, r.min, r.max, r.inclusive)
			}

		case reflect.Map:
			if isBetween(valueOf.Len(), r.min, r.max, r.inclusive) {
				return value, NewBetweenValidationError(ve.TypeBetween, ve.SubtypeMap, r.min, r.max, r.inclusive)
			}
		}
	}

	return value, nil
}

func isBetween[V, T numberType](v V, min, max T, inclusive bool) bool {
	if inclusive {
		return CompareNumbers(v, min) == -1 || CompareNumbers(v, max) == 1
	}

	return CompareNumbers(v, min) < 1 || CompareNumbers(v, max) > -1
}

func NewBetweenValidationError[T numberType](t, st string, min, max T, inclusive bool) BetweenValidationError[T] {
	return BetweenValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: fmt.Sprintf("%s.%s", t, st),
		},
		Min:       min,
		Max:       max,
		Inclusive: inclusive,
	}
}

type BetweenValidationError[T numberType] struct {
	ve.BasicValidationError

	Min       T    `json:"min"`
	Max       T    `json:"max"`
	Inclusive bool `json:"inclusive"`
}

func (e BetweenValidationError[T]) Error() string {
	var mode string
	if e.Inclusive {
		mode = "inclusive"
	} else {
		mode = "exclusive"
	}

	switch e.Rule {
	case ve.TypeBetween + "." + ve.SubtypeString:
		return fmt.Sprintf("must be between %v and %v characters (%s)", e.Min, e.Max, mode)

	case ve.TypeBetween + "." + ve.SubtypeSlice,
		ve.TypeBetween + "." + ve.SubtypeMap:
		return fmt.Sprintf("must have between %v and %v items (%s)", e.Min, e.Max, mode)

	default:
		return fmt.Sprintf("must be between %v and %v (%s)", e.Min, e.Max, mode)
	}
}
