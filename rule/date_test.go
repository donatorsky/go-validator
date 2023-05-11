package rule

import (
	"testing"
	"time"
)

func Test_DateRule(t *testing.T) {
	runRuleTestCases(t, dateRuleDataProvider)
}

func BenchmarkDateRule(b *testing.B) {
	runRuleBenchmarks(b, dateRuleDataProvider)
}

func dateRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		dateDummy         = time.Now()
		dateAsStringDummy = dateDummy.Format(time.RFC3339Nano)
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Date(),
			value:            nil,
			expectedNewValue: (*time.Time)(nil),
			expectedError:    nil,
		},

		"time.Time": {
			rule:             Date(),
			value:            dateDummy,
			expectedNewValue: dateDummy,
			expectedError:    nil,
		},
		"*time.Time": {
			rule:             Date(),
			value:            &dateDummy,
			expectedNewValue: &dateDummy,
			expectedError:    nil,
		},

		"date string": {
			rule:  Date(),
			value: dateAsStringDummy,
			expectedNewValueFunc: func(value any) bool {
				return dateDummy.Equal(value.(time.Time))
			},
			expectedError: nil,
		},
		"pointer to date string": {
			rule:  Date(),
			value: &dateAsStringDummy,
			expectedNewValueFunc: func(value any) bool {
				return dateDummy.Equal(value.(time.Time))
			},
			expectedError: nil,
		},
		"invalid date string": {
			rule:             Date(),
			value:            "invalid date",
			expectedNewValue: "invalid date",
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},

		// unsupported values
		"int": {
			rule:             Date(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		"float": {
			rule:             Date(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		"complex": {
			rule:             Date(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		"bool": {
			rule:             Date(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		"slice": {
			rule:             Date(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		"array": {
			rule:             Date(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		"map": {
			rule:             Date(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
		"struct": {
			rule:             Date(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewDateFormatValidationError(time.RFC3339Nano),
		},
	}
}
