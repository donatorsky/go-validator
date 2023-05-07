package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmailRule(t *testing.T) {
	// given
	for ttIdx, tt := range emailRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_EmailValidationError(t *testing.T) {
	// when
	err := NewEmailValidationError()

	// then
	require.EqualError(t, err, "must be a valid email address")
}

func BenchmarkEmailRule(b *testing.B) {
	for ttIdx, tt := range emailRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func emailRuleDataProvider() []*ruleTestCaseData {
	var emailDummy = fakerInstance.Internet().Email()

	return []*ruleTestCaseData{
		{
			rule:             Email(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		{
			rule:             Email(),
			value:            emailDummy,
			expectedNewValue: emailDummy,
			expectedError:    nil,
		},
		{
			rule:             Email(),
			value:            &emailDummy,
			expectedNewValue: &emailDummy,
			expectedError:    nil,
		},
		{
			rule:             Email(),
			value:            "invalid email address",
			expectedNewValue: "invalid email address",
			expectedError:    NewEmailValidationError(),
		},

		// unsupported values
		{
			rule:             Email(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             Email(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             Email(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             Email(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             Email(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             Email(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             Email(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             Email(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewEmailValidationError(),
		},
	}
}
