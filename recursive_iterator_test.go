package validator

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

func Test_recursiveIterator(t *testing.T) {
	// given
	var (
		ruleMock1 = ruleMock{1}
		ruleMock2 = ruleMock{2}
		ruleMock3 = ruleMock{3}
		ruleMock4 = ruleMock{4}
		ruleMock5 = ruleMock{5}
		ruleMock6 = ruleMock{6}
	)

	for ttIdx, tt := range []struct {
		rules         []vr.Rule
		expectedRules []vr.Rule
	}{
		{
			rules:         nil,
			expectedRules: []vr.Rule{},
		},
		{
			rules: []vr.Rule{
				ruleMock1,
				ruleMock2,
				ruleMock3,
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2, ruleMock3},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock(nil),
			},
			expectedRules: []vr.Rule{},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock(nil),
				newRuleWithSubRulesMock(nil),
			},
			expectedRules: []vr.Rule{},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock([]vr.Rule{
					newRuleWithSubRulesMock(nil),
				}),
			},
			expectedRules: []vr.Rule{},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock([]vr.Rule{
					ruleMock1,
					ruleMock2,
					ruleMock3,
				}),
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2, ruleMock3},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock([]vr.Rule{
					ruleMock1,
					ruleMock2,
					ruleMock3,
				}).WithRulesCondition(func(_ context.Context, _ any, _ any) bool {
					return false
				}),
			},
			expectedRules: []vr.Rule{},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock([]vr.Rule{
					newRuleWithSubRulesMock([]vr.Rule{
						ruleMock1,
						ruleMock2,
						ruleMock3,
					}),
				}),
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2, ruleMock3},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock([]vr.Rule{
					newRuleWithSubRulesMock([]vr.Rule{
						ruleMock1,
						ruleMock2,
						ruleMock3,
					}),
				}).WithRulesCondition(func(_ context.Context, _ any, _ any) bool {
					return false
				}),
			},
			expectedRules: []vr.Rule{},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock([]vr.Rule{
					newRuleWithSubRulesMock([]vr.Rule{
						ruleMock1,
						ruleMock2,
						ruleMock3,
					}).WithRulesCondition(func(_ context.Context, _ any, _ any) bool {
						return false
					}),
				}),
			},
			expectedRules: []vr.Rule{},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock([]vr.Rule{
					ruleMock1,
					newRuleWithSubRulesMock([]vr.Rule{ruleMock2}).
						WithRulesCondition(func(_ context.Context, _ any, _ any) bool {
							return false
						}),
					ruleMock3,
				}),
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock3},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock([]vr.Rule{
					ruleMock1,
					newRuleWithSubRulesMock([]vr.Rule{
						ruleMock2,
					}),
					ruleMock3,
				}),
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2, ruleMock3},
		},
		{
			rules: []vr.Rule{
				ruleMock1,
				newRuleWithSubRulesMock([]vr.Rule{
					ruleMock2,
				}),
				ruleMock3,
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2, ruleMock3},
		},
		{
			rules: []vr.Rule{
				ruleMock1,
				ruleMock2,
				newRuleWithSubRulesMock([]vr.Rule{
					ruleMock3,
				}),
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2, ruleMock3},
		},
		{
			rules: []vr.Rule{
				ruleMock1,
				ruleMock2,
				newRuleWithSubRulesMock(nil),
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2},
		},
		{
			rules: []vr.Rule{
				ruleMock1,
				ruleMock2,
				newRuleWithSubRulesMock([]vr.Rule{
					newRuleWithSubRulesMock(nil),
				}),
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2},
		},
		{
			rules: []vr.Rule{
				ruleMock1,
				ruleMock2,
				newRuleWithSubRulesMock([]vr.Rule{
					newRuleWithSubRulesMock([]vr.Rule{
						ruleMock3,
					}),
				}),
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2, ruleMock3},
		},
		{
			rules: []vr.Rule{
				newRuleWithSubRulesMock([]vr.Rule{
					ruleMock1,
					newRuleWithSubRulesMock([]vr.Rule{
						ruleMock2,
					}),
				}),
				ruleMock3,
				newRuleWithSubRulesMock([]vr.Rule{
					newRuleWithSubRulesMock([]vr.Rule{
						ruleMock4,
					}),
					ruleMock5,
				}),
				ruleMock6,
			},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2, ruleMock3, ruleMock4, ruleMock5, ruleMock6},
		},
	} {
		t.Run(fmt.Sprintf("Test data #%d", ttIdx), func(t *testing.T) {
			var rules = make([]vr.Rule, 0, len(tt.expectedRules))

			i := newRecursiveIterator(tt.rules, context.Background(), nil, nil)

			// when
			for i.Valid() {
				rules = append(rules, i.Current())

				i.Next(context.Background(), nil, nil)
			}

			// then
			require.Equal(t, tt.expectedRules, rules)
		})
	}
}

func Test_recursiveIterator_ReturnsNilForEmptyIterator(t *testing.T) {
	// given
	i := newRecursiveIterator(nil, context.Background(), nil, nil)

	// then
	require.Nil(t, i.Current())
	require.False(t, i.Valid())
	require.Nil(t, i.iterator)

	// when
	i.Next(context.Background(), nil, nil)

	// then
	require.Nil(t, i.Current())
	require.False(t, i.Valid())
	require.Nil(t, i.iterator)
}

func Test_stack(t *testing.T) {
	// given
	s := stack{}

	// then
	require.Len(t, s, 0)
	require.True(t, s.Empty())

	// and when
	s.Push(nil)

	// then
	require.Len(t, s, 1)
	require.False(t, s.Empty())

	// and when
	s.Pop()

	// then
	require.Len(t, s, 0)
	require.True(t, s.Empty())
}

func newRuleWithSubRulesMock(rules []vr.Rule) *ruleWithSubRulesMock {
	return &ruleWithSubRulesMock{rules: rules}
}

type ruleWithSubRulesMock struct {
	rules          []vr.Rule
	rulesCondition func(ctx context.Context, value any, data any) bool
}

func (*ruleWithSubRulesMock) Apply(_ context.Context, _ any, _ any) (any, ve.ValidationError) {
	panic("unexpected method call")
}

func (*ruleWithSubRulesMock) Bails() bool {
	panic("unexpected method call")
}

func (r *ruleWithSubRulesMock) Rules(ctx context.Context, value any, data any) []vr.Rule {
	if r.rulesCondition == nil || r.rulesCondition(ctx, value, data) {
		return r.rules
	}

	return nil
}

func (r *ruleWithSubRulesMock) WithRulesCondition(rulesCondition func(ctx context.Context, value any, data any) bool) *ruleWithSubRulesMock {
	r.rulesCondition = rulesCondition

	return r
}

func (r ruleWithSubRulesMock) String() string {
	return fmt.Sprintf("%s", r.rules)
}
