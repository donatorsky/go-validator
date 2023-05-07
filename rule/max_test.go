package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_MaxRule(t *testing.T) {
	// given
	for ttIdx, tt := range maxRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_MaxValidationError(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		error           error
		expectedMessage string
	}{
		{
			error:           NewMaxValidationError(ve.SubtypeString, 5),
			expectedMessage: "must be at most 5 characters",
		},
		{
			error:           NewMaxValidationError(ve.SubtypeString, 5.1),
			expectedMessage: "must be at most 5.1 characters",
		},

		{
			error:           NewMaxValidationError(ve.SubtypeSlice, 5),
			expectedMessage: "must have at most 5 items",
		},
		{
			error:           NewMaxValidationError(ve.SubtypeSlice, 5.1),
			expectedMessage: "must have at most 5.1 items",
		},

		{
			error:           NewMaxValidationError(ve.SubtypeMap, 5),
			expectedMessage: "must have at most 5 items",
		},
		{
			error:           NewMaxValidationError(ve.SubtypeMap, 5.1),
			expectedMessage: "must have at most 5.1 items",
		},

		{
			error:           NewMaxValidationError(ve.SubtypeNumber, 5),
			expectedMessage: "must be at most 5",
		},
		{
			error:           NewMaxValidationError(ve.SubtypeNumber, 5.1),
			expectedMessage: "must be at most 5.1",
		},
	} {
		t.Run(fmt.Sprintf("#%d", ttIdx), func(t *testing.T) {
			// then
			require.EqualError(t, tt.error, tt.expectedMessage)
		})
	}
}

func BenchmarkMaxRule(b *testing.B) {
	for ttIdx, tt := range maxRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func maxRuleDataProvider() []*ruleTestCaseData {
	return []*ruleTestCaseData{
		{
			rule:             Max(3),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		{
			rule:             Max(3),
			value:            "abc",
			expectedNewValue: "abc",
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            "abcd",
			expectedNewValue: "abcd",
			expectedError:    NewMaxValidationError(ve.SubtypeString, 3),
		},

		{
			rule:             Max(3),
			value:            3,
			expectedNewValue: 3,
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            4,
			expectedNewValue: 4,
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Max(3),
			value:            ptr(3),
			expectedNewValue: ptr(3),
			expectedError:    nil,
		},

		{
			rule:             Max(3),
			value:            int8(3),
			expectedNewValue: int8(3),
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            int8(4),
			expectedNewValue: int8(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            int16(3),
			expectedNewValue: int16(3),
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            int16(4),
			expectedNewValue: int16(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            int32(3),
			expectedNewValue: int32(3),
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            int32(4),
			expectedNewValue: int32(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            int64(3),
			expectedNewValue: int64(3),
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            int64(4),
			expectedNewValue: int64(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            uint(3),
			expectedNewValue: uint(3),
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            uint(4),
			expectedNewValue: uint(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            uint8(3),
			expectedNewValue: uint8(3),
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            uint8(4),
			expectedNewValue: uint8(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            uint16(3),
			expectedNewValue: uint16(3),
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            uint16(4),
			expectedNewValue: uint16(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            uint32(3),
			expectedNewValue: uint32(3),
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            uint32(4),
			expectedNewValue: uint32(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            uint64(3),
			expectedNewValue: uint64(3),
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            uint64(4),
			expectedNewValue: uint64(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            float32(3.0),
			expectedNewValue: float32(3.0),
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            float32(3.01),
			expectedNewValue: float32(3.01),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            3.0,
			expectedNewValue: 3.0,
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            3.01,
			expectedNewValue: 3.01,
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3),
		},

		{
			rule:             Max(3),
			value:            []int{1, 2, 3},
			expectedNewValue: []int{1, 2, 3},
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            []int{1, 2, 3, 4},
			expectedNewValue: []int{1, 2, 3, 4},
			expectedError:    NewMaxValidationError(ve.SubtypeSlice, 3),
		},

		{
			rule:             Max(3),
			value:            [3]int{1, 2, 3},
			expectedNewValue: [3]int{1, 2, 3},
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            [4]int{1, 2, 3, 4},
			expectedNewValue: [4]int{1, 2, 3, 4},
			expectedError:    NewMaxValidationError(ve.SubtypeSlice, 3),
		},

		{
			rule:             Max(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedError:    NewMaxValidationError(ve.SubtypeMap, 3),
		},

		// unsupported values
		{
			rule:             Max(3),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            true,
			expectedNewValue: true,
			expectedError:    nil,
		},
		{
			rule:             Max(3),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    nil,
		},
	}
}