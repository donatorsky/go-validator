package rule

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_AfterOrEqualRule(t *testing.T) {
	// given
	for ttIdx, tt := range afterOrEqualRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
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
	for ttIdx, tt := range afterOrEqualRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func afterOrEqualRuleDataProvider() []*ruleTestCaseData {
	var (
		nowDummy                          = time.Now()
		tomorrowDummy                     = time.Now().AddDate(0, 0, 1)
		yesterdayDummy                    = time.Now().AddDate(0, 0, -1)
		customAfterOrEqualComparable1Mock = newAfterOrEqualComparableMock(true, false)
		customAfterOrEqualComparable2Mock = newAfterOrEqualComparableMock(false, true)
		customAfterOrEqualComparable3Mock = newAfterOrEqualComparableMock(false, false)
	)

	return []*ruleTestCaseData{
		{
			rule:             AfterOrEqual(nowDummy),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		{
			rule:             AfterOrEqual(nowDummy),
			value:            yesterdayDummy,
			expectedNewValue: yesterdayDummy,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            &yesterdayDummy,
			expectedNewValue: &yesterdayDummy,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		{
			rule:             AfterOrEqual(nowDummy),
			value:            nowDummy,
			expectedNewValue: nowDummy,
			expectedError:    nil,
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            &nowDummy,
			expectedNewValue: &nowDummy,
			expectedError:    nil,
		},

		{
			rule:             AfterOrEqual(nowDummy),
			value:            tomorrowDummy,
			expectedNewValue: tomorrowDummy,
			expectedError:    nil,
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            &tomorrowDummy,
			expectedNewValue: &tomorrowDummy,
			expectedError:    nil,
		},

		{
			rule:             AfterOrEqual(nowDummy),
			value:            customAfterOrEqualComparable1Mock,
			expectedNewValue: customAfterOrEqualComparable1Mock,
			expectedError:    nil,
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            customAfterOrEqualComparable2Mock,
			expectedNewValue: customAfterOrEqualComparable2Mock,
			expectedError:    nil,
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            customAfterOrEqualComparable3Mock,
			expectedNewValue: customAfterOrEqualComparable3Mock,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		// unsupported values
		{
			rule:             AfterOrEqual(nowDummy),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             AfterOrEqual(nowDummy),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewAfterOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
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
