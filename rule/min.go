package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Min[T numberType](min T) *minRule[T] {
	return &minRule[T]{
		min: min,
	}
}

type minRule[T numberType] struct {
	min T
}

func (r *minRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	switch v := v.(type) {
	case string:
		if CompareNumbers(len(v), r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeString, r.min)
		}

	case int:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case int8:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case int16:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case int32:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case int64:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case uint:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case uint8:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case uint16:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case uint32:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case uint64:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case float32:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	case float64:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.SubtypeNumber, r.min)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice:
			if CompareNumbers(valueOf.Len(), r.min) == -1 {
				return value, NewMinValidationError(ve.SubtypeSlice, r.min)
			}

		case reflect.Array:
			if CompareNumbers(valueOf.Len(), r.min) == -1 {
				return value, NewMinValidationError(ve.SubtypeArray, r.min)
			}

		case reflect.Map:
			if CompareNumbers(valueOf.Len(), r.min) == -1 {
				return value, NewMinValidationError(ve.SubtypeMap, r.min)
			}
		}
	}

	return value, nil
}

func NewMinValidationError[T numberType](st string, threshold T) MinValidationError[T] {
	return MinValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: fmt.Sprintf("%s.%s", ve.TypeMin, st),
		},
		Threshold: threshold,
	}
}

type MinValidationError[T numberType] struct {
	ve.BasicValidationError

	Threshold T `json:"threshold"`
}

func (e MinValidationError[T]) Error() string {
	switch e.Rule {
	case ve.TypeMin + "." + ve.SubtypeString:
		return fmt.Sprintf("must be at least %v characters", e.Threshold)

	case ve.TypeMin + "." + ve.SubtypeSlice,
		ve.TypeMin + "." + ve.SubtypeArray,
		ve.TypeMin + "." + ve.SubtypeMap:
		return fmt.Sprintf("must have at least %v items", e.Threshold)

	default:
		return fmt.Sprintf("must be at least %v", e.Threshold)
	}
}
