package rule

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_DurationRule(t *testing.T) {
	runRuleTestCases(t, durationRuleDataProvider)
}

func Test_DurationValidationError(t *testing.T) {
	// when
	err := NewDurationValidationError()

	// then
	require.EqualError(t, err, "must be a valid duration")
}

func BenchmarkDurationRule(b *testing.B) {
	runRuleBenchmarks(b, durationRuleDataProvider)
}

func durationRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		durationDummy         = time.Minute + 2*time.Second
		durationAsStringDummy = "1m2s"
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Duration(),
			value:            nil,
			expectedNewValue: (*time.Duration)(nil),
			expectedError:    nil,
		},

		"time.Duration": {
			rule:             Duration(),
			value:            durationDummy,
			expectedNewValue: durationDummy,
			expectedError:    nil,
		},
		"*time.Duration": {
			rule:             Duration(),
			value:            &durationDummy,
			expectedNewValue: &durationDummy,
			expectedError:    nil,
		},

		"duration string": {
			rule:             Duration(),
			value:            durationAsStringDummy,
			expectedNewValue: durationDummy,
			expectedError:    nil,
		},
		"pointer to duration string": {
			rule:             Duration(),
			value:            &durationAsStringDummy,
			expectedNewValue: durationDummy,
			expectedError:    nil,
		},
		"invalid duration string": {
			rule:             Duration(),
			value:            "invalid duration",
			expectedNewValue: "invalid duration",
			expectedError:    NewDurationValidationError(),
		},

		// unsupported values
		"int": {
			rule:             Duration(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewDurationValidationError(),
		},
		"float": {
			rule:             Duration(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewDurationValidationError(),
		},
		"complex": {
			rule:             Duration(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewDurationValidationError(),
		},
		"bool": {
			rule:             Duration(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewDurationValidationError(),
		},
		"slice": {
			rule:             Duration(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewDurationValidationError(),
		},
		"array": {
			rule:             Duration(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewDurationValidationError(),
		},
		"map": {
			rule:             Duration(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewDurationValidationError(),
		},
		"struct": {
			rule:             Duration(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewDurationValidationError(),
		},
	}
}
