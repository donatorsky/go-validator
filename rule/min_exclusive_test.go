package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_MinExclusiveRule(t *testing.T) {
	runRuleTestCases(t, minExclusiveRuleDataProvider)
}

func Test_MinExclusiveValidationError(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		error           error
		expectedMessage string
	}{
		{
			error:           NewMinValidationError(ve.TypeString, 5, false),
			expectedMessage: "must be more than 5 characters",
		},
		{
			error:           NewMinValidationError(ve.TypeString, 5.1, false),
			expectedMessage: "must be more than 5.1 characters",
		},

		{
			error:           NewMinValidationError(ve.TypeSlice, 5, false),
			expectedMessage: "must have more than 5 items",
		},
		{
			error:           NewMinValidationError(ve.TypeSlice, 5.1, false),
			expectedMessage: "must have more than 5.1 items",
		},

		{
			error:           NewMinValidationError(ve.TypeArray, 5, false),
			expectedMessage: "must have more than 5 items",
		},
		{
			error:           NewMinValidationError(ve.TypeArray, 5.1, false),
			expectedMessage: "must have more than 5.1 items",
		},

		{
			error:           NewMinValidationError(ve.TypeMap, 5, false),
			expectedMessage: "must have more than 5 items",
		},
		{
			error:           NewMinValidationError(ve.TypeMap, 5.1, false),
			expectedMessage: "must have more than 5.1 items",
		},

		{
			error:           NewMinValidationError(ve.TypeNumber, 5, false),
			expectedMessage: "must be greater than 5",
		},
		{
			error:           NewMinValidationError(ve.TypeNumber, 5.1, false),
			expectedMessage: "must be greater than 5.1",
		},

		{
			error:           NewMinValidationError(ve.TypeInvalid, fakerInstance.Int(), false),
			expectedMessage: "min cannot be determined",
		},
	} {
		t.Run(fmt.Sprintf("#%d", ttIdx), func(t *testing.T) {
			// then
			require.EqualError(t, tt.error, tt.expectedMessage)
		})
	}
}

func BenchmarkMinExclusiveRule(b *testing.B) {
	runRuleBenchmarks(b, minExclusiveRuleDataProvider)
}

func minExclusiveRuleDataProvider() map[string]*ruleTestCaseData {
	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             MinExclusive(3),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string of length 3": {
			rule:             MinExclusive(3),
			value:            "abc",
			expectedNewValue: "abc",
			expectedError:    NewMinValidationError(ve.TypeString, 3, false),
		},
		"string of length 4": {
			rule:             MinExclusive(3),
			value:            "abcd",
			expectedNewValue: "abcd",
			expectedError:    nil,
		},
		"*string of length 3": {
			rule:             MinExclusive(3),
			value:            ptr("abc"),
			expectedNewValue: ptr("abc"),
			expectedError:    NewMinValidationError(ve.TypeString, 3, false),
		},
		"*string of length 4": {
			rule:             MinExclusive(3),
			value:            ptr("abcd"),
			expectedNewValue: ptr("abcd"),
			expectedError:    nil,
		},

		"int(3)": {
			rule:             MinExclusive(3),
			value:            3,
			expectedNewValue: 3,
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"int(4)": {
			rule:             MinExclusive(3),
			value:            4,
			expectedNewValue: 4,
			expectedError:    nil,
		},
		"*int(3)": {
			rule:             MinExclusive(3),
			value:            ptr(3),
			expectedNewValue: ptr(3),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*int(4)": {
			rule:             MinExclusive(3),
			value:            ptr(4),
			expectedNewValue: ptr(4),
			expectedError:    nil,
		},

		"int8(3)": {
			rule:             MinExclusive(3),
			value:            int8(3),
			expectedNewValue: int8(3),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"int8(4)": {
			rule:             MinExclusive(3),
			value:            int8(4),
			expectedNewValue: int8(4),
			expectedError:    nil,
		},
		"*int8(3)": {
			rule:             MinExclusive(3),
			value:            ptr(int8(3)),
			expectedNewValue: ptr(int8(3)),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*int8(4)": {
			rule:             MinExclusive(3),
			value:            ptr(int8(4)),
			expectedNewValue: ptr(int8(4)),
			expectedError:    nil,
		},

		"int16(3)": {
			rule:             MinExclusive(3),
			value:            int16(3),
			expectedNewValue: int16(3),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"int16(4)": {
			rule:             MinExclusive(3),
			value:            int16(4),
			expectedNewValue: int16(4),
			expectedError:    nil,
		},
		"*int16(3)": {
			rule:             MinExclusive(3),
			value:            ptr(int16(3)),
			expectedNewValue: ptr(int16(3)),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*int16(4)": {
			rule:             MinExclusive(3),
			value:            ptr(int16(4)),
			expectedNewValue: ptr(int16(4)),
			expectedError:    nil,
		},

		"int32(3)": {
			rule:             MinExclusive(3),
			value:            int32(3),
			expectedNewValue: int32(3),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"int32(4)": {
			rule:             MinExclusive(3),
			value:            int32(4),
			expectedNewValue: int32(4),
			expectedError:    nil,
		},
		"*int32(3)": {
			rule:             MinExclusive(3),
			value:            ptr(int32(3)),
			expectedNewValue: ptr(int32(3)),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*int32(4)": {
			rule:             MinExclusive(3),
			value:            ptr(int32(4)),
			expectedNewValue: ptr(int32(4)),
			expectedError:    nil,
		},

		"int64(3)": {
			rule:             MinExclusive(3),
			value:            int64(3),
			expectedNewValue: int64(3),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"int64(4)": {
			rule:             MinExclusive(3),
			value:            int64(4),
			expectedNewValue: int64(4),
			expectedError:    nil,
		},
		"*int64(3)": {
			rule:             MinExclusive(3),
			value:            ptr(int64(3)),
			expectedNewValue: ptr(int64(3)),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*int64(4)": {
			rule:             MinExclusive(3),
			value:            ptr(int64(4)),
			expectedNewValue: ptr(int64(4)),
			expectedError:    nil,
		},

		"uint(3)": {
			rule:             MinExclusive(3),
			value:            uint(3),
			expectedNewValue: uint(3),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"uint(4)": {
			rule:             MinExclusive(3),
			value:            uint(4),
			expectedNewValue: uint(4),
			expectedError:    nil,
		},
		"*uint(3)": {
			rule:             MinExclusive(3),
			value:            ptr(uint(3)),
			expectedNewValue: ptr(uint(3)),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*uint(4)": {
			rule:             MinExclusive(3),
			value:            ptr(uint(4)),
			expectedNewValue: ptr(uint(4)),
			expectedError:    nil,
		},

		"uint8(3)": {
			rule:             MinExclusive(3),
			value:            uint8(3),
			expectedNewValue: uint8(3),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"uint8(4)": {
			rule:             MinExclusive(3),
			value:            uint8(4),
			expectedNewValue: uint8(4),
			expectedError:    nil,
		},
		"*uint8(3)": {
			rule:             MinExclusive(3),
			value:            ptr(uint8(3)),
			expectedNewValue: ptr(uint8(3)),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*uint8(4)": {
			rule:             MinExclusive(3),
			value:            ptr(uint8(4)),
			expectedNewValue: ptr(uint8(4)),
			expectedError:    nil,
		},

		"uint16(3)": {
			rule:             MinExclusive(3),
			value:            uint16(3),
			expectedNewValue: uint16(3),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"uint16(4)": {
			rule:             MinExclusive(3),
			value:            uint16(4),
			expectedNewValue: uint16(4),
			expectedError:    nil,
		},
		"*uint16(3)": {
			rule:             MinExclusive(3),
			value:            ptr(uint16(3)),
			expectedNewValue: ptr(uint16(3)),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*uint16(4)": {
			rule:             MinExclusive(3),
			value:            ptr(uint16(4)),
			expectedNewValue: ptr(uint16(4)),
			expectedError:    nil,
		},

		"uint32(3)": {
			rule:             MinExclusive(3),
			value:            uint32(3),
			expectedNewValue: uint32(3),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"uint32(4)": {
			rule:             MinExclusive(3),
			value:            uint32(4),
			expectedNewValue: uint32(4),
			expectedError:    nil,
		},
		"*uint32(3)": {
			rule:             MinExclusive(3),
			value:            ptr(uint32(3)),
			expectedNewValue: ptr(uint32(3)),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*uint32(4)": {
			rule:             MinExclusive(3),
			value:            ptr(uint32(4)),
			expectedNewValue: ptr(uint32(4)),
			expectedError:    nil,
		},

		"uint64(3)": {
			rule:             MinExclusive(3),
			value:            uint64(3),
			expectedNewValue: uint64(3),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"uint64(4)": {
			rule:             MinExclusive(3),
			value:            uint64(4),
			expectedNewValue: uint64(4),
			expectedError:    nil,
		},
		"*uint64(3)": {
			rule:             MinExclusive(3),
			value:            ptr(uint64(3)),
			expectedNewValue: ptr(uint64(3)),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*uint64(4)": {
			rule:             MinExclusive(3),
			value:            ptr(uint64(4)),
			expectedNewValue: ptr(uint64(4)),
			expectedError:    nil,
		},

		"float32(3.0)": {
			rule:             MinExclusive(3),
			value:            float32(3.0),
			expectedNewValue: float32(3.0),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"float32(3.01)": {
			rule:             MinExclusive(3),
			value:            float32(3.01),
			expectedNewValue: float32(3.01),
			expectedError:    nil,
		},
		"*float32(3.0)": {
			rule:             MinExclusive(3),
			value:            ptr(float32(3.0)),
			expectedNewValue: ptr(float32(3.0)),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*float32(3.01)": {
			rule:             MinExclusive(3),
			value:            ptr(float32(3.01)),
			expectedNewValue: ptr(float32(3.01)),
			expectedError:    nil,
		},

		"float64(3.0)": {
			rule:             MinExclusive(3),
			value:            3.0,
			expectedNewValue: 3.0,
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"float64(3.01)": {
			rule:             MinExclusive(3),
			value:            3.01,
			expectedNewValue: 3.01,
			expectedError:    nil,
		},
		"*float64(3.0)": {
			rule:             MinExclusive(3),
			value:            ptr(3.0),
			expectedNewValue: ptr(3.0),
			expectedError:    NewMinValidationError(ve.TypeNumber, 3, false),
		},
		"*float64(3.01)": {
			rule:             MinExclusive(3),
			value:            ptr(3.01),
			expectedNewValue: ptr(3.01),
			expectedError:    nil,
		},

		"slice with 3 items": {
			rule:             MinExclusive(3),
			value:            []int{1, 2, 3},
			expectedNewValue: []int{1, 2, 3},
			expectedError:    NewMinValidationError(ve.TypeSlice, 3, false),
		},
		"slice with 4 items": {
			rule:             MinExclusive(3),
			value:            []int{1, 2, 3, 4},
			expectedNewValue: []int{1, 2, 3, 4},
			expectedError:    nil,
		},

		"array with 3 items": {
			rule:             MinExclusive(3),
			value:            [3]int{1, 2, 3},
			expectedNewValue: [3]int{1, 2, 3},
			expectedError:    NewMinValidationError(ve.TypeArray, 3, false),
		},
		"array with 4 items": {
			rule:             MinExclusive(3),
			value:            [4]int{1, 2, 3, 4},
			expectedNewValue: [4]int{1, 2, 3, 4},
			expectedError:    nil,
		},

		"map with 3 keys": {
			rule:             MinExclusive(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedError:    NewMinValidationError(ve.TypeMap, 3, false),
		},
		"map with 4 keys": {
			rule:             MinExclusive(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedError:    nil,
		},

		// unsupported values
		"complex": {
			rule:             MinExclusive(3),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewMinValidationError(ve.TypeInvalid, 3, false),
		},
		"bool": {
			rule:             MinExclusive(3),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewMinValidationError(ve.TypeInvalid, 3, false),
		},
		"struct": {
			rule:             MinExclusive(3),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewMinValidationError(ve.TypeInvalid, 3, false),
		},
	}
}
