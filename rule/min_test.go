package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_MinRule(t *testing.T) {
	runRuleTestCases(t, minRuleDataProvider)
}

func Test_MinValidationError(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		error           error
		expectedMessage string
	}{
		{
			error:           NewMinValidationError(ve.SubtypeString, 5, true),
			expectedMessage: "must be at least 5 characters",
		},
		{
			error:           NewMinValidationError(ve.SubtypeString, 5.1, true),
			expectedMessage: "must be at least 5.1 characters",
		},

		{
			error:           NewMinValidationError(ve.SubtypeSlice, 5, true),
			expectedMessage: "must have at least 5 items",
		},
		{
			error:           NewMinValidationError(ve.SubtypeSlice, 5.1, true),
			expectedMessage: "must have at least 5.1 items",
		},

		{
			error:           NewMinValidationError(ve.SubtypeArray, 5, true),
			expectedMessage: "must have at least 5 items",
		},
		{
			error:           NewMinValidationError(ve.SubtypeArray, 5.1, true),
			expectedMessage: "must have at least 5.1 items",
		},

		{
			error:           NewMinValidationError(ve.SubtypeMap, 5, true),
			expectedMessage: "must have at least 5 items",
		},
		{
			error:           NewMinValidationError(ve.SubtypeMap, 5.1, true),
			expectedMessage: "must have at least 5.1 items",
		},

		{
			error:           NewMinValidationError(ve.SubtypeNumber, 5, true),
			expectedMessage: "must be at least 5",
		},
		{
			error:           NewMinValidationError(ve.SubtypeNumber, 5.1, true),
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
	runRuleBenchmarks(b, minRuleDataProvider)
}

func minRuleDataProvider() map[string]*ruleTestCaseData {
	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Min(3),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string of length 2": {
			rule:             Min(3),
			value:            "ab",
			expectedNewValue: "ab",
			expectedError:    NewMinValidationError(ve.SubtypeString, 3, true),
		},
		"string of length 3": {
			rule:             Min(3),
			value:            "abc",
			expectedNewValue: "abc",
			expectedError:    nil,
		},
		"*string of length 2": {
			rule:             Min(3),
			value:            ptr("ab"),
			expectedNewValue: ptr("ab"),
			expectedError:    NewMinValidationError(ve.SubtypeString, 3, true),
		},
		"*string of length 3": {
			rule:             Min(3),
			value:            ptr("abc"),
			expectedNewValue: ptr("abc"),
			expectedError:    nil,
		},

		"int(2)": {
			rule:             Min(3),
			value:            2,
			expectedNewValue: 2,
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"int(3)": {
			rule:             Min(3),
			value:            3,
			expectedNewValue: 3,
			expectedError:    nil,
		},
		"*int(2)": {
			rule:             Min(3),
			value:            ptr(2),
			expectedNewValue: ptr(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*int(3)": {
			rule:             Min(3),
			value:            ptr(3),
			expectedNewValue: ptr(3),
			expectedError:    nil,
		},

		"int8(2)": {
			rule:             Min(3),
			value:            int8(2),
			expectedNewValue: int8(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"int8(3)": {
			rule:             Min(3),
			value:            int8(3),
			expectedNewValue: int8(3),
			expectedError:    nil,
		},
		"*int8(2)": {
			rule:             Min(3),
			value:            ptr(int8(2)),
			expectedNewValue: ptr(int8(2)),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*int8(3)": {
			rule:             Min(3),
			value:            ptr(int8(3)),
			expectedNewValue: ptr(int8(3)),
			expectedError:    nil,
		},

		"int16(2)": {
			rule:             Min(3),
			value:            int16(2),
			expectedNewValue: int16(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"int16(3)": {
			rule:             Min(3),
			value:            int16(3),
			expectedNewValue: int16(3),
			expectedError:    nil,
		},
		"*int16(2)": {
			rule:             Min(3),
			value:            ptr(int16(2)),
			expectedNewValue: ptr(int16(2)),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*int16(3)": {
			rule:             Min(3),
			value:            ptr(int16(3)),
			expectedNewValue: ptr(int16(3)),
			expectedError:    nil,
		},

		"int32(2)": {
			rule:             Min(3),
			value:            int32(2),
			expectedNewValue: int32(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"int32(3)": {
			rule:             Min(3),
			value:            int32(3),
			expectedNewValue: int32(3),
			expectedError:    nil,
		},
		"*int32(2)": {
			rule:             Min(3),
			value:            ptr(int32(2)),
			expectedNewValue: ptr(int32(2)),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*int32(3)": {
			rule:             Min(3),
			value:            ptr(int32(3)),
			expectedNewValue: ptr(int32(3)),
			expectedError:    nil,
		},

		"int64(2)": {
			rule:             Min(3),
			value:            int64(2),
			expectedNewValue: int64(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"int64(3)": {
			rule:             Min(3),
			value:            int64(3),
			expectedNewValue: int64(3),
			expectedError:    nil,
		},
		"*int64(2)": {
			rule:             Min(3),
			value:            ptr(int64(2)),
			expectedNewValue: ptr(int64(2)),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*int64(3)": {
			rule:             Min(3),
			value:            ptr(int64(3)),
			expectedNewValue: ptr(int64(3)),
			expectedError:    nil,
		},

		"uint(2)": {
			rule:             Min(3),
			value:            uint(2),
			expectedNewValue: uint(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"uint(3)": {
			rule:             Min(3),
			value:            uint(3),
			expectedNewValue: uint(3),
			expectedError:    nil,
		},
		"*uint(2)": {
			rule:             Min(3),
			value:            ptr(uint(2)),
			expectedNewValue: ptr(uint(2)),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*uint(3)": {
			rule:             Min(3),
			value:            ptr(uint(3)),
			expectedNewValue: ptr(uint(3)),
			expectedError:    nil,
		},

		"uint8(2)": {
			rule:             Min(3),
			value:            uint8(2),
			expectedNewValue: uint8(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"uint8(3)": {
			rule:             Min(3),
			value:            uint8(3),
			expectedNewValue: uint8(3),
			expectedError:    nil,
		},
		"*uint8(2)": {
			rule:             Min(3),
			value:            ptr(uint8(2)),
			expectedNewValue: ptr(uint8(2)),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*uint8(3)": {
			rule:             Min(3),
			value:            ptr(uint8(3)),
			expectedNewValue: ptr(uint8(3)),
			expectedError:    nil,
		},

		"uint16(2)": {
			rule:             Min(3),
			value:            uint16(2),
			expectedNewValue: uint16(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"uint16(3)": {
			rule:             Min(3),
			value:            uint16(3),
			expectedNewValue: uint16(3),
			expectedError:    nil,
		},
		"*uint16(2)": {
			rule:             Min(3),
			value:            ptr(uint16(2)),
			expectedNewValue: ptr(uint16(2)),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*uint16(3)": {
			rule:             Min(3),
			value:            ptr(uint16(3)),
			expectedNewValue: ptr(uint16(3)),
			expectedError:    nil,
		},

		"uint32(2)": {
			rule:             Min(3),
			value:            uint32(2),
			expectedNewValue: uint32(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"uint32(3)": {
			rule:             Min(3),
			value:            uint32(3),
			expectedNewValue: uint32(3),
			expectedError:    nil,
		},
		"*uint32(2)": {
			rule:             Min(3),
			value:            ptr(uint32(2)),
			expectedNewValue: ptr(uint32(2)),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*uint32(3)": {
			rule:             Min(3),
			value:            ptr(uint32(3)),
			expectedNewValue: ptr(uint32(3)),
			expectedError:    nil,
		},

		"uint64(2)": {
			rule:             Min(3),
			value:            uint64(2),
			expectedNewValue: uint64(2),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"uint64(3)": {
			rule:             Min(3),
			value:            uint64(3),
			expectedNewValue: uint64(3),
			expectedError:    nil,
		},
		"*uint64(2)": {
			rule:             Min(3),
			value:            ptr(uint64(2)),
			expectedNewValue: ptr(uint64(2)),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*uint64(3)": {
			rule:             Min(3),
			value:            ptr(uint64(3)),
			expectedNewValue: ptr(uint64(3)),
			expectedError:    nil,
		},

		"float32(2.99)": {
			rule:             Min(3),
			value:            float32(2.99),
			expectedNewValue: float32(2.99),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"float32(3.0)": {
			rule:             Min(3),
			value:            float32(3.0),
			expectedNewValue: float32(3.0),
			expectedError:    nil,
		},
		"*float32(2.99)": {
			rule:             Min(3),
			value:            ptr(float32(2.99)),
			expectedNewValue: ptr(float32(2.99)),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*float32(3.0)": {
			rule:             Min(3),
			value:            ptr(float32(3.0)),
			expectedNewValue: ptr(float32(3.0)),
			expectedError:    nil,
		},

		"float64(2.99)": {
			rule:             Min(3),
			value:            2.99,
			expectedNewValue: 2.99,
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"float64(3.0)": {
			rule:             Min(3),
			value:            3.0,
			expectedNewValue: 3.0,
			expectedError:    nil,
		},
		"*float64(2.99)": {
			rule:             Min(3),
			value:            ptr(2.99),
			expectedNewValue: ptr(2.99),
			expectedError:    NewMinValidationError(ve.SubtypeNumber, 3, true),
		},
		"*float64(3.0)": {
			rule:             Min(3),
			value:            ptr(3.0),
			expectedNewValue: ptr(3.0),
			expectedError:    nil,
		},

		"slice with 2 items": {
			rule:             Min(3),
			value:            []int{1, 2},
			expectedNewValue: []int{1, 2},
			expectedError:    NewMinValidationError(ve.SubtypeSlice, 3, true),
		},
		"slice with 3 items": {
			rule:             Min(3),
			value:            []int{1, 2, 3},
			expectedNewValue: []int{1, 2, 3},
			expectedError:    nil,
		},

		"array with 2 items": {
			rule:             Min(3),
			value:            [2]int{1, 2},
			expectedNewValue: [2]int{1, 2},
			expectedError:    NewMinValidationError(ve.SubtypeArray, 3, true),
		},
		"array with 3 items": {
			rule:             Min(3),
			value:            [3]int{1, 2, 3},
			expectedNewValue: [3]int{1, 2, 3},
			expectedError:    nil,
		},

		"map with 2 keys": {
			rule:             Min(3),
			value:            map[any]int{1: 1, "a": 2},
			expectedNewValue: map[any]int{1: 1, "a": 2},
			expectedError:    NewMinValidationError(ve.SubtypeMap, 3, true),
		},
		"map with 3 keys": {
			rule:             Min(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedError:    nil,
		},

		// unsupported values
		"complex": {
			rule:             Min(3),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    nil,
		},
		"bool": {
			rule:             Min(3),
			value:            true,
			expectedNewValue: true,
			expectedError:    nil,
		},
		"struct": {
			rule:             Min(3),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    nil,
		},
	}
}
