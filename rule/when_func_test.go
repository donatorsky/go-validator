package rule

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_WhenFuncRule(t *testing.T) {
	runRuleTestCases(t, whenFuncRuleDataProvider)
}

func BenchmarkWhenFuncRule(b *testing.B) {
	runRuleBenchmarks(b, whenFuncRuleDataProvider)
}

func TestWhenFunc_Rules(t *testing.T) {
	// given
	var stringDummy = fakerInstance.Lorem().Sentence(6)

	for ttName, tt := range map[string]struct {
		rule          *whenFuncRule
		value         any
		expectedRules []Rule
	}{
		"nil value, condition is false": {
			rule: WhenFunc(func(_ context.Context, _ any, _ any) bool {
				return false
			}, newRuleMock(0), newRuleMock(1)),
			value:         nil,
			expectedRules: nil,
		},
		"nil value, condition is true": {
			rule: WhenFunc(func(_ context.Context, _ any, _ any) bool {
				return true
			}, newRuleMock(0), newRuleMock(1)),
			value: nil,
			expectedRules: []Rule{
				newRuleMock(0),
				newRuleMock(1),
			},
		},

		"string value, condition is false": {
			rule: WhenFunc(func(_ context.Context, value any, _ any) bool {
				return value != stringDummy
			}, newRuleMock(0), newRuleMock(1)),
			value:         stringDummy,
			expectedRules: nil,
		},
		"string value, condition is true": {
			rule: WhenFunc(func(_ context.Context, value any, _ any) bool {
				return value == stringDummy
			}, newRuleMock(0), newRuleMock(1)),
			value: stringDummy,
			expectedRules: []Rule{
				newRuleMock(0),
				newRuleMock(1),
			},
		},

		"*string value, condition is false": {
			rule: WhenFunc(func(_ context.Context, value any, _ any) bool {
				if value, ok := value.(*string); ok {
					return *value != stringDummy
				}

				return true
			}, newRuleMock(0), newRuleMock(1)),
			value:         &stringDummy,
			expectedRules: nil,
		},
		"*string value, condition is true": {
			rule: WhenFunc(func(_ context.Context, value any, _ any) bool {
				if value, ok := value.(*string); ok {
					return *value == stringDummy
				}

				return false
			}, newRuleMock(0), newRuleMock(1)),
			value: &stringDummy,
			expectedRules: []Rule{
				newRuleMock(0),
				newRuleMock(1),
			},
		},
	} {
		t.Run(ttName, func(t *testing.T) {
			// when
			rules := tt.rule.Rules(context.Background(), tt.value, nil)

			// then
			require.Equal(t, tt.expectedRules, rules)
		})
	}
}

func whenFuncRuleDataProvider() map[string]*ruleTestCaseData {
	var stringDummy = fakerInstance.Lorem().Sentence(6)

	return map[string]*ruleTestCaseData{
		"nil, condition fails": {
			rule:             WhenFunc(func(_ context.Context, _ any, _ any) bool { return false }, nil),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"nil, condition succeeds": {
			rule:             WhenFunc(func(_ context.Context, _ any, _ any) bool { return true }, nil),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string, condition fails": {
			rule:             WhenFunc(func(_ context.Context, _ any, _ any) bool { return false }, nil),
			value:            stringDummy,
			expectedNewValue: nil,
			expectedError:    nil,
		},
		"string, condition succeeds": {
			rule:             WhenFunc(func(_ context.Context, _ any, _ any) bool { return true }, nil),
			value:            stringDummy,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string pointer, condition fails": {
			rule:             WhenFunc(func(_ context.Context, _ any, _ any) bool { return false }, nil),
			value:            &stringDummy,
			expectedNewValue: nil,
			expectedError:    nil,
		},
		"string pointer, condition succeeds": {
			rule:             WhenFunc(func(_ context.Context, _ any, _ any) bool { return true }, nil),
			value:            &stringDummy,
			expectedNewValue: nil,
			expectedError:    nil,
		},
	}
}
