package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StructRule(t *testing.T) {
	// given
	for ttIdx, tt := range structRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_StructValidationError(t *testing.T) {
	// when
	err := NewStructValidationError()

	// then
	require.EqualError(t, err, "must be a struct")
}

func BenchmarkStructRule(b *testing.B) {
	for ttIdx, tt := range structRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func structRuleDataProvider() []*ruleTestCaseData {
	var structDummy = someStruct{}

	return []*ruleTestCaseData{
		{
			rule:             Struct(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   false,
		},

		{
			rule:             Struct(),
			value:            structDummy,
			expectedNewValue: structDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Struct(),
			value:            &structDummy,
			expectedNewValue: &structDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		{
			rule:             Struct(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Struct(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Struct(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Struct(),
			value:            "foo",
			expectedNewValue: "foo",
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Struct(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Struct(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Struct(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Struct(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
	}
}
