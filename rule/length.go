package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Length[T number](length T) *lengthRule[T] {
	return &lengthRule[T]{
		length: length,
	}
}

type lengthRule[T number] struct {
	length T
}

func (r lengthRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	switch v := v.(type) {
	case string:
		if CompareNumbers(len(v), r.length) == 0 {
			return value, NewLengthValidationError(ve.TypeLength, ve.SubtypeString, r.length)
		}

	default:
		switch valueOf := reflect.ValueOf(v); valueOf.Kind() {
		case reflect.Slice, reflect.Array:
			if CompareNumbers(valueOf.Len(), r.length) == 0 {
				return value, NewLengthValidationError(ve.TypeLength, ve.SubtypeSlice, r.length)
			}

		case reflect.Map:
			if CompareNumbers(valueOf.Len(), r.length) == 0 {
				return value, NewLengthValidationError(ve.TypeLength, ve.SubtypeMap, r.length)
			}
		}
	}

	return value, nil
}

func NewLengthValidationError[T number](t, st string, threshold T) LengthValidationError[T] {
	return LengthValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: fmt.Sprintf("%s.%s", t, st),
		},
		Threshold: threshold,
	}
}

type LengthValidationError[T number] struct {
	ve.BasicValidationError

	Threshold T `json:"threshold"`
}

func (e LengthValidationError[T]) Error() string {
	return fmt.Sprintf("lengthRule{Threshold=%v}", e.Threshold)
}
