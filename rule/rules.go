package rule

import (
	"context"
	"math/big"
	"reflect"

	ve "github.com/donatorsky/go-validator/error"
)

type Rule interface {
	Apply(ctx context.Context, value any, data any) (newValue any, err ve.ValidationError)
}

type WithSubRulesRule interface {
	Rules(ctx context.Context, value any, data any) []Rule
}

type BailingRule interface {
	Bails() bool
}

type Bailer struct {
	bailed bool
}

func (b *Bailer) MarkBailed() {
	b.bailed = true
}

func (b *Bailer) Bails() bool {
	defer func() { b.bailed = false }()

	return b.bailed
}

func Dereference(reference any) (value any, isNil bool) {
	switch valueOf := reflect.ValueOf(reference); valueOf.Kind() {
	case reflect.Invalid:
		return nil, true

	case reflect.Ptr,
		reflect.Interface:
		if valueOf.IsNil() {
			return nil, true
		}

		return Dereference(valueOf.Elem().Interface())

	case reflect.Slice,
		reflect.Map,
		reflect.Func,
		reflect.Chan:
		if valueOf.IsNil() {
			return nil, true
		}
	}

	return reference, false
}

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func CompareNumbers[N1, N2 number](n1 N1, n2 N2) int {
	return toBigFloat(n1).Cmp(toBigFloat(n2))
}

func toBigFloat[T number](n T) *big.Float {
	v := &big.Float{}

	switch valueOf := reflect.ValueOf(n); {
	case valueOf.CanInt():
		v.SetInt64(valueOf.Int())

	case valueOf.CanUint():
		v.SetUint64(valueOf.Uint())

	case valueOf.CanFloat():
		v.SetFloat64(valueOf.Float())
	}

	return v
}
