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
			return value, NewLengthValidationError(ve.SubtypeString, r.length)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice:
			if CompareNumbers(valueOf.Len(), r.length) != 0 {
				return value, NewLengthValidationError(ve.SubtypeSlice, r.length)
			}

		case reflect.Array:
			if CompareNumbers(valueOf.Len(), r.length) != 0 {
				return value, NewLengthValidationError(ve.SubtypeArray, r.length)
			}

		case reflect.Map:
			if CompareNumbers(valueOf.Len(), r.length) != 0 {
				return value, NewLengthValidationError(ve.SubtypeMap, r.length)
			}
		}
	}

	return value, nil
}

func NewLengthValidationError[T integerType](st string, threshold T) LengthValidationError[T] {
	return LengthValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: fmt.Sprintf("%s.%s", ve.TypeLength, st),
		},
		Threshold: threshold,
	}
}

type LengthValidationError[T integerType] struct {
	ve.BasicValidationError

	Threshold T `json:"threshold"`
}

func (e LengthValidationError[T]) Error() string {
	switch e.Rule {
	case ve.TypeLength + "." + ve.SubtypeSlice,
		ve.TypeLength + "." + ve.SubtypeArray,
		ve.TypeLength + "." + ve.SubtypeMap:
		return fmt.Sprintf("must have exactly %v items", e.Threshold)

	default:
		return fmt.Sprintf("must be exactly %v characters long", e.Threshold)
	}
}
