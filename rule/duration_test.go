package rule

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_DurationRule(t *testing.T) {
	// given
	for ttIdx, tt := range durationRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_DurationValidationError(t *testing.T) {
	// when
	err := NewDurationValidationError()

	// then
	require.EqualError(t, err, "must be a valid duration")
}

func BenchmarkDurationRule(b *testing.B) {
	for ttIdx, tt := range durationRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func durationRuleDataProvider() []*ruleTestCaseData {
	var (
		durationDummy         = time.Minute + 2*time.Second
		durationAsStringDummy = "1m2s"
	)

	return []*ruleTestCaseData{
		{
			rule:             Duration(),
			value:            nil,
			expectedNewValue: (*time.Duration)(nil),
			expectedError:    nil,
		},

		{
			rule:             Duration(),
			value:            durationDummy,
			expectedNewValue: durationDummy,
			expectedError:    nil,
		},
		{
			rule:             Duration(),
			value:            &durationDummy,
			expectedNewValue: &durationDummy,
			expectedError:    nil,
		},

		{
			rule:             Duration(),
			value:            durationAsStringDummy,
			expectedNewValue: durationDummy,
			expectedError:    nil,
		},
		{
			rule:             Duration(),
			value:            &durationAsStringDummy,
			expectedNewValue: durationDummy,
			expectedError:    nil,
		},
		{
			rule:             Duration(),
			value:            "invalid duration",
			expectedNewValue: "invalid duration",
			expectedError:    NewDurationValidationError(),
		},

		// unsupported values
		{
			rule:             Duration(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewDurationValidationError(),
		},
		{
			rule:             Duration(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewDurationValidationError(),
		},
		{
			rule:             Duration(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewDurationValidationError(),
		},
		{
			rule:             Duration(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewDurationValidationError(),
		},
		{
			rule:             Duration(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewDurationValidationError(),
		},
		{
			rule:             Duration(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewDurationValidationError(),
		},
		{
			rule:             Duration(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewDurationValidationError(),
		},
		{
			rule:             Duration(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewDurationValidationError(),
		},
	}
}
