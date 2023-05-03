package rule

import (
	"context"

	ve "github.com/donatorsky/go-validator/error"
)

type whenFuncCondition[T any] func(ctx context.Context, value T, data any) bool

func WhenFunc[T any](condition whenFuncCondition[T], rules ...Rule) *whenFuncRule[T] {
	return &whenFuncRule[T]{
		condition: condition,
		rules:     rules,
	}
}

type whenFuncRule[T any] struct {
	Bailer

	condition whenFuncCondition[T]
	rules     []Rule
}

func (r *whenFuncRule[T]) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	return nil, nil
}

func (r *whenFuncRule[T]) Rules(ctx context.Context, value any, data any) []Rule {
	value, isNil := Dereference(value)
	if isNil {
		return nil
	}

	newValue, ok := value.(T)
	if !ok {
		return nil
	}

	if !r.condition(ctx, newValue, data) {
		return nil
	}

	return r.rules
}

//func (r *whenFuncRule[T]) Apply(ctx context.Context, value any, data any) (any, ve.ValidationError) {
//	newValue, ok := value.(T)
//	if !ok {
//		return value, nil
//	}
//
//	if !r.condition(ctx, newValue, data) {
//		return value, nil
//	}
//
//	anyRuleFailed := false
//	errors := ve.CompositeValidationError{}
//
//	for _, rule := range r.rules {
//		var err ve.ValidationError
//		if value, err = rule.Apply(ctx, value, data); err != nil {
//			if compositeError, ok := err.(ve.CompositeValidationError); ok {
//				for _, validationError := range compositeError.Errors() {
//					errors.Add(validationError)
//				}
//			} else {
//				errors.Add(err)
//			}
//
//			anyRuleFailed = true
//		}
//
//		if bailingRule, ok := rule.(BailingRule); ok && anyRuleFailed && bailingRule.Bails() {
//			r.bailed = true
//
//			break
//		}
//	}
//
//	if !errors.Empty() {
//		return value, errors
//	}
//
//	return newValue, nil
//}
//
//func (r *whenFuncRule[T]) Rules(ctx context.Context, value any, data any) []Rule {
//	return r.rules
//}
