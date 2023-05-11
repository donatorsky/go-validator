package rule

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_AfterOrEqualRule(t *testing.T) {
	runRuleTestCases(t, afterOrEqualRuleDataProvider)
}

func Test_AfterOrEqualValidationError(t *testing.T) {
	// given
	var afterOrEqualDummy = fakerInstance.Lorem().Sentence(6)

	// when
	err := NewAfterOrEqualValidationError(afterOrEqualDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must be a date after or equal to %s",
		afterOrEqualDummy,
	))
}

func BenchmarkAfterOrEqualRule(b *testing.B) {
	runRuleBenchmarks(b, afterOrEqualRuleDataProvider)
}

func afterOrEqualRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		nowDummy                          = time.Now()
		tomorrowDummy                     = time.Now().AddDate(0, 0, 1)
		yesterdayDummy                    = time.Now().AddDate(0, 0, -1)
		customAfterOrEqualComparable1Mock = newAfterOrEqualComparableMock(true, false)
		customAfterOrEqualComparable2Mock = newAfterOrEqualComparableMock(false, true)
		customAfterOrEqualComparable3Mock = newAfterOrEqualComparableMock(false, false)
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             AfterOrEqual(nowDummy),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"date yesterday": {
			rule:             AfterOrEqual(nowDummy),
			value:            yesterdayDummy,
			expectedNewValue: yesterdayDummy,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"date yesterday (pointer)": {
			rule:             AfterOrEqual(nowDummy),
			value:            &yesterdayDummy,
			expectedNewValue: &yesterdayDummy,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		"date today": {
			rule:             AfterOrEqual(nowDummy),
			value:            nowDummy,
			expectedNewValue: nowDummy,
			expectedError:    nil,
		},
		"date today (pointer)": {
			rule:             AfterOrEqual(nowDummy),
			value:            &nowDummy,
			expectedNewValue: &nowDummy,
			expectedError:    nil,
		},

		"date tomorrow": {
			rule:             AfterOrEqual(nowDummy),
			value:            tomorrowDummy,
			expectedNewValue: tomorrowDummy,
			expectedError:    nil,
		},
		"date tomorrow (pointer)": {
			rule:             AfterOrEqual(nowDummy),
			value:            &tomorrowDummy,
			expectedNewValue: &tomorrowDummy,
			expectedError:    nil,
		},

		"custom afterOrEqualComparable object: after": {
			rule:             AfterOrEqual(nowDummy),
			value:            customAfterOrEqualComparable1Mock,
			expectedNewValue: customAfterOrEqualComparable1Mock,
			expectedError:    nil,
		},
		"custom afterOrEqualComparable object: equal": {
			rule:             AfterOrEqual(nowDummy),
			value:            customAfterOrEqualComparable2Mock,
			expectedNewValue: customAfterOrEqualComparable2Mock,
			expectedError:    nil,
		},
		"custom afterOrEqualComparable object: before": {
			rule:             AfterOrEqual(nowDummy),
			value:            customAfterOrEqualComparable3Mock,
			expectedNewValue: customAfterOrEqualComparable3Mock,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		// unsupported values
		"int": {
			rule:             AfterOrEqual(nowDummy),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"float": {
			rule:             AfterOrEqual(nowDummy),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"complex": {
			rule:             AfterOrEqual(nowDummy),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"bool": {
			rule:             AfterOrEqual(nowDummy),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"slice": {
			rule:             AfterOrEqual(nowDummy),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"array": {
			rule:             AfterOrEqual(nowDummy),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"map": {
			rule:             AfterOrEqual(nowDummy),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"struct": {
			rule:             AfterOrEqual(nowDummy),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
	}
}

func newAfterOrEqualComparableMock(after, equal bool) *afterOrEqualComparableMock {
	return &afterOrEqualComparableMock{
		after: after,
		equal: equal,
	}
}

type afterOrEqualComparableMock struct {
	after bool
	equal bool
}

func (m *afterOrEqualComparableMock) After(_ time.Time) bool {
	return m.after
}

func (m *afterOrEqualComparableMock) Equal(_ time.Time) bool {
	return m.equal
}
