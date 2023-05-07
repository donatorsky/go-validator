package rule

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_InRule(t *testing.T) {
	// given
	for ttIdx, tt := range inRuleDataProvider() {
		t.Run(fmt.Sprintf("#%[1]d: for value %[2]T(%[2]v)", ttIdx, tt.value), func(t *testing.T) {
			// when
			newValue, err := tt.rule.Apply(context.Background(), tt.value, tt.data)

			// then
			if tt.expectedError == nil {
				require.NoError(t, err, "Rule is expected to not return error")
			} else {
				require.EqualError(t, err, tt.expectedError.Error(), "Rule returned unexpected error")
			}

			require.Equal(t, tt.expectedNewValue, newValue, "Rule returned unexpected value")
		})
	}
}

func Test_InValidationError(t *testing.T) {
	// given
	var valuesDummy = []any{1, "2", true}

	// when
	err := NewInValidationError(valuesDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf("does not exist in %v", valuesDummy))
}

func BenchmarkInRule(b *testing.B) {
	for ttIdx, tt := range inRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func inRuleDataProvider() []*ruleTestCaseData {
	var (
		intDummy              = fakerInstance.IntBetween(-1000, 1000)
		intValuesDummy        = []int{intDummy, intDummy + 1}
		intPointerValuesDummy = []*int{&intDummy, ptr(intDummy + 1)}
	)

	return []*ruleTestCaseData{
		{
			rule:             In(intValuesDummy),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    NewInValidationError(intValuesDummy),
		},
		{
			rule: In(intValuesDummy, InRuleWithComparator(func(value, expectedValue any) bool {
				if value == nil {
					return true
				}

				return value == expectedValue
			})),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},
		{
			rule:             In([]int{}),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    NewInValidationError([]int{}),
		},
		{
			rule:             In(intValuesDummy),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    nil,
		},
		{
			rule:             In(intPointerValuesDummy),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    NewInValidationError(intPointerValuesDummy),
		},
		{
			rule: In(intValuesDummy, InRuleWithComparator(func(value, expectedValue any) bool {
				if value == nil {
					return true
				}

				return value == expectedValue
			})),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    nil,
		},
		{
			rule:             In(intValuesDummy),
			value:            intDummy - 1,
			expectedNewValue: intDummy - 1,
			expectedError:    NewInValidationError(intValuesDummy),
		},
		{
			rule:             In(intValuesDummy),
			value:            "not an integer",
			expectedNewValue: "not an integer",
			expectedError:    NewInValidationError(intValuesDummy),
		},

		{
			rule:             In(intValuesDummy),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    nil,
		},
		{
			rule:             In(intPointerValuesDummy),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    NewInValidationError(intPointerValuesDummy),
		},
		{
			rule:             In(intPointerValuesDummy, InRuleWithoutAutoDereference()),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    nil,
		},
		{
			rule:             In(intValuesDummy, InRuleWithoutAutoDereference()),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    NewInValidationError(intValuesDummy),
		},
		{
			rule:             In(intValuesDummy),
			value:            ptr(intDummy - 1),
			expectedNewValue: ptr(intDummy - 1),
			expectedError:    NewInValidationError(intValuesDummy),
		},
	}
}
