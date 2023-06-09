package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StringRule(t *testing.T) {
	runRuleTestCases(t, stringRuleDataProvider)
}

func Test_StringValidationError(t *testing.T) {
	// when
	err := NewStringValidationError()

	// then
	require.EqualError(t, err, "must be a string")
}

func BenchmarkStringRule(b *testing.B) {
	runRuleBenchmarks(b, stringRuleDataProvider)
}

func stringRuleDataProvider() map[string]*ruleTestCaseData {
	var stringDummy = fakerInstance.Lorem().Sentence(6)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             String(),
			value:            nil,
			expectedNewValue: (*string)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"string": {
			rule:             String(),
			value:            stringDummy,
			expectedNewValue: stringDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*string": {
			rule:             String(),
			value:            &stringDummy,
			expectedNewValue: &stringDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		"int": {
			rule:             String(),
			value:            0,
			expectedNewValue: nil,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		"float": {
			rule:             String(),
			value:            0.0,
			expectedNewValue: nil,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		"complex": {
			rule:             String(),
			value:            1 + 2i,
			expectedNewValue: nil,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		"bool": {
			rule:             String(),
			value:            true,
			expectedNewValue: nil,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		"slice": {
			rule:             String(),
			value:            []int{},
			expectedNewValue: nil,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		"array": {
			rule:             String(),
			value:            [1]int{},
			expectedNewValue: nil,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		"map": {
			rule:             String(),
			value:            map[any]any{},
			expectedNewValue: nil,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
		"struct": {
			rule:             String(),
			value:            someStruct{},
			expectedNewValue: nil,
			expectedError:    NewStringValidationError(),
			expectedToBail:   true,
		},
	}
}
