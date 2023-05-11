package rule

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_WhenRule(t *testing.T) {
	runRuleTestCases(t, whenRuleDataProvider)
}

func BenchmarkWhenRule(b *testing.B) {
	runRuleBenchmarks(b, whenRuleDataProvider)
}

func TestWhen_Rules(t *testing.T) {
	// given
	var stringDummy = fakerInstance.Lorem().Sentence(6)

	for ttName, tt := range map[string]struct {
		rule          *whenFuncRule
		value         any
		expectedRules []Rule
	}{
		"nil value, condition is false": {
			rule:          When(false, newRuleMock(0), newRuleMock(1)),
			value:         nil,
			expectedRules: nil,
		},
		"nil value, condition is true": {
			rule:  When(true, newRuleMock(0), newRuleMock(1)),
			value: nil,
			expectedRules: []Rule{
				newRuleMock(0),
				newRuleMock(1),
			},
		},

		"string value, condition is false": {
			rule:          When(false, newRuleMock(0), newRuleMock(1)),
			value:         stringDummy,
			expectedRules: nil,
		},
		"string value, condition is true": {
			rule:  When(true, newRuleMock(0), newRuleMock(1)),
			value: stringDummy,
			expectedRules: []Rule{
				newRuleMock(0),
				newRuleMock(1),
			},
		},

		"*string value, condition is false": {
			rule:          When(false, newRuleMock(0), newRuleMock(1)),
			value:         &stringDummy,
			expectedRules: nil,
		},
		"*string value, condition is true": {
			rule:  When(true, newRuleMock(0), newRuleMock(1)),
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

func whenRuleDataProvider() map[string]*ruleTestCaseData {
	var stringDummy = fakerInstance.Lorem().Sentence(6)

	return map[string]*ruleTestCaseData{
		"nil, condition fails": {
			rule:             When(false, nil),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"nil, condition succeeds": {
			rule:             When(true, nil),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string, condition fails": {
			rule:             When(false, nil),
			value:            stringDummy,
			expectedNewValue: nil,
			expectedError:    nil,
		},
		"string, condition succeeds": {
			rule:             When(true, nil),
			value:            stringDummy,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string pointer, condition fails": {
			rule:             When(false, nil),
			value:            &stringDummy,
			expectedNewValue: nil,
			expectedError:    nil,
		},
		"string pointer, condition succeeds": {
			rule:             When(true, nil),
			value:            &stringDummy,
			expectedNewValue: nil,
			expectedError:    nil,
		},
	}
}
