package validator

import (
	"context"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

type forStructValidatorOption func(options *validatorOptions) error

func ForStruct(data any, rules RulesMap, options ...forStructValidatorOption) (ve.ErrorsBag, error) {
	return ForStructWithContext(context.Background(), data, rules, options...)
}

func ForStructWithContext(ctx context.Context, data any, rules RulesMap, options ...forStructValidatorOption) (ve.ErrorsBag, error) {
	data, _ = vr.Dereference(data)

	if reflect.TypeOf(data).Kind() != reflect.Struct {
		return nil, ve.ErrNotStructType
	}

	opts := &validatorOptions{}

	for _, option := range options {
		if err := option(opts); err != nil {
			return nil, err
		}
	}

	errorsBag := ve.NewErrorsBag()

	for field, rules := range rules {
		for fieldValue := range newFieldsIterator(field, data) {
			if err := applyRules(ctx, data, rules, fieldValue, errorsBag, opts); err != nil {
				return nil, err
			}
		}
	}

	return errorsBag, nil
}

func ForStructWithDataCollector(collector DataCollector) forStructValidatorOption {
	return func(options *validatorOptions) error {
		options.dataCollector = collector

		return nil
	}
}
