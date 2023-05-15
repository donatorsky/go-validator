package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_MaxExclusiveRule(t *testing.T) {
	runRuleTestCases(t, maxExclusiveRuleDataProvider)
}

func Test_MaxExclusiveValidationError(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		error           error
		expectedMessage string
	}{
		{
			error:           NewMaxValidationError(ve.TypeString, 5, false),
			expectedMessage: "must be less than 5 characters",
		},
		{
			error:           NewMaxValidationError(ve.TypeString, 5.1, false),
			expectedMessage: "must be less than 5.1 characters",
		},

		{
			error:           NewMaxValidationError(ve.TypeSlice, 5, false),
			expectedMessage: "must have less than 5 items",
		},
		{
			error:           NewMaxValidationError(ve.TypeSlice, 5.1, false),
			expectedMessage: "must have less than 5.1 items",
		},

		{
			error:           NewMaxValidationError(ve.TypeArray, 5, false),
			expectedMessage: "must have less than 5 items",
		},
		{
			error:           NewMaxValidationError(ve.TypeArray, 5.1, false),
			expectedMessage: "must have less than 5.1 items",
		},

		{
			error:           NewMaxValidationError(ve.TypeMap, 5, false),
			expectedMessage: "must have less than 5 items",
		},
		{
			error:           NewMaxValidationError(ve.TypeMap, 5.1, false),
			expectedMessage: "must have less than 5.1 items",
		},

		{
			error:           NewMaxValidationError(ve.TypeNumber, 5, false),
			expectedMessage: "must be less than 5",
		},
		{
			error:           NewMaxValidationError(ve.TypeNumber, 5.1, false),
			expectedMessage: "must be less than 5.1",
		},

		{
			error:           NewMaxValidationError(ve.TypeInvalid, fakerInstance.Int(), false),
			expectedMessage: "max cannot be determined",
		},
	} {
		t.Run(fmt.Sprintf("#%d", ttIdx), func(t *testing.T) {
			// then
			require.EqualError(t, tt.error, tt.expectedMessage)
		})
	}
}

func BenchmarkMaxExclusiveRule(b *testing.B) {
	runRuleBenchmarks(b, maxExclusiveRuleDataProvider)
}

func maxExclusiveRuleDataProvider() map[string]*ruleTestCaseData {
	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             MaxExclusive(3),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string of length 2": {
			rule:             MaxExclusive(3),
			value:            "ab",
			expectedNewValue: "ab",
			expectedError:    nil,
		},
		"string of length 3": {
			rule:             MaxExclusive(3),
			value:            "abc",
			expectedNewValue: "abc",
			expectedError:    NewMaxValidationError(ve.TypeString, 3, false),
		},
		"*string of length 2": {
			rule:             MaxExclusive(3),
			value:            ptr("ab"),
			expectedNewValue: ptr("ab"),
			expectedError:    nil,
		},
		"*string of length 3": {
			rule:             MaxExclusive(3),
			value:            ptr("abc"),
			expectedNewValue: ptr("abc"),
			expectedError:    NewMaxValidationError(ve.TypeString, 3, false),
		},

		"int(2)": {
			rule:             MaxExclusive(3),
			value:            2,
			expectedNewValue: 2,
			expectedError:    nil,
		},
		"int(3)": {
			rule:             MaxExclusive(3),
			value:            3,
			expectedNewValue: 3,
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*int(2)": {
			rule:             MaxExclusive(3),
			value:            ptr(2),
			expectedNewValue: ptr(2),
			expectedError:    nil,
		},
		"*int(3)": {
			rule:             MaxExclusive(3),
			value:            ptr(3),
			expectedNewValue: ptr(3),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"int8(2)": {
			rule:             MaxExclusive(3),
			value:            int8(2),
			expectedNewValue: int8(2),
			expectedError:    nil,
		},
		"int8(3)": {
			rule:             MaxExclusive(3),
			value:            int8(3),
			expectedNewValue: int8(3),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*int8(2)": {
			rule:             MaxExclusive(3),
			value:            ptr(int8(2)),
			expectedNewValue: ptr(int8(2)),
			expectedError:    nil,
		},
		"*int8(3)": {
			rule:             MaxExclusive(3),
			value:            ptr(int8(3)),
			expectedNewValue: ptr(int8(3)),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"int16(2)": {
			rule:             MaxExclusive(3),
			value:            int16(2),
			expectedNewValue: int16(2),
			expectedError:    nil,
		},
		"int16(3)": {
			rule:             MaxExclusive(3),
			value:            int16(3),
			expectedNewValue: int16(3),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*int16(2)": {
			rule:             MaxExclusive(3),
			value:            ptr(int16(2)),
			expectedNewValue: ptr(int16(2)),
			expectedError:    nil,
		},
		"*int16(3)": {
			rule:             MaxExclusive(3),
			value:            ptr(int16(3)),
			expectedNewValue: ptr(int16(3)),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"int32(2)": {
			rule:             MaxExclusive(3),
			value:            int32(2),
			expectedNewValue: int32(2),
			expectedError:    nil,
		},
		"int32(3)": {
			rule:             MaxExclusive(3),
			value:            int32(3),
			expectedNewValue: int32(3),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*int32(2)": {
			rule:             MaxExclusive(3),
			value:            ptr(int32(2)),
			expectedNewValue: ptr(int32(2)),
			expectedError:    nil,
		},
		"*int32(3)": {
			rule:             MaxExclusive(3),
			value:            ptr(int32(3)),
			expectedNewValue: ptr(int32(3)),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"int64(2)": {
			rule:             MaxExclusive(3),
			value:            int64(2),
			expectedNewValue: int64(2),
			expectedError:    nil,
		},
		"int64(3)": {
			rule:             MaxExclusive(3),
			value:            int64(3),
			expectedNewValue: int64(3),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*int64(2)": {
			rule:             MaxExclusive(3),
			value:            ptr(int64(2)),
			expectedNewValue: ptr(int64(2)),
			expectedError:    nil,
		},
		"*int64(3)": {
			rule:             MaxExclusive(3),
			value:            ptr(int64(3)),
			expectedNewValue: ptr(int64(3)),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"uint(2)": {
			rule:             MaxExclusive(3),
			value:            uint(2),
			expectedNewValue: uint(2),
			expectedError:    nil,
		},
		"uint(3)": {
			rule:             MaxExclusive(3),
			value:            uint(3),
			expectedNewValue: uint(3),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*uint(2)": {
			rule:             MaxExclusive(3),
			value:            ptr(uint(2)),
			expectedNewValue: ptr(uint(2)),
			expectedError:    nil,
		},
		"*uint(3)": {
			rule:             MaxExclusive(3),
			value:            ptr(uint(3)),
			expectedNewValue: ptr(uint(3)),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"uint8(2)": {
			rule:             MaxExclusive(3),
			value:            uint8(2),
			expectedNewValue: uint8(2),
			expectedError:    nil,
		},
		"uint8(3)": {
			rule:             MaxExclusive(3),
			value:            uint8(3),
			expectedNewValue: uint8(3),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*uint8(2)": {
			rule:             MaxExclusive(3),
			value:            ptr(uint8(2)),
			expectedNewValue: ptr(uint8(2)),
			expectedError:    nil,
		},
		"*uint8(3)": {
			rule:             MaxExclusive(3),
			value:            ptr(uint8(3)),
			expectedNewValue: ptr(uint8(3)),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"uint16(2)": {
			rule:             MaxExclusive(3),
			value:            uint16(2),
			expectedNewValue: uint16(2),
			expectedError:    nil,
		},
		"uint16(3)": {
			rule:             MaxExclusive(3),
			value:            uint16(3),
			expectedNewValue: uint16(3),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*uint16(2)": {
			rule:             MaxExclusive(3),
			value:            ptr(uint16(2)),
			expectedNewValue: ptr(uint16(2)),
			expectedError:    nil,
		},
		"*uint16(3)": {
			rule:             MaxExclusive(3),
			value:            ptr(uint16(3)),
			expectedNewValue: ptr(uint16(3)),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"uint32(2)": {
			rule:             MaxExclusive(3),
			value:            uint32(2),
			expectedNewValue: uint32(2),
			expectedError:    nil,
		},
		"uint32(3)": {
			rule:             MaxExclusive(3),
			value:            uint32(3),
			expectedNewValue: uint32(3),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*uint32(2)": {
			rule:             MaxExclusive(3),
			value:            ptr(uint32(2)),
			expectedNewValue: ptr(uint32(2)),
			expectedError:    nil,
		},
		"*uint32(3)": {
			rule:             MaxExclusive(3),
			value:            ptr(uint32(3)),
			expectedNewValue: ptr(uint32(3)),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"uint64(2)": {
			rule:             MaxExclusive(3),
			value:            uint64(2),
			expectedNewValue: uint64(2),
			expectedError:    nil,
		},
		"uint64(3)": {
			rule:             MaxExclusive(3),
			value:            uint64(3),
			expectedNewValue: uint64(3),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*uint64(2)": {
			rule:             MaxExclusive(3),
			value:            ptr(uint64(2)),
			expectedNewValue: ptr(uint64(2)),
			expectedError:    nil,
		},
		"*uint64(3)": {
			rule:             MaxExclusive(3),
			value:            ptr(uint64(3)),
			expectedNewValue: ptr(uint64(3)),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"float32(3.0)": {
			rule:             MaxExclusive(3),
			value:            float32(2.99),
			expectedNewValue: float32(2.99),
			expectedError:    nil,
		},
		"float32(3.01)": {
			rule:             MaxExclusive(3),
			value:            float32(3.0),
			expectedNewValue: float32(3.0),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*float32(3.0)": {
			rule:             MaxExclusive(3),
			value:            float32(2.99),
			expectedNewValue: float32(2.99),
			expectedError:    nil,
		},
		"*float32(3.01)": {
			rule:             MaxExclusive(3),
			value:            float32(3.0),
			expectedNewValue: float32(3.0),
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"float64(3.0)": {
			rule:             MaxExclusive(3),
			value:            2.99,
			expectedNewValue: 2.99,
			expectedError:    nil,
		},
		"float64(3.01)": {
			rule:             MaxExclusive(3),
			value:            3.0,
			expectedNewValue: 3.0,
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},
		"*float64(3.0)": {
			rule:             MaxExclusive(3),
			value:            2.99,
			expectedNewValue: 2.99,
			expectedError:    nil,
		},
		"*float64(3.01)": {
			rule:             MaxExclusive(3),
			value:            3.0,
			expectedNewValue: 3.0,
			expectedError:    NewMaxValidationError(ve.TypeNumber, 3, false),
		},

		"slice with 2 items": {
			rule:             MaxExclusive(3),
			value:            []int{1, 2},
			expectedNewValue: []int{1, 2},
			expectedError:    nil,
		},
		"slice with 3 items": {
			rule:             MaxExclusive(3),
			value:            []int{1, 2, 3},
			expectedNewValue: []int{1, 2, 3},
			expectedError:    NewMaxValidationError(ve.TypeSlice, 3, false),
		},

		"array with 2 items": {
			rule:             MaxExclusive(3),
			value:            [2]int{1, 2},
			expectedNewValue: [2]int{1, 2},
			expectedError:    nil,
		},
		"array with 3 items": {
			rule:             MaxExclusive(3),
			value:            [3]int{1, 2, 3},
			expectedNewValue: [3]int{1, 2, 3},
			expectedError:    NewMaxValidationError(ve.TypeArray, 3, false),
		},

		"map with 2 keys": {
			rule:             MaxExclusive(3),
			value:            map[any]int{1: 1, "a": 2},
			expectedNewValue: map[any]int{1: 1, "a": 2},
			expectedError:    nil,
		},
		"map with 3 keys": {
			rule:             MaxExclusive(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedError:    NewMaxValidationError(ve.TypeMap, 3, false),
		},

		// unsupported values
		"complex": {
			rule:             MaxExclusive(3),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewMaxValidationError(ve.TypeInvalid, 3, false),
		},
		"bool": {
			rule:             MaxExclusive(3),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewMaxValidationError(ve.TypeInvalid, 3, false),
		},
		"struct": {
			rule:             MaxExclusive(3),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewMaxValidationError(ve.TypeInvalid, 3, false),
		},
	}
}
