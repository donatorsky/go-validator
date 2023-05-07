package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_MinRule(t *testing.T) {
	// given
	for ttIdx, tt := range minRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_MinValidationError(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		error           error
		expectedMessage string
	}{
		{
			error:           NewMinValidationError(ve.SubtypeString, 5),
			expectedMessage: "must be at least 5 characters",
		},
		{
			error:           NewMinValidationError(ve.SubtypeString, 5.1),
			expectedMessage: "must be at least 5.1 characters",
		},

		{
			error:           NewMinValidationError(ve.SubtypeSlice, 5),
			expectedMessage: "must have at least 5 items",
		},
		{
			error:           NewMinValidationError(ve.SubtypeSlice, 5.1),
			expectedMessage: "must have at least 5.1 items",
		},

		{
			error:           NewMinValidationError(ve.SubtypeMap, 5),
			expectedMessage: "must have at least 5 items",
		},
		{
			error:           NewMinValidationError(ve.SubtypeMap, 5.1),
			expectedMessage: "must have at least 5.1 items",
		},

		{
			error:           NewMinValidationError(ve.SubtypeNumber, 5),
			expectedMessage: "must be at least 5",
		},
		{
			error:           NewMinValidationError(ve.SubtypeNumber, 5.1),
			expectedMessage: "must be at least 5.1",
		},
	} {
		t.Run(fmt.Sprintf("#%d", ttIdx), func(t *testing.T) {
			// then
			require.EqualError(t, tt.error, tt.expectedMessage)
		})
	}
}

func BenchmarkMinRule(b *testing.B) {
	for ttIdx, tt := range minRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func minRuleDataProvider() []*ruleTestCaseData {
	return []*ruleTestCaseData{
		{
			rule:             Min(3),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            "ab",
			expectedNewValue: "ab",
			expectedError:    NewMinValidationError(ve.SubtypeString, 3),
		},
		{
			rule:             Min(3),
			value:            "abc",
			expectedNewValue: "abc",
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            2,
			expectedNewValue: 2,
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            3,
			expectedNewValue: 3,
			expectedError:    nil,
		},
		{
			rule:             Min(3),
			value:            ptr(3),
			expectedNewValue: ptr(3),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            int8(2),
			expectedNewValue: int8(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            int8(3),
			expectedNewValue: int8(3),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            int16(2),
			expectedNewValue: int16(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            int16(3),
			expectedNewValue: int16(3),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            int32(2),
			expectedNewValue: int32(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            int32(3),
			expectedNewValue: int32(3),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            int64(2),
			expectedNewValue: int64(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            int64(3),
			expectedNewValue: int64(3),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            uint(2),
			expectedNewValue: uint(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            uint(3),
			expectedNewValue: uint(3),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            uint8(2),
			expectedNewValue: uint8(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            uint8(3),
			expectedNewValue: uint8(3),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            uint16(2),
			expectedNewValue: uint16(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            uint16(3),
			expectedNewValue: uint16(3),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            uint32(2),
			expectedNewValue: uint32(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            uint32(3),
			expectedNewValue: uint32(3),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            uint64(2),
			expectedNewValue: uint64(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            uint64(3),
			expectedNewValue: uint64(3),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            float32(2.99),
			expectedNewValue: float32(2.99),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            float32(3.0),
			expectedNewValue: float32(3.0),
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            2.99,
			expectedNewValue: 2.99,
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3),
		},
		{
			rule:             Min(3),
			value:            3.0,
			expectedNewValue: 3.0,
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            []int{1, 2},
			expectedNewValue: []int{1, 2},
			expectedError:    NewMinValidationError(ve.SubtypeSlice, 3),
		},
		{
			rule:             Min(3),
			value:            []int{1, 2, 3},
			expectedNewValue: []int{1, 2, 3},
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            [2]int{1, 2},
			expectedNewValue: [2]int{1, 2},
			expectedError:    NewMinValidationError(ve.SubtypeSlice, 3),
		},
		{
			rule:             Min(3),
			value:            [3]int{1, 2, 3},
			expectedNewValue: [3]int{1, 2, 3},
			expectedError:    nil,
		},

		{
			rule:             Min(3),
			value:            map[any]int{1: 1, "a": 2},
			expectedNewValue: map[any]int{1: 1, "a": 2},
			expectedError:    NewMinValidationError(ve.SubtypeMap, 3),
		},
		{
			rule:             Min(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedError:    nil,
		},

		// unsupported values
		{
			rule:             Min(3),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    nil,
		},
		{
			rule:             Min(3),
			value:            true,
			expectedNewValue: true,
			expectedError:    nil,
		},
		{
			rule:             Min(3),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    nil,
		},
	}
}
