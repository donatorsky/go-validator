package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RequiredRule(t *testing.T) {
	runRuleTestCases(t, requiredRuleDataProvider)
}

func Test_RequiredValidationError(t *testing.T) {
	// when
	err := NewRequiredValidationError()

	// then
	require.EqualError(t, err, "is required")
}

func BenchmarkRequiredRule(b *testing.B) {
	runRuleBenchmarks(b, requiredRuleDataProvider)
}

func requiredRuleDataProvider() map[string]*ruleTestCaseData {
	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Required(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    NewRequiredValidationError(),
			expectedToBail:   true,
		},

		"pointer to string nil pointer": {
			rule:             Required(),
			value:            (*string)(nil),
			expectedNewValue: nil,
			expectedError:    NewRequiredValidationError(),
			expectedToBail:   true,
		},

		"int": {
			rule:             Required(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"float": {
			rule:             Required(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"complex": {
			rule:             Required(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"bool": {
			rule:             Required(),
			value:            true,
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"slice": {
			rule:             Required(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    nil,
			expectedToBail:   false,
		},
		"array": {
			rule:             Required(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    nil,
			expectedToBail:   false,
		},
		"map": {
			rule:             Required(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    nil,
			expectedToBail:   false,
		},
		"struct": {
			rule:             Required(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    nil,
			expectedToBail:   false,
		},
	}
}
