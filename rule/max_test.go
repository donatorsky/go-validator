package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_MaxRule(t *testing.T) {
	runRuleTestCases(t, maxRuleDataProvider)
}

func Test_MaxValidationError(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		error           error
		expectedMessage string
	}{
		{
			error:           NewMaxValidationError(ve.SubtypeString, 5, true),
			expectedMessage: "must be at most 5 characters",
		},
		{
			error:           NewMaxValidationError(ve.SubtypeString, 5.1, true),
			expectedMessage: "must be at most 5.1 characters",
		},

		{
			error:           NewMaxValidationError(ve.SubtypeSlice, 5, true),
			expectedMessage: "must have at most 5 items",
		},
		{
			error:           NewMaxValidationError(ve.SubtypeSlice, 5.1, true),
			expectedMessage: "must have at most 5.1 items",
		},

		{
			error:           NewMaxValidationError(ve.SubtypeArray, 5, true),
			expectedMessage: "must have at most 5 items",
		},
		{
			error:           NewMaxValidationError(ve.SubtypeArray, 5.1, true),
			expectedMessage: "must have at most 5.1 items",
		},

		{
			error:           NewMaxValidationError(ve.SubtypeMap, 5, true),
			expectedMessage: "must have at most 5 items",
		},
		{
			error:           NewMaxValidationError(ve.SubtypeMap, 5.1, true),
			expectedMessage: "must have at most 5.1 items",
		},

		{
			error:           NewMaxValidationError(ve.SubtypeNumber, 5, true),
			expectedMessage: "must be at most 5",
		},
		{
			error:           NewMaxValidationError(ve.SubtypeNumber, 5.1, true),
			expectedMessage: "must be at most 5.1",
		},

		{
			error:           NewMaxValidationError(ve.SubtypeInvalid, fakerInstance.Int(), true),
			expectedMessage: "max cannot be determined",
		},
	} {
		t.Run(fmt.Sprintf("#%d", ttIdx), func(t *testing.T) {
			// then
			require.EqualError(t, tt.error, tt.expectedMessage)
		})
	}
}

func BenchmarkMaxRule(b *testing.B) {
	runRuleBenchmarks(b, maxRuleDataProvider)
}

func maxRuleDataProvider() map[string]*ruleTestCaseData {
	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Max(3),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string of length 3": {
			rule:             Max(3),
			value:            "abc",
			expectedNewValue: "abc",
			expectedError:    nil,
		},
		"string of length 4": {
			rule:             Max(3),
			value:            "abcd",
			expectedNewValue: "abcd",
			expectedError:    NewMaxValidationError(ve.SubtypeString, 3, true),
		},
		"*string of length 3": {
			rule:             Max(3),
			value:            ptr("abc"),
			expectedNewValue: ptr("abc"),
			expectedError:    nil,
		},
		"*string of length 4": {
			rule:             Max(3),
			value:            ptr("abcd"),
			expectedNewValue: ptr("abcd"),
			expectedError:    NewMaxValidationError(ve.SubtypeString, 3, true),
		},

		"int(3)": {
			rule:             Max(3),
			value:            3,
			expectedNewValue: 3,
			expectedError:    nil,
		},
		"int(4)": {
			rule:             Max(3),
			value:            4,
			expectedNewValue: 4,
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*int(3)": {
			rule:             Max(3),
			value:            ptr(3),
			expectedNewValue: ptr(3),
			expectedError:    nil,
		},
		"*int(4)": {
			rule:             Max(3),
			value:            ptr(4),
			expectedNewValue: ptr(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"int8(3)": {
			rule:             Max(3),
			value:            int8(3),
			expectedNewValue: int8(3),
			expectedError:    nil,
		},
		"int8(4)": {
			rule:             Max(3),
			value:            int8(4),
			expectedNewValue: int8(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*int8(3)": {
			rule:             Max(3),
			value:            ptr(int8(3)),
			expectedNewValue: ptr(int8(3)),
			expectedError:    nil,
		},
		"*int8(4)": {
			rule:             Max(3),
			value:            ptr(int8(4)),
			expectedNewValue: ptr(int8(4)),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"int16(3)": {
			rule:             Max(3),
			value:            int16(3),
			expectedNewValue: int16(3),
			expectedError:    nil,
		},
		"int16(4)": {
			rule:             Max(3),
			value:            int16(4),
			expectedNewValue: int16(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*int16(3)": {
			rule:             Max(3),
			value:            ptr(int16(3)),
			expectedNewValue: ptr(int16(3)),
			expectedError:    nil,
		},
		"*int16(4)": {
			rule:             Max(3),
			value:            ptr(int16(4)),
			expectedNewValue: ptr(int16(4)),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"int32(3)": {
			rule:             Max(3),
			value:            int32(3),
			expectedNewValue: int32(3),
			expectedError:    nil,
		},
		"int32(4)": {
			rule:             Max(3),
			value:            int32(4),
			expectedNewValue: int32(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*int32(3)": {
			rule:             Max(3),
			value:            ptr(int32(3)),
			expectedNewValue: ptr(int32(3)),
			expectedError:    nil,
		},
		"*int32(4)": {
			rule:             Max(3),
			value:            ptr(int32(4)),
			expectedNewValue: ptr(int32(4)),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"int64(3)": {
			rule:             Max(3),
			value:            int64(3),
			expectedNewValue: int64(3),
			expectedError:    nil,
		},
		"int64(4)": {
			rule:             Max(3),
			value:            int64(4),
			expectedNewValue: int64(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*int64(3)": {
			rule:             Max(3),
			value:            ptr(int64(3)),
			expectedNewValue: ptr(int64(3)),
			expectedError:    nil,
		},
		"*int64(4)": {
			rule:             Max(3),
			value:            ptr(int64(4)),
			expectedNewValue: ptr(int64(4)),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"uint(3)": {
			rule:             Max(3),
			value:            uint(3),
			expectedNewValue: uint(3),
			expectedError:    nil,
		},
		"uint(4)": {
			rule:             Max(3),
			value:            uint(4),
			expectedNewValue: uint(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*uint(3)": {
			rule:             Max(3),
			value:            ptr(uint(3)),
			expectedNewValue: ptr(uint(3)),
			expectedError:    nil,
		},
		"*uint(4)": {
			rule:             Max(3),
			value:            ptr(uint(4)),
			expectedNewValue: ptr(uint(4)),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"uint8(3)": {
			rule:             Max(3),
			value:            uint8(3),
			expectedNewValue: uint8(3),
			expectedError:    nil,
		},
		"uint8(4)": {
			rule:             Max(3),
			value:            uint8(4),
			expectedNewValue: uint8(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*uint8(3)": {
			rule:             Max(3),
			value:            ptr(uint8(3)),
			expectedNewValue: ptr(uint8(3)),
			expectedError:    nil,
		},
		"*uint8(4)": {
			rule:             Max(3),
			value:            ptr(uint8(4)),
			expectedNewValue: ptr(uint8(4)),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"uint16(3)": {
			rule:             Max(3),
			value:            uint16(3),
			expectedNewValue: uint16(3),
			expectedError:    nil,
		},
		"uint16(4)": {
			rule:             Max(3),
			value:            uint16(4),
			expectedNewValue: uint16(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*uint16(3)": {
			rule:             Max(3),
			value:            ptr(uint16(3)),
			expectedNewValue: ptr(uint16(3)),
			expectedError:    nil,
		},
		"*uint16(4)": {
			rule:             Max(3),
			value:            ptr(uint16(4)),
			expectedNewValue: ptr(uint16(4)),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"uint32(3)": {
			rule:             Max(3),
			value:            uint32(3),
			expectedNewValue: uint32(3),
			expectedError:    nil,
		},
		"uint32(4)": {
			rule:             Max(3),
			value:            uint32(4),
			expectedNewValue: uint32(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*uint32(3)": {
			rule:             Max(3),
			value:            ptr(uint32(3)),
			expectedNewValue: ptr(uint32(3)),
			expectedError:    nil,
		},
		"*uint32(4)": {
			rule:             Max(3),
			value:            ptr(uint32(4)),
			expectedNewValue: ptr(uint32(4)),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"uint64(3)": {
			rule:             Max(3),
			value:            uint64(3),
			expectedNewValue: uint64(3),
			expectedError:    nil,
		},
		"uint64(4)": {
			rule:             Max(3),
			value:            uint64(4),
			expectedNewValue: uint64(4),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*uint64(3)": {
			rule:             Max(3),
			value:            ptr(uint64(3)),
			expectedNewValue: ptr(uint64(3)),
			expectedError:    nil,
		},
		"*uint64(4)": {
			rule:             Max(3),
			value:            ptr(uint64(4)),
			expectedNewValue: ptr(uint64(4)),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"float32(3.0)": {
			rule:             Max(3),
			value:            float32(3.0),
			expectedNewValue: float32(3.0),
			expectedError:    nil,
		},
		"float32(3.01)": {
			rule:             Max(3),
			value:            float32(3.01),
			expectedNewValue: float32(3.01),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*float32(3.0)": {
			rule:             Max(3),
			value:            ptr(float32(3.0)),
			expectedNewValue: ptr(float32(3.0)),
			expectedError:    nil,
		},
		"*float32(3.01)": {
			rule:             Max(3),
			value:            ptr(float32(3.01)),
			expectedNewValue: ptr(float32(3.01)),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"float64(3.0)": {
			rule:             Max(3),
			value:            3.0,
			expectedNewValue: 3.0,
			expectedError:    nil,
		},
		"float64(3.01)": {
			rule:             Max(3),
			value:            3.01,
			expectedNewValue: 3.01,
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},
		"*float64(3.0)": {
			rule:             Max(3),
			value:            ptr(3.0),
			expectedNewValue: ptr(3.0),
			expectedError:    nil,
		},
		"*float64(3.01)": {
			rule:             Max(3),
			value:            ptr(3.01),
			expectedNewValue: ptr(3.01),
			expectedError:    NewMaxValidationError(ve.SubtypeNumber, 3, true),
		},

		"slice with 3 items": {
			rule:             Max(3),
			value:            []int{1, 2, 3},
			expectedNewValue: []int{1, 2, 3},
			expectedError:    nil,
		},
		"slice with 4 items": {
			rule:             Max(3),
			value:            []int{1, 2, 3, 4},
			expectedNewValue: []int{1, 2, 3, 4},
			expectedError:    NewMaxValidationError(ve.SubtypeSlice, 3, true),
		},

		"array with 3 items": {
			rule:             Max(3),
			value:            [3]int{1, 2, 3},
			expectedNewValue: [3]int{1, 2, 3},
			expectedError:    nil,
		},
		"array with 4 items": {
			rule:             Max(3),
			value:            [4]int{1, 2, 3, 4},
			expectedNewValue: [4]int{1, 2, 3, 4},
			expectedError:    NewMaxValidationError(ve.SubtypeArray, 3, true),
		},

		"map with 3 keys": {
			rule:             Max(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedError:    nil,
		},
		"map with 4 keys": {
			rule:             Max(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedError:    NewMaxValidationError(ve.SubtypeMap, 3, true),
		},

		// unsupported values
		"complex": {
			rule:             Max(3),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewMaxValidationError(ve.SubtypeInvalid, 3, true),
		},
		"bool": {
			rule:             Max(3),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewMaxValidationError(ve.SubtypeInvalid, 3, true),
		},
		"struct": {
			rule:             Max(3),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewMaxValidationError(ve.SubtypeInvalid, 3, true),
		},
	}
}
