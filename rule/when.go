package rule

import (
	"context"
)

func When(condition bool, rules ...Rule) *whenFuncRule {
	return &whenFuncRule{
		condition: func(_ context.Context, _ any, _ any) bool { return condition },
		rules:     rules,
	}
}
