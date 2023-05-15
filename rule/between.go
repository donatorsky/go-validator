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
		if !isBetween(len(v), r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeString, r.min, r.max, r.inclusive)
		}

	case int:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case int8:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case int16:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case int32:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case int64:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case uint:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case uint8:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case uint16:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case uint32:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case uint64:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case float32:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	case float64:
		if !isBetween(v, r.min, r.max, r.inclusive) {
			return value, NewBetweenValidationError(ve.SubtypeNumber, r.min, r.max, r.inclusive)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice:
			if !isBetween(valueOf.Len(), r.min, r.max, r.inclusive) {
				return value, NewBetweenValidationError(ve.SubtypeSlice, r.min, r.max, r.inclusive)
			}

		case reflect.Array:
			if !isBetween(valueOf.Len(), r.min, r.max, r.inclusive) {
				return value, NewBetweenValidationError(ve.SubtypeArray, r.min, r.max, r.inclusive)
			}

		case reflect.Map:
			if !isBetween(valueOf.Len(), r.min, r.max, r.inclusive) {
				return value, NewBetweenValidationError(ve.SubtypeMap, r.min, r.max, r.inclusive)
			}

		default:
			return value, NewBetweenValidationError(ve.SubtypeInvalid, r.min, r.max, r.inclusive)
		}
	}

	return value, nil
}

func isBetween[V, T numberType](v V, min, max T, inclusive bool) bool {
	if inclusive {
		// v >= min && v <= max
		return CompareNumbers(v, min) != -1 && CompareNumbers(v, max) != 1
	}

	// v > min && v < max
	return CompareNumbers(v, min) == 1 && CompareNumbers(v, max) == -1
}

func NewBetweenValidationError[T numberType](st string, min, max T, inclusive bool) BetweenValidationError[T] {
	return BetweenValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeBetween,
		},
		Type:      st,
		Min:       min,
		Max:       max,
		Inclusive: inclusive,
	}
}

type BetweenValidationError[T numberType] struct {
	ve.BasicValidationError

	Type      string `json:"type"`
	Min       T      `json:"min"`
	Max       T      `json:"max"`
	Inclusive bool   `json:"inclusive"`
}

func (e BetweenValidationError[T]) Error() string {
	var mode string
	if e.Inclusive {
		mode = "inclusive"
	} else {
		mode = "exclusive"
	}

	switch e.Type {
	case ve.SubtypeNumber:
		return fmt.Sprintf("must be between %v and %v (%s)", e.Min, e.Max, mode)

	case ve.SubtypeString:
		return fmt.Sprintf("must be between %v and %v characters (%s)", e.Min, e.Max, mode)

	case ve.SubtypeSlice,
		ve.SubtypeArray,
		ve.SubtypeMap:
		return fmt.Sprintf("must have between %v and %v items (%s)", e.Min, e.Max, mode)

	default:
		return "cannot be determined"
	}
}
