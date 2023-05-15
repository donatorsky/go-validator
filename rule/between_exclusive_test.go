package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_BetweenExclusiveRule(t *testing.T) {
	runRuleTestCases(t, betweenExclusiveRuleDataProvider)
}

func Test_BetweenValidationError_Exclusive(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		error           error
		expectedMessage string
	}{
		{
			error:           NewBetweenValidationError(ve.TypeString, 5, 6, false),
			expectedMessage: "must be between 5 and 6 characters (exclusive)",
		},
		{
			error:           NewBetweenValidationError(ve.TypeString, 5.1, 6.9, false),
			expectedMessage: "must be between 5.1 and 6.9 characters (exclusive)",
		},

		{
			error:           NewBetweenValidationError(ve.TypeSlice, 5, 6, false),
			expectedMessage: "must have between 5 and 6 items (exclusive)",
		},
		{
			error:           NewBetweenValidationError(ve.TypeSlice, 5.1, 6.9, false),
			expectedMessage: "must have between 5.1 and 6.9 items (exclusive)",
		},

		{
			error:           NewBetweenValidationError(ve.TypeMap, 5, 6, false),
			expectedMessage: "must have between 5 and 6 items (exclusive)",
		},
		{
			error:           NewBetweenValidationError(ve.TypeMap, 5.1, 6.9, false),
			expectedMessage: "must have between 5.1 and 6.9 items (exclusive)",
		},

		{
			error:           NewBetweenValidationError(ve.TypeNumber, 5, 6, false),
			expectedMessage: "must be between 5 and 6 (exclusive)",
		},
		{
			error:           NewBetweenValidationError(ve.TypeNumber, 5.1, 6.9, false),
			expectedMessage: "must be between 5.1 and 6.9 (exclusive)",
		},

		{
			error:           NewBetweenValidationError(ve.TypeInvalid, fakerInstance.Int(), fakerInstance.Int(), false),
			expectedMessage: "between cannot be determined",
		},
	} {
		t.Run(fmt.Sprintf("#%d", ttIdx), func(t *testing.T) {
			// then
			require.EqualError(t, tt.error, tt.expectedMessage)
		})
	}
}

func BenchmarkBetweenExclusiveRule(b *testing.B) {
	runRuleBenchmarks(b, betweenExclusiveRuleDataProvider)
}

func betweenExclusiveRuleDataProvider() map[string]*ruleTestCaseData {
	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             BetweenExclusive(2, 4),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string of length 2": {
			rule:             BetweenExclusive(2, 4),
			value:            "ab",
			expectedNewValue: "ab",
			expectedError:    NewBetweenValidationError(ve.TypeString, 2, 4, false),
		},
		"string of length 3": {
			rule:             BetweenExclusive(2, 4),
			value:            "abc",
			expectedNewValue: "abc",
			expectedError:    nil,
		},
		"pointer to string of length 3": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr("abc"),
			expectedNewValue: ptr("abc"),
			expectedError:    nil,
		},
		"string of length 4": {
			rule:             BetweenExclusive(2, 4),
			value:            "abcd",
			expectedNewValue: "abcd",
			expectedError:    NewBetweenValidationError(ve.TypeString, 2, 4, false),
		},

		"int(2)": {
			rule:             BetweenExclusive(2, 4),
			value:            2,
			expectedNewValue: 2,
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"int(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            3,
			expectedNewValue: 3,
			expectedError:    nil,
		},
		"*int(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(3),
			expectedNewValue: ptr(3),
			expectedError:    nil,
		},
		"int(4)": {
			rule:             BetweenExclusive(2, 4),
			value:            4,
			expectedNewValue: 4,
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"int8(2)": {
			rule:             BetweenExclusive(2, 4),
			value:            int8(2),
			expectedNewValue: int8(2),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"int8(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            int8(3),
			expectedNewValue: int8(3),
			expectedError:    nil,
		},
		"*int8(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(int8(3)),
			expectedNewValue: ptr(int8(3)),
			expectedError:    nil,
		},
		"int8(4)": {
			rule:             BetweenExclusive(2, 4),
			value:            int8(4),
			expectedNewValue: int8(4),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"int16(2)": {
			rule:             BetweenExclusive(2, 4),
			value:            int16(2),
			expectedNewValue: int16(2),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"int16(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            int16(3),
			expectedNewValue: int16(3),
			expectedError:    nil,
		},
		"*int16(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(int16(3)),
			expectedNewValue: ptr(int16(3)),
			expectedError:    nil,
		},
		"int16(4)": {
			rule:             BetweenExclusive(2, 4),
			value:            int16(4),
			expectedNewValue: int16(4),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"int32(2)": {
			rule:             BetweenExclusive(2, 4),
			value:            int32(2),
			expectedNewValue: int32(2),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"int32(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            int32(3),
			expectedNewValue: int32(3),
			expectedError:    nil,
		},
		"*int32(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(int32(3)),
			expectedNewValue: ptr(int32(3)),
			expectedError:    nil,
		},
		"int32(4)": {
			rule:             BetweenExclusive(2, 4),
			value:            int32(4),
			expectedNewValue: int32(4),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"int64(2)": {
			rule:             BetweenExclusive(2, 4),
			value:            int64(2),
			expectedNewValue: int64(2),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"int64(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            int64(3),
			expectedNewValue: int64(3),
			expectedError:    nil,
		},
		"*int64(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(int64(3)),
			expectedNewValue: ptr(int64(3)),
			expectedError:    nil,
		},
		"int64(4)": {
			rule:             BetweenExclusive(2, 4),
			value:            int64(4),
			expectedNewValue: int64(4),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"uint(2)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint(2),
			expectedNewValue: uint(2),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"uint(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint(3),
			expectedNewValue: uint(3),
			expectedError:    nil,
		},
		"*uint(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(uint(3)),
			expectedNewValue: ptr(uint(3)),
			expectedError:    nil,
		},
		"uint(4)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint(4),
			expectedNewValue: uint(4),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"uint8(2)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint8(2),
			expectedNewValue: uint8(2),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"uint8(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint8(3),
			expectedNewValue: uint8(3),
			expectedError:    nil,
		},
		"*uint8(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(uint8(3)),
			expectedNewValue: ptr(uint8(3)),
			expectedError:    nil,
		},
		"uint8(4)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint8(4),
			expectedNewValue: uint8(4),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"uint16(2)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint16(2),
			expectedNewValue: uint16(2),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"uint16(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint16(3),
			expectedNewValue: uint16(3),
			expectedError:    nil,
		},
		"*uint16(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(uint16(3)),
			expectedNewValue: ptr(uint16(3)),
			expectedError:    nil,
		},
		"uint16(4)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint16(4),
			expectedNewValue: uint16(4),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"uint32(2)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint32(2),
			expectedNewValue: uint32(2),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"uint32(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint32(3),
			expectedNewValue: uint32(3),
			expectedError:    nil,
		},
		"*uint32(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(uint32(3)),
			expectedNewValue: ptr(uint32(3)),
			expectedError:    nil,
		},
		"uint32(4)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint32(4),
			expectedNewValue: uint32(4),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"uint64(2)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint64(2),
			expectedNewValue: uint64(2),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"uint64(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint64(3),
			expectedNewValue: uint64(3),
			expectedError:    nil,
		},
		"*uint64(3)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(uint64(3)),
			expectedNewValue: ptr(uint64(3)),
			expectedError:    nil,
		},
		"uint64(4)": {
			rule:             BetweenExclusive(2, 4),
			value:            uint64(4),
			expectedNewValue: uint64(4),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"float32(2.0)": {
			rule:             BetweenExclusive(2, 4),
			value:            float32(2.0),
			expectedNewValue: float32(2.0),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"float32(2.01)": {
			rule:             BetweenExclusive(2, 4),
			value:            float32(2.01),
			expectedNewValue: float32(2.01),
			expectedError:    nil,
		},
		"*float32(3.0)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(float32(3.0)),
			expectedNewValue: ptr(float32(3.0)),
			expectedError:    nil,
		},
		"float32(3.99)": {
			rule:             BetweenExclusive(2, 4),
			value:            float32(3.99),
			expectedNewValue: float32(3.99),
			expectedError:    nil,
		},
		"float32(4.0)": {
			rule:             BetweenExclusive(2, 4),
			value:            float32(4.0),
			expectedNewValue: float32(4.0),
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"float64(2.0)": {
			rule:             BetweenExclusive(2, 4),
			value:            2.0,
			expectedNewValue: 2.0,
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},
		"float64(2.01)": {
			rule:             BetweenExclusive(2, 4),
			value:            2.01,
			expectedNewValue: 2.01,
			expectedError:    nil,
		},
		"*float64(3.0)": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(3.0),
			expectedNewValue: ptr(3.0),
			expectedError:    nil,
		},
		"float64(3.99)": {
			rule:             BetweenExclusive(2, 4),
			value:            3.99,
			expectedNewValue: 3.99,
			expectedError:    nil,
		},
		"float64(4.0)": {
			rule:             BetweenExclusive(2, 4),
			value:            4.0,
			expectedNewValue: 4.0,
			expectedError:    NewBetweenValidationError(ve.TypeNumber, 2, 4, false),
		},

		"slice with 2 items": {
			rule:             BetweenExclusive(2, 4),
			value:            []int{1, 2},
			expectedNewValue: []int{1, 2},
			expectedError:    NewBetweenValidationError(ve.TypeSlice, 2, 4, false),
		},
		"slice with 3 items": {
			rule:             BetweenExclusive(2, 4),
			value:            []int{1, 2, 3},
			expectedNewValue: []int{1, 2, 3},
			expectedError:    nil,
		},
		"pointer to slice with 3 items": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr([]int{1, 2, 3}),
			expectedNewValue: ptr([]int{1, 2, 3}),
			expectedError:    nil,
		},
		"slice with 4 items": {
			rule:             BetweenExclusive(2, 4),
			value:            []int{1, 2, 3, 4},
			expectedNewValue: []int{1, 2, 3, 4},
			expectedError:    NewBetweenValidationError(ve.TypeSlice, 2, 4, false),
		},

		"array with 2 items": {
			rule:             BetweenExclusive(2, 4),
			value:            [2]int{1, 2},
			expectedNewValue: [2]int{1, 2},
			expectedError:    NewBetweenValidationError(ve.TypeArray, 2, 4, false),
		},
		"array with 3 items": {
			rule:             BetweenExclusive(2, 4),
			value:            [3]int{1, 2, 3},
			expectedNewValue: [3]int{1, 2, 3},
			expectedError:    nil,
		},
		"pointer to array with 3 items": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr([3]int{1, 2}),
			expectedNewValue: ptr([3]int{1, 2}),
			expectedError:    nil,
		},
		"array with 4 items": {
			rule:             BetweenExclusive(2, 4),
			value:            [4]int{1, 2, 3, 4},
			expectedNewValue: [4]int{1, 2, 3, 4},
			expectedError:    NewBetweenValidationError(ve.TypeArray, 2, 4, false),
		},

		"map with 2 keys": {
			rule:             BetweenExclusive(2, 4),
			value:            map[any]int{1: 1, "a": 2},
			expectedNewValue: map[any]int{1: 1, "a": 2},
			expectedError:    NewBetweenValidationError(ve.TypeMap, 2, 4, false),
		},
		"map with 3 keys": {
			rule:             BetweenExclusive(2, 4),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedError:    nil,
		},
		"pointer to map with 3 keys": {
			rule:             BetweenExclusive(2, 4),
			value:            ptr(map[any]int{1: 1, "a": 2, 3.0: 3}),
			expectedNewValue: ptr(map[any]int{1: 1, "a": 2, 3.0: 3}),
			expectedError:    nil,
		},
		"map with 4 keys": {
			rule:             BetweenExclusive(2, 4),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedError:    NewBetweenValidationError(ve.TypeMap, 2, 4, false),
		},

		// unsupported values
		"complex": {
			rule:             BetweenExclusive(2, 4),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewBetweenValidationError(ve.TypeInvalid, 2, 4, false),
		},
		"bool": {
			rule:             BetweenExclusive(2, 4),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewBetweenValidationError(ve.TypeInvalid, 2, 4, false),
		},
		"struct": {
			rule:             BetweenExclusive(2, 4),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewBetweenValidationError(ve.TypeInvalid, 2, 4, false),
		},
	}
}
