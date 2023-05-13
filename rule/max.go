package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Max[T numberType](max T) *maxRule[T] {
	return &maxRule[T]{
		max: max,
	}
}

type maxRule[T numberType] struct {
	max T
}

func (r *maxRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	switch v := v.(type) {
	case string:
		if CompareNumbers(len(v), r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeString, r.max)
		}

	case int:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case int8:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case int16:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case int32:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case int64:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case uint:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case uint8:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case uint16:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case uint32:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case uint64:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case float32:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	case float64:
		if CompareNumbers(v, r.max) == 1 {
			return value, NewMaxValidationError(ve.SubtypeNumber, r.max)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice:
			if CompareNumbers(valueOf.Len(), r.max) == 1 {
				return value, NewMaxValidationError(ve.SubtypeSlice, r.max)
			}

		case reflect.Array:
			if CompareNumbers(valueOf.Len(), r.max) == 1 {
				return value, NewMaxValidationError(ve.SubtypeArray, r.max)
			}

		case reflect.Map:
			if CompareNumbers(valueOf.Len(), r.max) == 1 {
				return value, NewMaxValidationError(ve.SubtypeMap, r.max)
			}
		}
	}

	return value, nil
}

func NewMaxValidationError[T numberType](st string, threshold T) MaxValidationError[T] {
	return MaxValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: fmt.Sprintf("%s.%s", ve.TypeMax, st),
		},
		Threshold: threshold,
	}
}

type MaxValidationError[T numberType] struct {
	ve.BasicValidationError

	Threshold T `json:"threshold"`
}

func (e MaxValidationError[T]) Error() string {
	switch e.Rule {
	case ve.TypeMax + "." + ve.SubtypeString:
		return fmt.Sprintf("must be at most %v characters", e.Threshold)

	case ve.TypeMax + "." + ve.SubtypeSlice,
		ve.TypeMax + "." + ve.SubtypeArray,
		ve.TypeMax + "." + ve.SubtypeMap:
		return fmt.Sprintf("must have at most %v items", e.Threshold)

	default:
		return fmt.Sprintf("must be at most %v", e.Threshold)
	}
}
