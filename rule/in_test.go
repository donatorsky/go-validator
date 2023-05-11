package rule

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_InRule(t *testing.T) {
	// given
	for ttName, tt := range inRuleDataProvider() {
		t.Run(ttName, func(t *testing.T) {
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
	runRuleBenchmarks(b, inRuleDataProvider)
}

func inRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		intDummy              = fakerInstance.IntBetween(-1000, 1000)
		intValuesDummy        = []int{intDummy, intDummy + 1}
		intPointerValuesDummy = []*int{&intDummy, ptr(intDummy + 1)}
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             In(intValuesDummy),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    NewInValidationError(intValuesDummy),
		},
		"nil with custom validator allowing it": {
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
		"existing int in empty slice": {
			rule:             In([]int{}),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    NewInValidationError([]int{}),
		},
		"existing int in slice of ints": {
			rule:             In(intValuesDummy),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    nil,
		},
		"existing int in slice of int pointers": {
			rule:             In(intPointerValuesDummy),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    NewInValidationError(intPointerValuesDummy),
		},
		"existing int in slice of ints with custom comparator": {
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
		"non-existing int in slice of ints": {
			rule:             In(intValuesDummy),
			value:            intDummy - 1,
			expectedNewValue: intDummy - 1,
			expectedError:    NewInValidationError(intValuesDummy),
		},
		"string in in slice of ints": {
			rule:             In(intValuesDummy),
			value:            "not an integer",
			expectedNewValue: "not an integer",
			expectedError:    NewInValidationError(intValuesDummy),
		},

		"existing int pointer in slice of ints": {
			rule:             In(intValuesDummy),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    nil,
		},
		"existing int pointer in slice int pointers": {
			rule:             In(intPointerValuesDummy),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    NewInValidationError(intPointerValuesDummy),
		},
		"existing int pointer in slice int pointers, without auto dereference": {
			rule:             In(intPointerValuesDummy, InRuleWithoutAutoDereference()),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    nil,
		},
		"existing int pointer in slice of ints, without auto dereference": {
			rule:             In(intValuesDummy, InRuleWithoutAutoDereference()),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    NewInValidationError(intValuesDummy),
		},
		"non-existing int pointer in slice of ints": {
			rule:             In(intValuesDummy),
			value:            ptr(intDummy - 1),
			expectedNewValue: ptr(intDummy - 1),
			expectedError:    NewInValidationError(intValuesDummy),
		},
	}
}
