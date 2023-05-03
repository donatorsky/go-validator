package rule

import (
	"context"

	ve "github.com/donatorsky/go-validator/error"
)

func Bail() *bailRule {
	return &bailRule{}
}

type bailRule struct {
}

func (*bailRule) Apply(_ context.Context, _ any, _ any) (any, ve.ValidationError) {
	return nil, nil
}

func (*bailRule) Bails() bool {
	return true
}
