package rule

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NotInRule(t *testing.T) {
	// given
	for ttName, tt := range notInRuleDataProvider() {
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

func Test_NotInValidationError(t *testing.T) {
	// given
	var valuesDummy = []any{1, "2", true}

	// when
	err := NewNotInValidationError(valuesDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf("exists in %v", valuesDummy))
}

func BenchmarkNotInRule(b *testing.B) {
	runRuleBenchmarks(b, notInRuleDataProvider)
}

func notInRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		intDummy              = fakerInstance.IntBetween(-1000, 1000)
		intValuesDummy        = []int{intDummy, intDummy + 1}
		intPointerValuesDummy = []*int{&intDummy, ptr(intDummy + 1)}
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             NotIn(intValuesDummy),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},
		"nil, without auto dereference": {
			rule:             NotIn(intValuesDummy, NotInRuleWithoutAutoDereference()),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},
		"nil with custom validator allowing it": {
			rule: NotIn(
				intValuesDummy,
				NotInRuleWithoutAutoDereference(),
				NotInRuleWithComparator(func(value, expectedValue any) bool {
					if value == nil {
						return true
					}

					return value == expectedValue
				}),
			),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    NewNotInValidationError(intValuesDummy),
		},
		"existing int in empty slice": {
			rule:             NotIn([]int{}),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    nil,
		},
		"existing int in slice of ints": {
			rule:             NotIn(intValuesDummy),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    NewNotInValidationError(intValuesDummy),
		},
		"existing int in slice of int pointers": {
			rule:             NotIn(intPointerValuesDummy),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    nil,
		},
		"existing int in slice of ints with custom comparator": {
			rule: NotIn(intValuesDummy, NotInRuleWithComparator(func(value, expectedValue any) bool {
				return value == expectedValue
			})),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    NewNotInValidationError(intValuesDummy),
		},
		"non-existing int in slice of ints": {
			rule:             NotIn(intValuesDummy),
			value:            intDummy - 1,
			expectedNewValue: intDummy - 1,
			expectedError:    nil,
		},
		"string in slice of ints": {
			rule:             NotIn(intValuesDummy),
			value:            "not an integer",
			expectedNewValue: "not an integer",
			expectedError:    nil,
		},

		"existing int pointer in slice of ints": {
			rule:             NotIn(intValuesDummy),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    NewNotInValidationError(intValuesDummy),
		},
		"existing int pointer in slice int pointers": {
			rule:             NotIn(intPointerValuesDummy),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    nil,
		},
		"existing int pointer in slice int pointers, without auto dereference": {
			rule:             NotIn(intPointerValuesDummy, NotInRuleWithoutAutoDereference()),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    NewNotInValidationError(intPointerValuesDummy),
		},
		"existing int pointer in slice of ints, without auto dereference": {
			rule:             NotIn(intValuesDummy, NotInRuleWithoutAutoDereference()),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    nil,
		},
		"non-existing int pointer in slice of ints": {
			rule:             NotIn(intValuesDummy),
			value:            ptr(intDummy - 1),
			expectedNewValue: ptr(intDummy - 1),
			expectedError:    nil,
		},
	}
}
