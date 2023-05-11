package rule

import (
	"context"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func Map() *mapRule {
	return &mapRule{}
}

type mapRule struct {
	Bailer
}

func (r *mapRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	if reflect.TypeOf(v).Kind() != reflect.Map {
		r.MarkBailed()

		return nil, NewMapValidationError()
	}

	return value, nil
}

func NewMapValidationError() MapValidationError {
	return MapValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeMap,
		},
	}
}

type MapValidationError struct {
	ve.BasicValidationError
}

func (MapValidationError) Error() string {
	return "must be a map"
}
