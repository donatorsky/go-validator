package rule

import (
	"testing"
	"time"
)

func Test_DateRule(t *testing.T) {
	// given
	for ttIdx, tt := range dateRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func BenchmarkDateRule(b *testing.B) {
	for ttIdx, tt := range dateRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func dateRuleDataProvider() []*ruleTestCaseData {
	var (
		dateDummy         = time.Now()
		dateAsStringDummy = dateDummy.Format(time.RFC3339Nano)
	)

	return []*ruleTestCaseData{
		{
			rule:             Date(),
			value:            nil,
			expectedNewValue: (*time.Time)(nil),
			expectedError:    nil,
		},

		{
			rule:             Date(),
			value:            dateDummy,
			expectedNewValue: dateDummy,
			expectedError:    nil,
		},
		{
			rule:             Date(),
			value:            &dateDummy,
			expectedNewValue: &dateDummy,
			expectedError:    nil,
		},

		{
			rule:  Date(),
			value: dateAsStringDummy,
			expectedNewValueFunc: func(value any) bool {
				return dateDummy.Equal(value.(time.Time))
			},
			expectedError: nil,
		},
		{
			rule:  Date(),
			value: &dateAsStringDummy,
			expectedNewValueFunc: func(value any) bool {
				return dateDummy.Equal(value.(time.Time))
			},
			expectedError: nil,
		},
		{
			rule:             Date(),
			value:            "invalid date",
			expectedNewValue: "invalid date",
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},

		// unsupported values
		{
			rule:             Date(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		{
			rule:             Date(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		{
			rule:             Date(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		{
			rule:             Date(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		{
			rule:             Date(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		{
			rule:             Date(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		{
			rule:             Date(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		{
			rule:             Date(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
	}
}
