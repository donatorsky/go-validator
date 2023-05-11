package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StructRule(t *testing.T) {
	runRuleTestCases(t, structRuleDataProvider)
}

func Test_StructValidationError(t *testing.T) {
	// when
	err := NewStructValidationError()

	// then
	require.EqualError(t, err, "must be a struct")
}

func BenchmarkStructRule(b *testing.B) {
	runRuleBenchmarks(b, structRuleDataProvider)
}

func structRuleDataProvider() map[string]*ruleTestCaseData {
	var structDummy = someStruct{}

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Struct(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"struct": {
			rule:             Struct(),
			value:            structDummy,
			expectedNewValue: structDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to struct": {
			rule:             Struct(),
			value:            &structDummy,
			expectedNewValue: &structDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		"int": {
			rule:             Struct(),
			value:            0,
			expectedNewValue: nil,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		"float": {
			rule:             Struct(),
			value:            0.0,
			expectedNewValue: nil,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		"complex": {
			rule:             Struct(),
			value:            1 + 2i,
			expectedNewValue: nil,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		"string": {
			rule:             Struct(),
			value:            "foo",
			expectedNewValue: nil,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		"bool": {
			rule:             Struct(),
			value:            true,
			expectedNewValue: nil,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		"slice": {
			rule:             Struct(),
			value:            []int{},
			expectedNewValue: nil,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		"array": {
			rule:             Struct(),
			value:            [1]int{},
			expectedNewValue: nil,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
		"map": {
			rule:             Struct(),
			value:            map[any]any{},
			expectedNewValue: nil,
			expectedError:    NewStructValidationError(),
			expectedToBail:   true,
		},
	}
}
