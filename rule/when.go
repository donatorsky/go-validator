package rule

import (
	"context"
)

func When(condition bool, rules ...Rule) *whenFuncRule[any] {
	return &whenFuncRule[any]{
		condition: func(_ context.Context, _ any, _ any) bool { return condition },
		rules:     rules,
	}
}
