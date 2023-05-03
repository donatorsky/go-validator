package validator

import (
	"context"

	vr "github.com/donatorsky/go-validator/rule"
)

type iterator interface {
	Current() vr.Rule
	Next(ctx context.Context, value any, data any)
	Valid() bool
}
