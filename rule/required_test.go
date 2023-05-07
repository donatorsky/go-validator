package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RequiredRule(t *testing.T) {
	// given
	for ttIdx, tt := range requiredRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_RequiredValidationError(t *testing.T) {
	// when
	err := NewRequiredValidationError()

	// then
	require.EqualError(t, err, "is required")
}

func BenchmarkRequiredRule(b *testing.B) {
	for ttIdx, tt := range requiredRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func requiredRuleDataProvider() []*ruleTestCaseData {
	return []*ruleTestCaseData{
		{
			rule:             Required(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    NewRequiredValidationError(),
			expectedToBail:   true,
		},

		{
			rule:             Required(),
			value:            (*string)(nil),
			expectedNewValue: nil,
			expectedError:    NewRequiredValidationError(),
			expectedToBail:   true,
		},

		{
			rule:             Required(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Required(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Required(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Required(),
			value:            true,
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Required(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Required(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Required(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Required(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    nil,
			expectedToBail:   false,
		},
	}
}
