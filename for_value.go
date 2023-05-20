package validator

import (
	"context"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

type forValueValidatorOption func(options *validatorOptions) error

func ForValue[In any](value In, rules []vr.Rule, options ...forValueValidatorOption) (errors []ve.ValidationError, _ error) {
	return ForValueWithContext[In](context.Background(), value, rules, options...)
}

func ForValueWithContext[In any](ctx context.Context, value In, rules []vr.Rule, options ...forValueValidatorOption) (errors []ve.ValidationError, _ error) {
	opts := &validatorOptions{}

	for _, option := range options {
		if err := option(opts); err != nil {
			return nil, err
		}
	}

	errorsBag := ve.NewErrorsBag()

	if err := applyRules(
		ctx,
		value,
		rules,
		fieldValue{
			field: "_",
			value: value,
		},
		errorsBag,
		opts,
	); err != nil {
		return nil, err
	}

	return errorsBag.Get("_"), nil
}

func ForValueWithValueExporter[Out any](value *Out) forValueValidatorOption {
	return func(options *validatorOptions) error {
		rv := reflect.ValueOf(value)
		options.valueExporter = &rv

		return nil
	}
}
