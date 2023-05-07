package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StringRule(t *testing.T) {
	// given
	for ttIdx, tt := range stringRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_StringValidationError(t *testing.T) {
	// when
	err := NewStringValidationError()

	// then
	require.EqualError(t, err, "must be a string")
}

func BenchmarkStringRule(b *testing.B) {
	for ttIdx, tt := range stringRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func stringRuleDataProvider() []*ruleTestCaseData {
	var stringDummy = fakerInstance.Lorem().Sentence(6)

	return []*ruleTestCaseData{
		{
			rule:             String(),
			value:            nil,
			expectedNewValue: (*string)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		{
			rule:             String(),
			value:            stringDummy,
			expectedNewValue: stringDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             String(),
			value:            &stringDummy,
			expectedNewValue: &stringDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		{
			rule:             String(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             String(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             String(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             String(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             String(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             String(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             String(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             String(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
	}
}
