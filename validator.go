package validator

import (
	"context"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

type RulesMap map[string][]vr.Rule

func ForMap(data map[string]any, rules RulesMap) ve.ErrorsBag {
	return ForMapWithContext(context.Background(), data, rules)
}

func ForMapWithContext(ctx context.Context, data map[string]any, rules RulesMap) ve.ErrorsBag {
	errorsBag := ve.NewErrorsBag()

	for field, rules := range rules {
		for fieldValue := range newFieldsIterator(field, data) {
			applyRules(ctx, data, rules, fieldValue, errorsBag)
		}
	}

	return errorsBag
}

func ForStruct(data any, rules RulesMap) ve.ErrorsBag {
	return ForStructWithContext(context.Background(), data, rules)
}

func ForStructWithContext(ctx context.Context, data any, rules RulesMap) ve.ErrorsBag {
	data, _ = vr.Dereference(data)

	if reflect.TypeOf(data).Kind() != reflect.Struct {
		panic("not a struct")
	}

	errorsBag := ve.NewErrorsBag()

	for field, rules := range rules {
		for fieldValue := range newFieldsIterator(field, data) {
			applyRules(ctx, data, rules, fieldValue, errorsBag)
		}
	}

	return errorsBag
}

func ForSlice(data any, rules ...vr.Rule) ve.ErrorsBag {
	return ForSliceWithContext(context.Background(), data, rules...)
}

func ForSliceWithContext(ctx context.Context, data any, rules ...vr.Rule) ve.ErrorsBag {
	data, _ = vr.Dereference(data)

	if kind := reflect.TypeOf(data).Kind(); kind != reflect.Slice && kind != reflect.Array {
		panic("not a slice or an array")
	}

	errorsBag := ve.NewErrorsBag()

	for fieldValue := range newFieldsIterator("*", data) {
		applyRules(ctx, data, rules, fieldValue, errorsBag)
	}

	return errorsBag
}

func ForValue[In any](value In, rules ...vr.Rule) (errors []ve.ValidationError) {
	return ForValueWithContext[In](context.Background(), value, rules...)
}

func ForValueWithContext[In any](ctx context.Context, value In, rules ...vr.Rule) (errors []ve.ValidationError) {
	errorsBag := ve.NewErrorsBag()

	applyRules(
		ctx,
		value,
		rules,
		fieldValue{
			pattern: false,
			field:   "*",
			value:   value,
		},
		errorsBag,
	)

	return errorsBag.Get("*")
}

func applyRules(ctx context.Context, data any, rules []vr.Rule, fieldValue fieldValue, errorsBag ve.ErrorsBag) {
	anyRuleFailed := false
	i := newRecursiveIterator(rules, ctx, fieldValue.value, data)

	for i.Valid() {
		rule := i.Current()

		var err ve.ValidationError

		value := fieldValue.value

		if value, err = rule.Apply(ctx, value, data); err != nil {
			if compositeError, ok := err.(ve.CompositeValidationError); ok {
				for _, validationError := range compositeError.Errors() {
					errorsBag.Add(fieldValue.field, validationError)
				}
			} else {
				errorsBag.Add(fieldValue.field, err)
			}

			anyRuleFailed = true
		}

		if bailingRule, ok := rule.(vr.BailingRule); ok && anyRuleFailed && bailingRule.Bails() {
			break
		}

		i.Next(ctx, value, data)
	}
}
