package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_BetweenRule(t *testing.T) {
	// given
	for ttIdx, tt := range betweenRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_BetweenValidationError(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		error           error
		expectedMessage string
	}{
		{
			error:           NewBetweenValidationError(ve.TypeBetween, ve.SubtypeString, 5, 6, true),
			expectedMessage: "must be between 5 and 6 characters (inclusive)",
		},
		{
			error:           NewBetweenValidationError(ve.TypeBetween, ve.SubtypeString, 5.1, 6.9, true),
			expectedMessage: "must be between 5.1 and 6.9 characters (inclusive)",
		},

		{
			error:           NewBetweenValidationError(ve.TypeBetween, ve.SubtypeSlice, 5, 6, true),
			expectedMessage: "must have between 5 and 6 items (inclusive)",
		},
		{
			error:           NewBetweenValidationError(ve.TypeBetween, ve.SubtypeSlice, 5.1, 6.9, true),
			expectedMessage: "must have between 5.1 and 6.9 items (inclusive)",
		},

		{
			error:           NewBetweenValidationError(ve.TypeBetween, ve.SubtypeMap, 5, 6, true),
			expectedMessage: "must have between 5 and 6 items (inclusive)",
		},
		{
			error:           NewBetweenValidationError(ve.TypeBetween, ve.SubtypeMap, 5.1, 6.9, true),
			expectedMessage: "must have between 5.1 and 6.9 items (inclusive)",
		},

		{
			error:           NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 5, 6, true),
			expectedMessage: "must be between 5 and 6 (inclusive)",
		},
		{
			error:           NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 5.1, 6.9, true),
			expectedMessage: "must be between 5.1 and 6.9 (inclusive)",
		},
	} {
		t.Run(fmt.Sprintf("#%d", ttIdx), func(t *testing.T) {
			// then
			require.EqualError(t, tt.error, tt.expectedMessage)
		})
	}
}

func BenchmarkBetweenRule(b *testing.B) {
	for ttIdx, tt := range betweenRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func betweenRuleDataProvider() []*ruleTestCaseData {
	return []*ruleTestCaseData{
		{
			rule:             Between(2, 4),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		{
			rule:             Between(2, 4),
			value:            "a",
			expectedNewValue: "a",
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeString, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            "ab",
			expectedNewValue: "ab",
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr("abc"),
			expectedNewValue: ptr("abc"),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            "abcd",
			expectedNewValue: "abcd",
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            "abcde",
			expectedNewValue: "abcde",
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeString, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            1,
			expectedNewValue: 1,
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            2,
			expectedNewValue: 2,
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            4,
			expectedNewValue: 4,
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(3),
			expectedNewValue: ptr(3),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            5,
			expectedNewValue: 5,
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            int8(1),
			expectedNewValue: int8(1),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            int8(2),
			expectedNewValue: int8(2),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            int8(4),
			expectedNewValue: int8(4),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(int8(3)),
			expectedNewValue: ptr(int8(3)),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            int8(5),
			expectedNewValue: int8(5),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            int16(1),
			expectedNewValue: int16(1),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            int16(2),
			expectedNewValue: int16(2),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            int16(4),
			expectedNewValue: int16(4),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(int16(3)),
			expectedNewValue: ptr(int16(3)),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            int16(5),
			expectedNewValue: int16(5),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            int32(1),
			expectedNewValue: int32(1),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            int32(2),
			expectedNewValue: int32(2),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            int32(4),
			expectedNewValue: int32(4),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(int32(3)),
			expectedNewValue: ptr(int32(3)),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            int32(5),
			expectedNewValue: int32(5),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            int64(1),
			expectedNewValue: int64(1),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            int64(2),
			expectedNewValue: int64(2),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            int64(4),
			expectedNewValue: int64(4),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(int64(3)),
			expectedNewValue: ptr(int64(3)),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            int64(5),
			expectedNewValue: int64(5),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            uint(1),
			expectedNewValue: uint(1),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            uint(2),
			expectedNewValue: uint(2),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            uint(4),
			expectedNewValue: uint(4),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(uint(3)),
			expectedNewValue: ptr(uint(3)),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            uint(5),
			expectedNewValue: uint(5),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            uint8(1),
			expectedNewValue: uint8(1),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            uint8(2),
			expectedNewValue: uint8(2),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            uint8(4),
			expectedNewValue: uint8(4),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(uint8(3)),
			expectedNewValue: ptr(uint8(3)),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            uint8(5),
			expectedNewValue: uint8(5),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            uint16(1),
			expectedNewValue: uint16(1),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            uint16(2),
			expectedNewValue: uint16(2),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            uint16(4),
			expectedNewValue: uint16(4),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(uint16(3)),
			expectedNewValue: ptr(uint16(3)),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            uint16(5),
			expectedNewValue: uint16(5),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            uint32(1),
			expectedNewValue: uint32(1),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            uint32(2),
			expectedNewValue: uint32(2),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            uint32(4),
			expectedNewValue: uint32(4),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(uint32(3)),
			expectedNewValue: ptr(uint32(3)),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            uint32(5),
			expectedNewValue: uint32(5),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            uint64(1),
			expectedNewValue: uint64(1),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            uint64(2),
			expectedNewValue: uint64(2),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            uint64(4),
			expectedNewValue: uint64(4),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(uint64(3)),
			expectedNewValue: ptr(uint64(3)),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            uint64(5),
			expectedNewValue: uint64(5),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            float32(1.99),
			expectedNewValue: float32(1.99),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            float32(2.0),
			expectedNewValue: float32(2.0),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            float32(4.0),
			expectedNewValue: float32(4.0),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(float32(3.0)),
			expectedNewValue: ptr(float32(3.0)),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            float32(4.01),
			expectedNewValue: float32(4.01),
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            1.99,
			expectedNewValue: 1.99,
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            2.0,
			expectedNewValue: 2.0,
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            4.0,
			expectedNewValue: 4.0,
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(3.0),
			expectedNewValue: ptr(3.0),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            4.01,
			expectedNewValue: 4.01,
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeNumber, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            []int{1},
			expectedNewValue: []int{1},
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeSlice, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            []int{1, 2},
			expectedNewValue: []int{1, 2},
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr([]int{1, 2, 3}),
			expectedNewValue: ptr([]int{1, 2, 3}),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            []int{1, 2, 3, 4},
			expectedNewValue: []int{1, 2, 3, 4},
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            []int{1, 2, 3, 4, 5},
			expectedNewValue: []int{1, 2, 3, 4, 5},
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeSlice, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            [1]int{1},
			expectedNewValue: [1]int{1},
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeSlice, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            [2]int{1, 2},
			expectedNewValue: [2]int{1, 2},
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr([3]int{1, 2}),
			expectedNewValue: ptr([3]int{1, 2}),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            [4]int{1, 2, 3, 4},
			expectedNewValue: [4]int{1, 2, 3, 4},
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            [5]int{1, 2, 3, 4, 5},
			expectedNewValue: [5]int{1, 2, 3, 4, 5},
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeSlice, 2, 4, true),
		},

		{
			rule:             Between(2, 4),
			value:            map[any]int{1: 1},
			expectedNewValue: map[any]int{1: 1},
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeMap, 2, 4, true),
		},
		{
			rule:             Between(2, 4),
			value:            map[any]int{1: 1, "a": 2},
			expectedNewValue: map[any]int{1: 1, "a": 2},
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            ptr(map[any]int{1: 1, "a": 2, 3.0: 3}),
			expectedNewValue: ptr(map[any]int{1: 1, "a": 2, 3.0: 3}),
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3, true: 4, 5i: 5},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3, true: 4, 5i: 5},
			expectedError:    NewBetweenValidationError(ve.TypeBetween, ve.SubtypeMap, 2, 4, true),
		},

		// unsupported values
		{
			rule:             Between(2, 4),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            true,
			expectedNewValue: true,
			expectedError:    nil,
		},
		{
			rule:             Between(2, 4),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    nil,
		},
	}
}
