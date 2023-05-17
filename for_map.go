package validator

import (
	"context"

	ve "github.com/donatorsky/go-validator/error"
)

type forMapValidatorOption func(options *validatorOptions) error

func ForMap(data map[string]any, rules RulesMap, options ...forMapValidatorOption) (ve.ErrorsBag, error) {
	return ForMapWithContext(context.Background(), data, rules, options...)
}

func ForMapWithContext(ctx context.Context, data map[string]any, rules RulesMap, options ...forMapValidatorOption) (ve.ErrorsBag, error) {
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

func ForMapWithDataCollector(collector DataCollector) forMapValidatorOption {
	return func(options *validatorOptions) error {
		options.dataCollector = collector

		return nil
	}
}
