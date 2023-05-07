package rule

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_BeforeOrEqualRule(t *testing.T) {
	// given
	for ttIdx, tt := range beforeOrEqualRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_BeforeOrEqualValidationError(t *testing.T) {
	// given
	var beforeOrEqualDummy = fakerInstance.Lorem().Sentence(6)

	// when
	err := NewBeforeOrEqualValidationError(beforeOrEqualDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must be a date before or equal to %s",
		beforeOrEqualDummy,
	))
}

func BenchmarkBeforeOrEqualRule(b *testing.B) {
	for ttIdx, tt := range beforeOrEqualRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func beforeOrEqualRuleDataProvider() []*ruleTestCaseData {
	var (
		nowDummy                           = time.Now()
		tomorrowDummy                      = time.Now().AddDate(0, 0, 1)
		yesterdayDummy                     = time.Now().AddDate(0, 0, -1)
		customBeforeOrEqualComparable1Mock = newBeforeOrEqualComparableMock(true, false)
		customBeforeOrEqualComparable2Mock = newBeforeOrEqualComparableMock(false, true)
		customBeforeOrEqualComparable3Mock = newBeforeOrEqualComparableMock(false, false)
	)

	return []*ruleTestCaseData{
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		{
			rule:             BeforeOrEqual(nowDummy),
			value:            yesterdayDummy,
			expectedNewValue: yesterdayDummy,
			expectedError:    nil,
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            &yesterdayDummy,
			expectedNewValue: &yesterdayDummy,
			expectedError:    nil,
		},

		{
			rule:             BeforeOrEqual(nowDummy),
			value:            nowDummy,
			expectedNewValue: nowDummy,
			expectedError:    nil,
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            &nowDummy,
			expectedNewValue: &nowDummy,
			expectedError:    nil,
		},

		{
			rule:             BeforeOrEqual(nowDummy),
			value:            tomorrowDummy,
			expectedNewValue: tomorrowDummy,
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            &tomorrowDummy,
			expectedNewValue: &tomorrowDummy,
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		{
			rule:             BeforeOrEqual(nowDummy),
			value:            customBeforeOrEqualComparable1Mock,
			expectedNewValue: customBeforeOrEqualComparable1Mock,
			expectedError:    nil,
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            customBeforeOrEqualComparable2Mock,
			expectedNewValue: customBeforeOrEqualComparable2Mock,
			expectedError:    nil,
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            customBeforeOrEqualComparable3Mock,
			expectedNewValue: customBeforeOrEqualComparable3Mock,
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		// unsupported values
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             BeforeOrEqual(nowDummy),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewBeforeOrEqualValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
	}
}

func newBeforeOrEqualComparableMock(before, equal bool) *beforeOrEqualComparableMock {
	return &beforeOrEqualComparableMock{
		before: before,
		equal:  equal,
	}
}

type beforeOrEqualComparableMock struct {
	before bool
	equal  bool
}

func (m *beforeOrEqualComparableMock) Before(_ time.Time) bool {
	return m.before
}

func (m *beforeOrEqualComparableMock) Equal(_ time.Time) bool {
	return m.equal
}
