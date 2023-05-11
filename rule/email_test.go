package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmailRule(t *testing.T) {
	runRuleTestCases(t, emailRuleDataProvider)
}

func Test_EmailValidationError(t *testing.T) {
	// when
	err := NewEmailValidationError()

	// then
	require.EqualError(t, err, "must be a valid email address")
}

func BenchmarkEmailRule(b *testing.B) {
	runRuleBenchmarks(b, emailRuleDataProvider)
}

func emailRuleDataProvider() map[string]*ruleTestCaseData {
	var emailDummy = fakerInstance.Internet().Email()

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Email(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"email address": {
			rule:             Email(),
			value:            emailDummy,
			expectedNewValue: emailDummy,
			expectedError:    nil,
		},
		"pointer to email address": {
			rule:             Email(),
			value:            &emailDummy,
			expectedNewValue: &emailDummy,
			expectedError:    nil,
		},
		"invalid email address": {
			rule:             Email(),
			value:            "invalid email address",
			expectedNewValue: "invalid email address",
			expectedError:    NewEmailValidationError(),
		},

		// unsupported values
		"int": {
			rule:             Email(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewEmailValidationError(),
		},
		"float": {
			rule:             Email(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewEmailValidationError(),
		},
		"complex": {
			rule:             Email(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewEmailValidationError(),
		},
		"bool": {
			rule:             Email(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewEmailValidationError(),
		},
		"slice": {
			rule:             Email(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewEmailValidationError(),
		},
		"array": {
			rule:             Email(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewEmailValidationError(),
		},
		"map": {
			rule:             Email(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewEmailValidationError(),
		},
		"struct": {
			rule:             Email(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewEmailValidationError(),
		},
	}
}
