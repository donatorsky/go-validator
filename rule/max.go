package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Max[T number](max T) *maxRule[T] {
	return &maxRule[T]{
		max: max,
	}
}

type maxRule[T number] struct {
	max T
}

func (r maxRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	switch v := v.(type) {
	case string:
		if CompareNumbers(len(v), r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeString, r.max)
		}

	case int:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case int8:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case int16:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case int32:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case int64:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case uint:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case uint8:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case uint16:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case uint32:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case uint64:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case float32:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	case float64:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeNumber, r.max)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice, reflect.Array:
			if CompareNumbers(valueOf.Len(), r.max) == 1 {
				return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeSlice, r.max)
			}

		case reflect.Map:
			if CompareNumbers(valueOf.Len(), r.max) == 1 {
				return value, NewMaxValidationError(ve.TypeMax, ve.SubtypeMap, r.max)
			}
		}
	}

	return value, nil
}

func NewMaxValidationError[T number](t, st string, threshold T) MaxValidationError[T] {
	return MaxValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: fmt.Sprintf("%s.%s", t, st),
		},
		Threshold: threshold,
	}
}

type MaxValidationError[T number] struct {
	ve.BasicValidationError

	Threshold T `json:"threshold"`
}

func (e MaxValidationError[T]) Error() string {
	return fmt.Sprintf("maxRule{Threshold=%v}", e.Threshold)
}
