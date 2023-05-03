package validator

import (
	"context"

	vr "github.com/donatorsky/go-validator/rule"
)

func newRulesIterator(rules []vr.Rule) *rulesIterator {
	return &rulesIterator{
		rules: rules,
		index: 0,
	}
}

type rulesIterator struct {
	rules []vr.Rule
	index int
}

func (i *rulesIterator) Current() vr.Rule {
	if !i.Valid() {
		return nil
	}

	return i.rules[i.index]
}

func (i *rulesIterator) Next(_ context.Context, _ any, _ any) {
	if !i.Valid() {
		return
	}

	i.index++
}

func (i *rulesIterator) Valid() bool {
	return i.index < len(i.rules)
}
