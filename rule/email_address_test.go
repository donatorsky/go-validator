package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmailAddressRule(t *testing.T) {
	// given
	for ttIdx, tt := range emailAddressRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_EmailAddressValidationError(t *testing.T) {
	// when
	err := NewEmailValidationError()

	// then
	require.EqualError(t, err, "must be a valid email address")
}

func BenchmarkEmailAddressRule(b *testing.B) {
	for ttIdx, tt := range emailAddressRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func emailAddressRuleDataProvider() []*ruleTestCaseData {
	var (
		email1Dummy = fakerInstance.Internet().Email()
		email2Dummy = fmt.Sprintf(
			"%s <%s> (%s)",
			fakerInstance.Person().FirstName(),
			email1Dummy,
			fakerInstance.Lorem().Sentence(3),
		)
	)

	return []*ruleTestCaseData{
		{
			rule:             EmailAddress(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		{
			rule:             EmailAddress(),
			value:            email1Dummy,
			expectedNewValue: email1Dummy,
			expectedError:    nil,
		},
		{
			rule:             EmailAddress(),
			value:            &email1Dummy,
			expectedNewValue: email1Dummy,
			expectedError:    nil,
		},
		{
			rule:             EmailAddress(),
			value:            email2Dummy,
			expectedNewValue: email1Dummy,
			expectedError:    nil,
		},
		{
			rule:             EmailAddress(),
			value:            &email2Dummy,
			expectedNewValue: email1Dummy,
			expectedError:    nil,
		},
		{
			rule:             EmailAddress(),
			value:            "invalid email address",
			expectedNewValue: "invalid email address",
			expectedError:    NewEmailValidationError(),
		},

		// unsupported values
		{
			rule:             EmailAddress(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             EmailAddress(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             EmailAddress(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             EmailAddress(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             EmailAddress(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             EmailAddress(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             EmailAddress(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewEmailValidationError(),
		},
		{
			rule:             EmailAddress(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewEmailValidationError(),
		},
	}
}
