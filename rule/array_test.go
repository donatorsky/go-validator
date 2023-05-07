package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ArrayRule(t *testing.T) {
	// given
	for ttIdx, tt := range arrayRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_ArrayValidationError(t *testing.T) {
	// when
	err := NewArrayValidationError()

	// then
	require.EqualError(t, err, "must be an array")
}

func BenchmarkArrayRule(b *testing.B) {
	for ttIdx, tt := range arrayRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func arrayRuleDataProvider() []*ruleTestCaseData {
	var arrayDummy = [3]int{1, 2, 3}

	return []*ruleTestCaseData{
		{
			rule:             Array(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   false,
		},

		{
			rule:             Array(),
			value:            (*[3]any)(nil),
			expectedNewValue: (*[3]any)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		{
			rule:             Array(),
			value:            arrayDummy,
			expectedNewValue: arrayDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Array(),
			value:            &arrayDummy,
			expectedNewValue: &arrayDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		{
			rule:             Array(),
			value:            1,
			expectedNewValue: 1,
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Array(),
			value:            1.0,
			expectedNewValue: 1.0,
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Array(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Array(),
			value:            "foo",
			expectedNewValue: "foo",
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Array(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Array(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Array(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Array(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
	}
}
