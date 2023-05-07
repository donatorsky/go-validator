package rule

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_BeforeRule(t *testing.T) {
	// given
	for ttIdx, tt := range beforeRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_BeforeValidationError(t *testing.T) {
	// given
	var beforeDummy = fakerInstance.Lorem().Sentence(6)

	// when
	err := NewBeforeValidationError(beforeDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must be a date before %s",
		beforeDummy,
	))
}

func BenchmarkBeforeRule(b *testing.B) {
	for ttIdx, tt := range beforeRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func beforeRuleDataProvider() []*ruleTestCaseData {
	var (
		nowDummy                    = time.Now()
		tomorrowDummy               = time.Now().AddDate(0, 0, 1)
		yesterdayDummy              = time.Now().AddDate(0, 0, -1)
		customBeforeComparable1Mock = newBeforeComparableMock(true)
		customBeforeComparable2Mock = newBeforeComparableMock(false)
	)

	return []*ruleTestCaseData{
		{
			rule:             Before(nowDummy),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		{
			rule:             Before(nowDummy),
			value:            yesterdayDummy,
			expectedNewValue: yesterdayDummy,
			expectedError:    nil,
		},
		{
			rule:             Before(nowDummy),
			value:            &yesterdayDummy,
			expectedNewValue: &yesterdayDummy,
			expectedError:    nil,
		},

		{
			rule:             After(nowDummy),
			value:            nowDummy,
			expectedNewValue: nowDummy,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             After(nowDummy),
			value:            &nowDummy,
			expectedNewValue: &nowDummy,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		{
			rule:             Before(nowDummy),
			value:            tomorrowDummy,
			expectedNewValue: tomorrowDummy,
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             Before(nowDummy),
			value:            &tomorrowDummy,
			expectedNewValue: &tomorrowDummy,
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		{
			rule:             Before(nowDummy),
			value:            customBeforeComparable1Mock,
			expectedNewValue: customBeforeComparable1Mock,
			expectedError:    nil,
		},
		{
			rule:             Before(nowDummy),
			value:            customBeforeComparable2Mock,
			expectedNewValue: customBeforeComparable2Mock,
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		// unsupported values
		{
			rule:             Before(nowDummy),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             Before(nowDummy),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             Before(nowDummy),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             Before(nowDummy),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             Before(nowDummy),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             Before(nowDummy),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             Before(nowDummy),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		{
			rule:             Before(nowDummy),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewBeforeValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
	}
}

func newBeforeComparableMock(before bool) *beforeComparableMock {
	return &beforeComparableMock{
		before: before,
	}
}

type beforeComparableMock struct {
	before bool
}

func (m *beforeComparableMock) Before(_ time.Time) bool {
	return m.before
}
