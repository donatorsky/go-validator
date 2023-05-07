package rule

import (
	"context"

	ve "github.com/donatorsky/go-validator/error"
)

type whenFuncCondition func(ctx context.Context, value any, data any) bool

func WhenFunc(condition whenFuncCondition, rules ...Rule) *whenFuncRule {
	return &whenFuncRule{
		condition: condition,
		rules:     rules,
	}
}

type whenFuncRule struct {
	Bailer

	condition whenFuncCondition
	rules     []Rule
}

func (r *whenFuncRule) Apply(_ context.Context, _ any, _ any) (any, ve.ValidationError) {
	return nil, nil
}

func (r *whenFuncRule) Rules(ctx context.Context, value any, data any) []Rule {
	if !r.condition(ctx, value, data) {
		return nil
	}

	return r.rules
}
