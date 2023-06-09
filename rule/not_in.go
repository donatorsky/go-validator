package rule

import (
	"context"
	"fmt"

	ve "github.com/donatorsky/go-validator/error"
)

type notInRuleOption func(rule *notInRuleOptions)

func NotIn[T comparable](values []T, options ...notInRuleOption) *notInRule[T] {
	opts := notInRuleOptions{
		comparator:      nil,
		autoDereference: true,
	}

	for _, option := range options {
		option(&opts)
	}

	r := notInRule[T]{
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

type notInRule[T comparable] struct {
	values             []T
	valuesMap          map[T]any
	customErrorMessage *string
	options            notInRuleOptions
}

func (r *notInRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
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
			return value, nil
		}

		_, exists := r.valuesMap[comparableValue]
		if exists {
			return value, NewNotInValidationError(r.values)
		}
	} else {
		for _, v := range r.values {
			if r.options.comparator(newValue, v) {
				return value, NewNotInValidationError(r.values)
			}
		}
	}

	return value, nil
}

func NewNotInValidationError[T any](values []T) NotInValidationError[T] {
	return NotInValidationError[T]{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleNotIn,
		},
		Values: values,
	}
}

type NotInValidationError[T any] struct {
	ve.BasicValidationError

	Values []T `json:"values"`
}

func (e NotInValidationError[T]) Error() string {
	return fmt.Sprintf("exists in %v", e.Values)
}

type notInRuleOptions struct {
	comparator      Comparator
	autoDereference bool
}

func NotInRuleWithComparator(comparator Comparator) notInRuleOption {
	return func(rule *notInRuleOptions) {
		rule.comparator = comparator
	}
}

func NotInRuleWithoutAutoDereference() notInRuleOption {
	return func(rule *notInRuleOptions) {
		rule.autoDereference = false
	}
}
