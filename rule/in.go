package rule

import (
	"context"
	"fmt"

	ve "github.com/donatorsky/go-validator/error"
)

type inRuleOption func(options *inRuleOptions)

func In[T comparable](values []T, options ...inRuleOption) *inRule[T] {
	opts := inRuleOptions{
		comparator:      nil,
		autoDereference: true,
	}

	for _, option := range options {
		option(&opts)
	}

	r := inRule[T]{
		values:    values,
		valuesMap: nil,
		options:   opts,
	}

	if opts.comparator == nil {
		r.valuesMap = map[T]any{}

		for _, value := range values {
			r.valuesMap[value] = nil
		}
	}

	return &r
}

type inRule[T comparable] struct {
	values             []T
	valuesMap          map[T]any
	customErrorMessage *string
	options            inRuleOptions
}

func (r *inRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	var newValue any
	if r.options.autoDereference {
		var isNil bool
		newValue, isNil = Dereference(value)

		if isNil {
			return value, nil
		}
	} else {
		newValue = value
	}

	if r.options.comparator == nil {
		comparableValue, ok := newValue.(T)
		if !ok {
			return value, NewInValidationError(r.values)
		}

		_, exists := r.valuesMap[comparableValue]
		if exists {
			return value, nil
		}
	} else {
		for _, v := range r.values {
			if r.options.comparator(newValue, v) {
				return value, nil
			}
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
	return fmt.Sprintf("does not exist in %v", e.Values)
}

type inRuleOptions struct {
	comparator      Comparator
	autoDereference bool
}

func InRuleWithComparator(comparator Comparator) inRuleOption {
	return func(options *inRuleOptions) {
		options.comparator = comparator
	}
}

func InRuleWithoutAutoDereference() inRuleOption {
	return func(options *inRuleOptions) {
		options.autoDereference = false
	}
}
