package validator

import (
	"context"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

type forSliceValidatorOption func(options *validatorOptions) error

func ForSlice(data any, rules []vr.Rule, options ...forSliceValidatorOption) (ve.ErrorsBag, error) {
	return ForSliceWithContext(context.Background(), data, rules, options...)
}

func ForSliceWithContext(ctx context.Context, data any, rules []vr.Rule, options ...forSliceValidatorOption) (ve.ErrorsBag, error) {
	data, _ = vr.Dereference(data)

	if kind := reflect.TypeOf(data).Kind(); kind != reflect.Slice && kind != reflect.Array {
		return nil, ve.NotListTypeError{}
	}

	opts := &validatorOptions{}

	for _, option := range options {
		if err := option(opts); err != nil {
			return nil, err
		}
	}

	errorsBag := ve.NewErrorsBag()

	for fieldValue := range newFieldsIterator("*", data) {
		if err := applyRules(ctx, data, rules, fieldValue, errorsBag, opts); err != nil {
			return nil, err
		}
	}

	return errorsBag, nil
}

func ForSliceWithDataCollector(collector DataCollector) forSliceValidatorOption {
	return func(options *validatorOptions) error {
		options.dataCollector = collector

		return nil
	}
}
