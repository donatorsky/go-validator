package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Max[T numberType](max T) *maxRule[T] {
	return &maxRule[T]{
		max:       max,
		inclusive: true,
	}
}

type maxRule[T numberType] struct {
	max       T
	inclusive bool
}

func (r *maxRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	switch v := v.(type) {
	case string:
		if !isMax(len(v), r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeString, r.max, r.inclusive)
		}

	case int:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case int8:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case int16:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case int32:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case int64:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case uint:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case uint8:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case uint16:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case uint32:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case uint64:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case float32:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	case float64:
		if !isMax(v, r.max, r.inclusive) {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max, r.inclusive)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice:
			if !isMax(valueOf.Len(), r.max, r.inclusive) {
				return value, NewMaxValidationError(ve.SubtypeSlice, r.max, r.inclusive)
			}

		case reflect.Array:
			if !isMax(valueOf.Len(), r.max, r.inclusive) {
				return value, NewMaxValidationError(ve.SubtypeArray, r.max, r.inclusive)
			}

		case reflect.Map:
			if !isMax(valueOf.Len(), r.max, r.inclusive) {
				return value, NewMaxValidationError(ve.SubtypeMap, r.max, r.inclusive)
			}

		default:
			return value, NewMaxValidationError(ve.SubtypeInvalid, r.max, r.inclusive)
		}
	}

	return value, nil
}

func isMax[V, T numberType](v V, min T, inclusive bool) bool {
	if inclusive {
		// v <= max
		return CompareNumbers(v, min) != 1
	}

	// v < max
	return CompareNumbers(v, min) == -1
}

func NewMaxValidationError[T numberType](st string, threshold T, inclusive bool) MaxValidationError[T] {
	return MaxValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeMax,
		},
		Type:      st,
		Threshold: threshold,
		Inclusive: inclusive,
	}
}

type MaxValidationError[T numberType] struct {
	ve.BasicValidationError

	Type      string `json:"type"`
	Threshold T      `json:"threshold"`
	Inclusive bool   `json:"inclusive"`
}

func (e MaxValidationError[T]) Error() string {
	switch e.Type {
	case ve.SubtypeNumber:
		if e.Inclusive {
			return fmt.Sprintf("must be at most %v", e.Threshold)
		} else {
			return fmt.Sprintf("must be less than %v", e.Threshold)
		}

	case ve.SubtypeString:
		if e.Inclusive {
			return fmt.Sprintf("must be at most %v characters", e.Threshold)
		} else {
			return fmt.Sprintf("must be less than %v characters", e.Threshold)
		}

	case ve.SubtypeSlice,
		ve.SubtypeArray,
		ve.SubtypeMap:
		if e.Inclusive {
			return fmt.Sprintf("must have at most %v items", e.Threshold)
		} else {
			return fmt.Sprintf("must have less than %v items", e.Threshold)
		}

	default:
		return "max cannot be determined"
	}
}
