package rule

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_DateFormatRule(t *testing.T) {
	// given
	for ttIdx, tt := range dateFormatRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_DateFormatValidationError(t *testing.T) {
	// given
	var dateFormatDummy = fakerInstance.Lorem().Sentence(6)

	// when
	err := NewDateFormatValidationError(dateFormatDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"does not match the date format %s",
		dateFormatDummy,
	))
}

func BenchmarkDateFormatRule(b *testing.B) {
	for ttIdx, tt := range dateFormatRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func dateFormatRuleDataProvider() []*ruleTestCaseData {
	var (
		dateDummy         = time.Now().Truncate(time.Second)
		dateAsStringDummy = dateDummy.Format(time.RFC1123Z)
	)

	return []*ruleTestCaseData{
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            nil,
			expectedNewValue: (*time.Time)(nil),
			expectedError:    nil,
		},

		{
			rule:             DateFormat(time.RFC1123Z),
			value:            dateDummy,
			expectedNewValue: dateDummy,
			expectedError:    nil,
		},
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            &dateDummy,
			expectedNewValue: &dateDummy,
			expectedError:    nil,
		},

		{
			rule:             DateFormat(time.RFC1123Z),
			value:            dateAsStringDummy,
			expectedNewValue: dateDummy,
			expectedError:    nil,
		},
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            &dateAsStringDummy,
			expectedNewValue: dateDummy,
			expectedError:    nil,
		},
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            "invalid date",
			expectedNewValue: "invalid date",
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},

		// unsupported values
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		{
			rule:             DateFormat(time.RFC1123Z),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
	}
}
