package validator

import (
	"context"
	"errors"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

var (
	ErrNotStructType = errors.New("not a struct type")
	ErrNotListType   = errors.New("not an array or a slice type")
)

type RulesMap map[string][]vr.Rule

func ForMap(data map[string]any, rules RulesMap) (ve.ErrorsBag, error) {
	return ForMapWithContext(context.Background(), data, rules)
}

func ForMapWithContext(ctx context.Context, data map[string]any, rules RulesMap) (ve.ErrorsBag, error) {
	errorsBag := ve.NewErrorsBag()

	for field, rules := range rules {
		for fieldValue := range newFieldsIterator(field, data) {
			applyRules(ctx, data, rules, fieldValue, errorsBag)
		}
	}

	return errorsBag, nil
}

func ForStruct(data any, rules RulesMap) (ve.ErrorsBag, error) {
	return ForStructWithContext(context.Background(), data, rules)
}

func ForStructWithContext(ctx context.Context, data any, rules RulesMap) (ve.ErrorsBag, error) {
	data, _ = vr.Dereference(data)

	if reflect.TypeOf(data).Kind() != reflect.Struct {
		return nil, ErrNotStructType
	}

	errorsBag := ve.NewErrorsBag()

	for field, rules := range rules {
		for fieldValue := range newFieldsIterator(field, data) {
			applyRules(ctx, data, rules, fieldValue, errorsBag)
		}
	}

	return errorsBag, nil
}

func ForSlice(data any, rules ...vr.Rule) (ve.ErrorsBag, error) {
	return ForSliceWithContext(context.Background(), data, rules...)
}

func ForSliceWithContext(ctx context.Context, data any, rules ...vr.Rule) (ve.ErrorsBag, error) {
	data, _ = vr.Dereference(data)

	if kind := reflect.TypeOf(data).Kind(); kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrNotListType
	}

	errorsBag := ve.NewErrorsBag()

	for fieldValue := range newFieldsIterator("*", data) {
		applyRules(ctx, data, rules, fieldValue, errorsBag)
	}

	return errorsBag, nil
}

func ForValue[In any](value In, rules ...vr.Rule) (errors []ve.ValidationError, _ error) {
	return ForValueWithContext[In](context.Background(), value, rules...)
}

func ForValueWithContext[In any](ctx context.Context, value In, rules ...vr.Rule) (errors []ve.ValidationError, _ error) {
	errorsBag := ve.NewErrorsBag()

	applyRules(
		ctx,
		value,
		rules,
		fieldValue{
			field: "_",
			value: value,
		},
		errorsBag,
	)

	return errorsBag.Get("_"), nil
}

func applyRules(ctx context.Context, data any, rules []vr.Rule, fieldValue fieldValue, errorsBag ve.ErrorsBag) {
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
}
