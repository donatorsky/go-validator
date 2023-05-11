package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmailAddressRule(t *testing.T) {
	runRuleTestCases(t, emailAddressRuleDataProvider)
}

func Test_EmailAddressValidationError(t *testing.T) {
	// when
	err := NewEmailValidationError()

	// then
	require.EqualError(t, err, "must be a valid email address")
}

func BenchmarkEmailAddressRule(b *testing.B) {
	runRuleBenchmarks(b, emailAddressRuleDataProvider)
}

func emailAddressRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		email1Dummy = fakerInstance.Internet().Email()
		email2Dummy = fmt.Sprintf(
			"%s <%s> (%s)",
			fakerInstance.Person().FirstName(),
			email1Dummy,
			fakerInstance.Lorem().Sentence(3),
		)
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             EmailAddress(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"email address": {
			rule:             EmailAddress(),
			value:            email1Dummy,
			expectedNewValue: email1Dummy,
			expectedError:    nil,
		},
		"pointer to email address": {
			rule:             EmailAddress(),
			value:            &email1Dummy,
			expectedNewValue: email1Dummy,
			expectedError:    nil,
		},
		"email address with name and comment": {
			rule:             EmailAddress(),
			value:            email2Dummy,
			expectedNewValue: email1Dummy,
			expectedError:    nil,
		},
		"pointer to email address with name and comment": {
			rule:             EmailAddress(),
			value:            &email2Dummy,
			expectedNewValue: email1Dummy,
			expectedError:    nil,
		},
		"invalid email address": {
			rule:             EmailAddress(),
			value:            "invalid email address",
			expectedNewValue: "invalid email address",
			expectedError:    NewEmailValidationError(),
		},

		// unsupported values
		"int": {
			rule:             EmailAddress(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewEmailValidationError(),
		},
		"float": {
			rule:             EmailAddress(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewEmailValidationError(),
		},
		"complex": {
			rule:             EmailAddress(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewEmailValidationError(),
		},
		"bool": {
			rule:             EmailAddress(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewEmailValidationError(),
		},
		"slice": {
			rule:             EmailAddress(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewEmailValidationError(),
		},
		"array": {
			rule:             EmailAddress(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewEmailValidationError(),
		},
		"map": {
			rule:             EmailAddress(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewEmailValidationError(),
		},
		"struct": {
			rule:             EmailAddress(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewEmailValidationError(),
		},
	}
}
