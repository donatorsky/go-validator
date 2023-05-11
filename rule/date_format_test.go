package rule

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_DateFormatRule(t *testing.T) {
	runRuleTestCases(t, dateFormatRuleDataProvider)
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
	runRuleBenchmarks(b, dateFormatRuleDataProvider)
}

func dateFormatRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		dateDummy         = time.Now().Truncate(time.Second)
		dateAsStringDummy = dateDummy.Format(time.RFC1123Z)
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             DateFormat(time.RFC1123Z),
			value:            nil,
			expectedNewValue: (*time.Time)(nil),
			expectedError:    nil,
		},

		"time.Time": {
			rule:             DateFormat(time.RFC1123Z),
			value:            dateDummy,
			expectedNewValue: dateDummy,
			expectedError:    nil,
		},
		"*time.Time": {
			rule:             DateFormat(time.RFC1123Z),
			value:            &dateDummy,
			expectedNewValue: &dateDummy,
			expectedError:    nil,
		},

		"date string": {
			rule:             DateFormat(time.RFC1123Z),
			value:            dateAsStringDummy,
			expectedNewValue: dateDummy,
			expectedError:    nil,
		},
		"pointer to date string": {
			rule:             DateFormat(time.RFC1123Z),
			value:            &dateAsStringDummy,
			expectedNewValue: dateDummy,
			expectedError:    nil,
		},
		"invalid date string": {
			rule:             DateFormat(time.RFC1123Z),
			value:            "invalid date",
			expectedNewValue: "invalid date",
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},

		// unsupported values
		"int": {
			rule:             DateFormat(time.RFC1123Z),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		"float": {
			rule:             DateFormat(time.RFC1123Z),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		"complex": {
			rule:             DateFormat(time.RFC1123Z),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		"bool": {
			rule:             DateFormat(time.RFC1123Z),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		"slice": {
			rule:             DateFormat(time.RFC1123Z),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		"array": {
			rule:             DateFormat(time.RFC1123Z),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		"map": {
			rule:             DateFormat(time.RFC1123Z),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
		"struct": {
			rule:             DateFormat(time.RFC1123Z),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewDateFormatValidationError(time.RFC1123Z),
		},
	}
}
