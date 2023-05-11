package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MapRule(t *testing.T) {
	runRuleTestCases(t, mapRuleDataProvider)
}

func Test_MapValidationError(t *testing.T) {
	// when
	err := NewMapValidationError()

	// then
	require.EqualError(t, err, "must be a map")
}

func BenchmarkMapRule(b *testing.B) {
	runRuleBenchmarks(b, mapRuleDataProvider)
}

func mapRuleDataProvider() map[string]*ruleTestCaseData {
	var mapDummy = map[int]int{0: 1, 1: 2, 2: 3}

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Map(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"pointer to map nil pointer": {
			rule:             Map(),
			value:            (*map[int]int)(nil),
			expectedNewValue: (*map[int]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"map": {
			rule:             Map(),
			value:            mapDummy,
			expectedNewValue: mapDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to map": {
			rule:             Map(),
			value:            &mapDummy,
			expectedNewValue: &mapDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		"int": {
			rule:             Map(),
			value:            1,
			expectedNewValue: nil,
			expectedError:    NewMapValidationError(),
			expectedToBail:   true,
		},
		"float": {
			rule:             Map(),
			value:            1.0,
			expectedNewValue: nil,
			expectedError:    NewMapValidationError(),
			expectedToBail:   true,
		},
		"complex": {
			rule:             Map(),
			value:            1 + 2i,
			expectedNewValue: nil,
			expectedError:    NewMapValidationError(),
			expectedToBail:   true,
		},
		"string": {
			rule:             Map(),
			value:            "foo",
			expectedNewValue: nil,
			expectedError:    NewMapValidationError(),
			expectedToBail:   true,
		},
		"bool": {
			rule:             Map(),
			value:            true,
			expectedNewValue: nil,
			expectedError:    NewMapValidationError(),
			expectedToBail:   true,
		},
		"slice": {
			rule:             Map(),
			value:            []int{},
			expectedNewValue: nil,
			expectedError:    NewMapValidationError(),
			expectedToBail:   true,
		},
		"array": {
			rule:             Map(),
			value:            [1]int{},
			expectedNewValue: nil,
			expectedError:    NewMapValidationError(),
			expectedToBail:   true,
		},
		"struct": {
			rule:             Map(),
			value:            someStruct{},
			expectedNewValue: nil,
			expectedError:    NewMapValidationError(),
			expectedToBail:   true,
		},
	}
}
