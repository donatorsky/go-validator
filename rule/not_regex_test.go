package rule

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NotRegexRule(t *testing.T) {
	runRuleTestCases(t, notRegexRuleDataProvider)
}

func Test_NotRegexValidationError(t *testing.T) {
	// when
	err := NewNotRegexValidationError()

	// then
	require.EqualError(t, err, "format is invalid")
}

func BenchmarkNotRegexRule(b *testing.B) {
	runRuleBenchmarks(b, notRegexRuleDataProvider)
}

func notRegexRuleDataProvider() map[string]*ruleTestCaseData {
	var regexDummy = regexp.MustCompile(`^a.*`)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             NotRegex(regexDummy),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"matching string": {
			rule:             NotRegex(regexDummy),
			value:            "abc",
			expectedNewValue: "abc",
			expectedError:    NewNotRegexValidationError(),
		},
		"matching *string": {
			rule:             NotRegex(regexDummy),
			value:            ptr("abc"),
			expectedNewValue: ptr("abc"),
			expectedError:    NewNotRegexValidationError(),
		},

		"not-matching string": {
			rule:             NotRegex(regexDummy),
			value:            "bc",
			expectedNewValue: "bc",
			expectedError:    nil,
		},
		"not-matching *string": {
			rule:             NotRegex(regexDummy),
			value:            ptr("bc"),
			expectedNewValue: ptr("bc"),
			expectedError:    nil,
		},

		// unsupported values
		"int": {
			rule:             NotRegex(regexDummy),
			value:            1,
			expectedNewValue: 1,
			expectedError:    NewNotRegexValidationError(),
		},
		"float": {
			rule:             NotRegex(regexDummy),
			value:            1.2,
			expectedNewValue: 1.2,
			expectedError:    NewNotRegexValidationError(),
		},
		"complex": {
			rule:             NotRegex(regexDummy),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewNotRegexValidationError(),
		},
		"bool": {
			rule:             NotRegex(regexDummy),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewNotRegexValidationError(),
		},
		"slice": {
			rule:             NotRegex(regexDummy),
			value:            []any{},
			expectedNewValue: []any{},
			expectedError:    NewNotRegexValidationError(),
		},
		"array": {
			rule:             NotRegex(regexDummy),
			value:            [1]any{},
			expectedNewValue: [1]any{},
			expectedError:    NewNotRegexValidationError(),
		},
		"map": {
			rule:             NotRegex(regexDummy),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewNotRegexValidationError(),
		},
		"struct": {
			rule:             NotRegex(regexDummy),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewNotRegexValidationError(),
		},
	}
}
