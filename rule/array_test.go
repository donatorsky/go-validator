package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ArrayRule(t *testing.T) {
	runRuleTestCases(t, arrayRuleDataProvider)
}

func Test_ArrayValidationError(t *testing.T) {
	// when
	err := NewArrayValidationError()

	// then
	require.EqualError(t, err, "must be an array")
}

func BenchmarkArrayRule(b *testing.B) {
	runRuleBenchmarks(b, arrayRuleDataProvider)
}

func arrayRuleDataProvider() map[string]*ruleTestCaseData {
	var arrayDummy = [3]int{1, 2, 3}

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Array(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"pointer to array nil pointer": {
			rule:             Array(),
			value:            (*[3]any)(nil),
			expectedNewValue: (*[3]any)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"array": {
			rule:             Array(),
			value:            arrayDummy,
			expectedNewValue: arrayDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to array": {
			rule:             Array(),
			value:            &arrayDummy,
			expectedNewValue: &arrayDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		"int": {
			rule:             Array(),
			value:            1,
			expectedNewValue: 1,
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		"float": {
			rule:             Array(),
			value:            1.0,
			expectedNewValue: 1.0,
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		"complex": {
			rule:             Array(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		"string": {
			rule:             Array(),
			value:            "foo",
			expectedNewValue: "foo",
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		"bool": {
			rule:             Array(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		"slice": {
			rule:             Array(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		"map": {
			rule:             Array(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
		"struct": {
			rule:             Array(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewArrayValidationError(),
			expectedToBail:   true,
		},
	}
}
