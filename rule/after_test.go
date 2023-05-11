package rule

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_AfterRule(t *testing.T) {
	runRuleTestCases(t, afterRuleDataProvider)
}

func Test_AfterValidationError(t *testing.T) {
	// given
	var afterDummy = fakerInstance.Lorem().Sentence(6)

	// when
	err := NewAfterValidationError(afterDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must be a date after %s",
		afterDummy,
	))
}

func BenchmarkAfterRule(b *testing.B) {
	runRuleBenchmarks(b, afterRuleDataProvider)
}

func afterRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		nowDummy                   = time.Now()
		tomorrowDummy              = time.Now().AddDate(0, 0, 1)
		yesterdayDummy             = time.Now().AddDate(0, 0, -1)
		customAfterComparable1Mock = newAfterComparableMock(true)
		customAfterComparable2Mock = newAfterComparableMock(false)
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             After(nowDummy),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"date yesterday": {
			rule:             After(nowDummy),
			value:            yesterdayDummy,
			expectedNewValue: yesterdayDummy,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"date yesterday (pointer)": {
			rule:             After(nowDummy),
			value:            &yesterdayDummy,
			expectedNewValue: &yesterdayDummy,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		"date today": {
			rule:             After(nowDummy),
			value:            nowDummy,
			expectedNewValue: nowDummy,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"date today (pointer)": {
			rule:             After(nowDummy),
			value:            &nowDummy,
			expectedNewValue: &nowDummy,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		"date tomorrow": {
			rule:             After(nowDummy),
			value:            tomorrowDummy,
			expectedNewValue: tomorrowDummy,
			expectedError:    nil,
		},
		"date tomorrow (pointer)": {
			rule:             After(nowDummy),
			value:            &tomorrowDummy,
			expectedNewValue: &tomorrowDummy,
			expectedError:    nil,
		},

		"custom afterComparable object: after": {
			rule:             After(nowDummy),
			value:            customAfterComparable1Mock,
			expectedNewValue: customAfterComparable1Mock,
			expectedError:    nil,
		},
		"custom afterComparable object: not after": {
			rule:             After(nowDummy),
			value:            customAfterComparable2Mock,
			expectedNewValue: customAfterComparable2Mock,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},

		// unsupported values
		"int": {
			rule:             After(nowDummy),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"float": {
			rule:             After(nowDummy),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"complex": {
			rule:             After(nowDummy),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"bool": {
			rule:             After(nowDummy),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"slice": {
			rule:             After(nowDummy),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"array": {
			rule:             After(nowDummy),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"map": {
			rule:             After(nowDummy),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
		"struct": {
			rule:             After(nowDummy),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewAfterValidationError(nowDummy.Format(time.RFC3339Nano)),
		},
	}
}

func newAfterComparableMock(after bool) *afterComparableMock {
	return &afterComparableMock{
		after: after,
	}
}

type afterComparableMock struct {
	after bool
}

func (m *afterComparableMock) After(_ time.Time) bool {
	return m.after
}
