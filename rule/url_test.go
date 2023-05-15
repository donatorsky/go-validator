package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UrlRule(t *testing.T) {
	runRuleTestCases(t, urlRuleDataProvider)
}

func Test_UrlValidationError(t *testing.T) {
	// when
	err := NewUrlValidationError()

	// then
	require.EqualError(t, err, "must be a valid URL format")
}

func BenchmarkUrlRule(b *testing.B) {
	runRuleBenchmarks(b, urlRuleDataProvider)
}

func urlRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		urlDummy        = fakerInstance.Internet().URL()
		invalidUrlDummy = "invalid-url"
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             URL(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string with valid URL": {
			rule:             URL(),
			value:            urlDummy,
			expectedNewValue: urlDummy,
			expectedError:    nil,
		},
		"*string with valid URL": {
			rule:             URL(),
			value:            &urlDummy,
			expectedNewValue: &urlDummy,
			expectedError:    nil,
		},

		"string with invalid URL": {
			rule:             URL(),
			value:            invalidUrlDummy,
			expectedNewValue: invalidUrlDummy,
			expectedError:    NewUrlValidationError(),
		},
		"*string with invalid URL": {
			rule:             URL(),
			value:            &invalidUrlDummy,
			expectedNewValue: &invalidUrlDummy,
			expectedError:    NewUrlValidationError(),
		},

		// unsupported values
		"int": {
			rule:             URL(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewUrlValidationError(),
		},
		"float": {
			rule:             URL(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewUrlValidationError(),
		},
		"complex": {
			rule:             URL(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewUrlValidationError(),
		},
		"bool": {
			rule:             URL(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewUrlValidationError(),
		},
		"slice": {
			rule:             URL(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewUrlValidationError(),
		},
		"array": {
			rule:             URL(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewUrlValidationError(),
		},
		"map": {
			rule:             URL(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewUrlValidationError(),
		},
		"struct": {
			rule:             URL(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewUrlValidationError(),
		},
	}
}
