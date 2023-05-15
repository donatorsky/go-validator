package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Min[T numberType](min T) *minRule[T] {
	return &minRule[T]{
		min:       min,
		inclusive: true,
	}
}

type minRule[T numberType] struct {
	min       T
	inclusive bool
}

func (r *minRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	switch v := v.(type) {
	case string:
		if !isMin(len(v), r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeString, r.min, r.inclusive)
		}

	case int:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case int8:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case int16:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case int32:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case int64:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case uint:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case uint8:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case uint16:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case uint32:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case uint64:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case float32:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	case float64:
		if !isMin(v, r.min, r.inclusive) {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min, r.inclusive)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice:
			if !isMin(valueOf.Len(), r.min, r.inclusive) {
				return value, NewMinValidationError(ve.SubtypeSlice, r.min, r.inclusive)
			}

		case reflect.Array:
			if !isMin(valueOf.Len(), r.min, r.inclusive) {
				return value, NewMinValidationError(ve.SubtypeArray, r.min, r.inclusive)
			}

		case reflect.Map:
			if !isMin(valueOf.Len(), r.min, r.inclusive) {
				return value, NewMinValidationError(ve.SubtypeMap, r.min, r.inclusive)
			}

		default:
			return value, NewMinValidationError(ve.SubtypeInvalid, r.min, r.inclusive)
		}
	}

	return value, nil
}

func isMin[V, T numberType](v V, min T, inclusive bool) bool {
	if inclusive {
		// v >= min
		return CompareNumbers(v, min) != -1
	}

	// v > min
	return CompareNumbers(v, min) == 1
}

func NewMinValidationError[T numberType](st string, threshold T, inclusive bool) MinValidationError[T] {
	return MinValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeMin,
		},
		Type:      st,
		Threshold: threshold,
		Inclusive: inclusive,
	}
}

type MinValidationError[T numberType] struct {
	ve.BasicValidationError

	Type      string `json:"type"`
	Threshold T      `json:"threshold"`
	Inclusive bool   `json:"inclusive"`
}

func (e MinValidationError[T]) Error() string {
	switch e.Type {
	case ve.SubtypeNumber:
		if e.Inclusive {
			return fmt.Sprintf("must be at least %v", e.Threshold)
		} else {
			return fmt.Sprintf("must be greater than %v", e.Threshold)
		}

	case ve.SubtypeString:
		if e.Inclusive {
			return fmt.Sprintf("must be at least %v characters", e.Threshold)
		} else {
			return fmt.Sprintf("must be more than %v characters", e.Threshold)
		}

	case ve.SubtypeSlice,
		ve.SubtypeArray,
		ve.SubtypeMap:
		if e.Inclusive {
			return fmt.Sprintf("must have at least %v items", e.Threshold)
		} else {
			return fmt.Sprintf("must have more than %v items", e.Threshold)
		}

	default:
		return "cannot be determined"
	}
}
