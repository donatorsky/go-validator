package rule

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RegexRule(t *testing.T) {
	runRuleTestCases(t, regexRuleDataProvider)
}

func Test_RegexValidationError(t *testing.T) {
	// when
	err := NewRegexValidationError()

	// then
	require.EqualError(t, err, "format is invalid")
}

func BenchmarkRegexRule(b *testing.B) {
	runRuleBenchmarks(b, regexRuleDataProvider)
}

func regexRuleDataProvider() map[string]*ruleTestCaseData {
	var regexDummy = regexp.MustCompile(`^a.*`)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Regex(regexDummy),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"matching string": {
			rule:             Regex(regexDummy),
			value:            "abc",
			expectedNewValue: "abc",
			expectedError:    nil,
		},
		"matching *string": {
			rule:             Regex(regexDummy),
			value:            ptr("abc"),
			expectedNewValue: ptr("abc"),
			expectedError:    nil,
		},

		"not-matching string": {
			rule:             Regex(regexDummy),
			value:            "bc",
			expectedNewValue: "bc",
			expectedError:    NewRegexValidationError(),
		},
		"not-matching *string": {
			rule:             Regex(regexDummy),
			value:            ptr("bc"),
			expectedNewValue: ptr("bc"),
			expectedError:    NewRegexValidationError(),
		},

		// unsupported values
		"int": {
			rule:             Regex(regexDummy),
			value:            1,
			expectedNewValue: 1,
			expectedError:    NewRegexValidationError(),
		},
		"float": {
			rule:             Regex(regexDummy),
			value:            1.2,
			expectedNewValue: 1.2,
			expectedError:    NewRegexValidationError(),
		},
		"complex": {
			rule:             Regex(regexDummy),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewRegexValidationError(),
		},
		"bool": {
			rule:             Regex(regexDummy),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewRegexValidationError(),
		},
		"slice": {
			rule:             Regex(regexDummy),
			value:            []any{},
			expectedNewValue: []any{},
			expectedError:    NewRegexValidationError(),
		},
		"array": {
			rule:             Regex(regexDummy),
			value:            [1]any{},
			expectedNewValue: [1]any{},
			expectedError:    NewRegexValidationError(),
		},
		"map": {
			rule:             Regex(regexDummy),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewRegexValidationError(),
		},
		"struct": {
			rule:             Regex(regexDummy),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewRegexValidationError(),
		},
	}
}
