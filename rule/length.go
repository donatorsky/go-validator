package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Length[T integerType](length T) *lengthRule[T] {
	return &lengthRule[T]{
		length: length,
	}
}

type lengthRule[T integerType] struct {
	length T
}

func (r *lengthRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	switch v := v.(type) {
	case string:
		if CompareNumbers(len(v), r.length) != 0 {
			return value, NewLengthValidationError(ve.TypeString, r.length)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice:
			if CompareNumbers(valueOf.Len(), r.length) != 0 {
				return value, NewLengthValidationError(ve.TypeSlice, r.length)
			}

		case reflect.Array:
			if CompareNumbers(valueOf.Len(), r.length) != 0 {
				return value, NewLengthValidationError(ve.TypeArray, r.length)
			}

		case reflect.Map:
			if CompareNumbers(valueOf.Len(), r.length) != 0 {
				return value, NewLengthValidationError(ve.TypeMap, r.length)
			}

		default:
			return value, NewLengthValidationError(ve.TypeInvalid, r.length)
		}
	}

	return value, nil
}

func NewLengthValidationError[T integerType](st string, threshold T) LengthValidationError[T] {
	return LengthValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleLength,
		},
		Type:   st,
		Length: threshold,
	}
}

type LengthValidationError[T integerType] struct {
	ve.BasicValidationError

	Type   string `json:"type"`
	Length T      `json:"length"`
}

func (e LengthValidationError[T]) Error() string {
	switch e.Type {
	case ve.TypeString:
		return fmt.Sprintf("must be exactly %v characters long", e.Length)

	case ve.TypeSlice,
		ve.TypeArray,
		ve.TypeMap:
		return fmt.Sprintf("must have exactly %v items", e.Length)

	default:
		return "length cannot be determined"
	}
}
