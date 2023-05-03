package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Min[T number](min T) *minRule[T] {
	return &minRule[T]{
		min: min,
	}
}

type minRule[T number] struct {
	min T
}

func (r minRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	switch v := v.(type) {
	case string:
		if CompareNumbers(len(v), r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeString, r.min)
		}

	case int:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case int8:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case int16:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case int32:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case int64:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case uint:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case uint8:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case uint16:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case uint32:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case uint64:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case float32:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	case float64:
		if CompareNumbers(v, r.min) == -1 {
			return value, NewMinValidationError(ve.TypeMin, ve.SubtypeNumber, r.min)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice, reflect.Array:
			if CompareNumbers(valueOf.Len(), r.min) == -1 {
				return value, NewMinValidationError(ve.TypeMin, ve.SubtypeSlice, r.min)
			}

		case reflect.Map:
			if CompareNumbers(valueOf.Len(), r.min) == -1 {
				return value, NewMinValidationError(ve.TypeMin, ve.SubtypeMap, r.min)
			}
		}
	}

	return value, nil
}

func NewMinValidationError[T number](t, st string, threshold T) MinValidationError[T] {
	return MinValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: fmt.Sprintf("%s.%s", t, st),
		},
		Threshold: threshold,
	}
}

type MinValidationError[T number] struct {
	ve.BasicValidationError

	Threshold T `json:"threshold"`
}

func (e MinValidationError[T]) Error() string {
	return fmt.Sprintf("minRule{Threshold=%v}", e.Threshold)
}
