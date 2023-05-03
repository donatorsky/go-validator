package validator

import (
	"context"

	vr "github.com/donatorsky/go-validator/rule"
)

func newRecursiveIterator(rules []vr.Rule, ctx context.Context, value any, data any) *recursiveIterator {
	ri := &recursiveIterator{iterator: newRulesIterator(rules)}

	if ri.iterate(ctx, value, data); !ri.Valid() {
		return &recursiveIterator{iterator: nil}
	}

	return ri
}

type recursiveIterator struct {
	iterator iterator
	parent   stack
}

func (i *recursiveIterator) Current() vr.Rule {
	if !i.Valid() {
		return nil
	}

	return i.iterator.Current()
}

func (i *recursiveIterator) Next(ctx context.Context, value any, data any) {
	if !i.Valid() {
		return
	}

	i.iterator.Next(ctx, value, data)

	i.iterate(ctx, value, data)
}

func (i *recursiveIterator) Valid() bool {
	return i.iterator != nil
}

func (i *recursiveIterator) iterate(ctx context.Context, value any, data any) {
	for {
		if !i.iterator.Valid() {
			if i.parent.Empty() {
				i.iterator = nil

				break
			}

			i.iterator = i.parent.Pop()

			continue
		}

		rule := i.iterator.Current()

		if withSubRules, ok := rule.(vr.WithSubRulesRule); ok {
			i.parent.Push(i.iterator)
			i.iterator.Next(ctx, value, data)

			i.iterator = newRulesIterator(withSubRules.Rules(ctx, value, data))

			continue
		}

		break
	}
}

type stack []iterator

func (s *stack) Push(v iterator) {
	*s = append(*s, v)
}

func (s *stack) Pop() iterator {
	l := len(*s)
	i := (*s)[l-1]
	*s = (*s)[:l-1]

	return i
}

func (s stack) Empty() bool {
	return len(s) == 0
}
