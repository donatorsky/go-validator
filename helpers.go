package validator

import (
	"context"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

type RulesMap map[string][]vr.Rule

func applyRules(ctx context.Context, data any, rules []vr.Rule, fieldValue fieldValue, errorsBag ve.ErrorsBag, options *validatorOptions) error {
	anyRuleFailed := false
	i := newRecursiveIterator(rules, ctx, fieldValue.value, data)

	value := fieldValue.value

	for i.Valid() {
		rule := i.Current()

		var err ve.ValidationError

		if value, err = rule.Apply(ctx, value, data); err != nil {
			errorsBag.Add(fieldValue.field, err)

			anyRuleFailed = true
		}

		if bailingRule, ok := rule.(vr.BailingRule); ok && anyRuleFailed && bailingRule.Bails() {
			break
		}

		i.Next(ctx, value, data)
	}

	if !anyRuleFailed && options.dataCollector != nil {
		options.dataCollector.Set(fieldValue.field, value)
	}

	if !anyRuleFailed && options.valueExporter != nil {
		targetValue := options.valueExporter.Elem()
		targetType := targetValue.Type()

		if valueType := reflect.TypeOf(value); !valueType.ConvertibleTo(targetType) {
			return ve.ValueExporterTypeMismatchError{
				ValueType:  valueType.String(),
				TargetType: targetType.String(),
			}
		}

		targetValue.Set(reflect.ValueOf(value).Convert(targetType))
	}

	return nil
}
