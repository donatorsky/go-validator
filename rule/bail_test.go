package rule

import (
	"testing"
)

func Test_BailRule(t *testing.T) {
	// given
	for ttIdx, tt := range bailRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func BenchmarkBailRule(b *testing.B) {
	for ttIdx, tt := range bailRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func bailRuleDataProvider() []*ruleTestCaseData {
	return []*ruleTestCaseData{
		{
			rule:             Bail(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		{
			rule:             Bail(),
			value:            1,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		{
			rule:             Bail(),
			value:            1.2,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		{
			rule:             Bail(),
			value:            1 + 2i,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		{
			rule:             Bail(),
			value:            true,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		{
			rule:             Bail(),
			value:            []int{},
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		{
			rule:             Bail(),
			value:            [1]int{},
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
		{
			rule:             Bail(),
			value:            map[string]int{},
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   true,
		},
	}
}
