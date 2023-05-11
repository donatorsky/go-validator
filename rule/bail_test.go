package rule

import (
	"testing"
)

func Test_BailRule(t *testing.T) {
	runRuleTestCases(t, bailRuleDataProvider)
}

func BenchmarkBailRule(b *testing.B) {
	runRuleBenchmarks(b, bailRuleDataProvider)
}

func bailRuleDataProvider() map[string]*ruleTestCaseData {
	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Bail(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		"int": {
			rule:             Bail(),
			value:            1,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		"float": {
			rule:             Bail(),
			value:            1.2,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		"complex": {
			rule:             Bail(),
			value:            1 + 2i,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		"bool": {
			rule:             Bail(),
			value:            true,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		"slice": {
			rule:             Bail(),
			value:            []int{},
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		"array": {
			rule:             Bail(),
			value:            [1]int{},
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		"map": {
			rule:             Bail(),
			value:            map[string]int{},
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		"struct": {
			rule:             Bail(),
			value:            someStruct{},
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
	}
}
