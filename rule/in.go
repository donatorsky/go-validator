package rule

import (
	"context"
	"fmt"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

func In[T any](values []T) *inRule[T] {
	return &inRule[T]{
		values: values,
	}
}

type inRule[T any] struct {
	values []T
}

func (r *inRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	if value == nil {
		return value, nil
	}

	newValue, ok := value.(T)
	if !ok {
		return value, NewInValidationError(r.values)
	}

	for _, v := range r.values {
		if reflect.DeepEqual(v, newValue) {
			return value, nil
		}
	}

	return value, NewInValidationError(r.values)
}

func NewInValidationError[T any](values []T) InValidationError[T] {
	return InValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeIn,
		},
		Values: values,
	}
}

type InValidationError[T any] struct {
	ve.BasicValidationError

	Values []T `json:"values"`
}

func (e InValidationError[T]) Error() string {
	return fmt.Sprintf("inRule{Values=%v}", e.Values)
}
