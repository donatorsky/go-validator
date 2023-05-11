package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BooleanRule(t *testing.T) {
	runRuleTestCases(t, booleanRuleDataProvider)
}

func Test_BooleanValidationError(t *testing.T) {
	// when
	err := NewBooleanValidationError()

	// then
	require.EqualError(t, err, "must be true or false")
}

func BenchmarkBooleanRule(b *testing.B) {
	runRuleBenchmarks(b, booleanRuleDataProvider)
}

func booleanRuleDataProvider() map[string]*ruleTestCaseData {
	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Boolean(),
			value:            nil,
			expectedNewValue: (*bool)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"false": {
			rule:             Boolean(),
			value:            false,
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to false": {
			rule:             Boolean(),
			value:            ptr(false),
			expectedNewValue: ptr(false),
			expectedError:    nil,
			expectedToBail:   false,
		},
		"true": {
			rule:             Boolean(),
			value:            true,
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to true": {
			rule:             Boolean(),
			value:            ptr(true),
			expectedNewValue: ptr(true),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"string(F)": {
			rule:             Boolean(),
			value:            "F",
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*string(F)": {
			rule:             Boolean(),
			value:            ptr("F"),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"string(T)": {
			rule:             Boolean(),
			value:            "T",
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*string(T)": {
			rule:             Boolean(),
			value:            ptr("T"),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"int(0)": {
			rule:             Boolean(),
			value:            0,
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int(0)": {
			rule:             Boolean(),
			value:            ptr(0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"int(1)": {
			rule:             Boolean(),
			value:            1,
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int(1)": {
			rule:             Boolean(),
			value:            ptr(1),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"int(2)": {
			rule:             Boolean(),
			value:            2,
			expectedNewValue: 2,
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"int8(0)": {
			rule:             Boolean(),
			value:            int8(0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int8(0)": {
			rule:             Boolean(),
			value:            ptr(int8(0)),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"int8(1)": {
			rule:             Boolean(),
			value:            int8(1),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int8(1)": {
			rule:             Boolean(),
			value:            ptr(int8(1)),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"int8(2)": {
			rule:             Boolean(),
			value:            int8(2),
			expectedNewValue: int8(2),
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"int16(0)": {
			rule:             Boolean(),
			value:            int16(0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int16(0)": {
			rule:             Boolean(),
			value:            ptr(int16(0)),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"int16(1)": {
			rule:             Boolean(),
			value:            int16(1),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int16(1)": {
			rule:             Boolean(),
			value:            ptr(int16(1)),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"int16(2)": {
			rule:             Boolean(),
			value:            int16(2),
			expectedNewValue: int16(2),
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"int32(0)": {
			rule:             Boolean(),
			value:            int32(0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int32(0)": {
			rule:             Boolean(),
			value:            ptr(int32(0)),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"int32(1)": {
			rule:             Boolean(),
			value:            int32(1),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int32(1)": {
			rule:             Boolean(),
			value:            ptr(int32(1)),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"int32(2)": {
			rule:             Boolean(),
			value:            int32(2),
			expectedNewValue: int32(2),
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"int64(0)": {
			rule:             Boolean(),
			value:            int64(0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int64(0)": {
			rule:             Boolean(),
			value:            ptr(int64(0)),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"int64(1)": {
			rule:             Boolean(),
			value:            int64(1),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int64(1)": {
			rule:             Boolean(),
			value:            ptr(int64(1)),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"int64(2)": {
			rule:             Boolean(),
			value:            int64(2),
			expectedNewValue: int64(2),
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"uint(0)": {
			rule:             Boolean(),
			value:            uint(0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint(0)": {
			rule:             Boolean(),
			value:            ptr(uint(0)),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"uint(1)": {
			rule:             Boolean(),
			value:            uint(1),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint(1)": {
			rule:             Boolean(),
			value:            ptr(uint(1)),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"uint(2)": {
			rule:             Boolean(),
			value:            uint(2),
			expectedNewValue: uint(2),
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"uint8(0)": {
			rule:             Boolean(),
			value:            uint8(0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint8(0)": {
			rule:             Boolean(),
			value:            ptr(uint8(0)),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"uint8(1)": {
			rule:             Boolean(),
			value:            uint8(1),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint8(1)": {
			rule:             Boolean(),
			value:            ptr(uint8(1)),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"uint8(2)": {
			rule:             Boolean(),
			value:            uint8(2),
			expectedNewValue: uint8(2),
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"uint16(0)": {
			rule:             Boolean(),
			value:            uint16(0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint16(0)": {
			rule:             Boolean(),
			value:            ptr(uint16(0)),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"uint16(1)": {
			rule:             Boolean(),
			value:            uint16(1),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint16(1)": {
			rule:             Boolean(),
			value:            ptr(uint16(1)),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"uint16(2)": {
			rule:             Boolean(),
			value:            uint16(2),
			expectedNewValue: uint16(2),
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"uint32(0)": {
			rule:             Boolean(),
			value:            uint32(0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint32(0)": {
			rule:             Boolean(),
			value:            ptr(uint32(0)),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"uint32(1)": {
			rule:             Boolean(),
			value:            uint32(1),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint32(1)": {
			rule:             Boolean(),
			value:            ptr(uint32(1)),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"uint32(2)": {
			rule:             Boolean(),
			value:            uint32(2),
			expectedNewValue: uint32(2),
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"uint64(0)": {
			rule:             Boolean(),
			value:            uint64(0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint64(0)": {
			rule:             Boolean(),
			value:            ptr(uint64(0)),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"uint64(1)": {
			rule:             Boolean(),
			value:            uint64(1),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint64(1)": {
			rule:             Boolean(),
			value:            ptr(uint64(1)),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"uint64(2)": {
			rule:             Boolean(),
			value:            uint64(2),
			expectedNewValue: uint64(2),
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"float32(0.0)": {
			rule:             Boolean(),
			value:            float32(0.0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*float32(0.0)": {
			rule:             Boolean(),
			value:            ptr(float32(0.0)),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"float32(1.0)": {
			rule:             Boolean(),
			value:            float32(1.0),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*float32(1.0)": {
			rule:             Boolean(),
			value:            ptr(float32(1.0)),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"float32(2.0)": {
			rule:             Boolean(),
			value:            float32(2.0),
			expectedNewValue: float32(2.0),
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		"float64(0.0)": {
			rule:             Boolean(),
			value:            0.0,
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*float64(0.0)": {
			rule:             Boolean(),
			value:            ptr(0.0),
			expectedNewValue: false,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"float64(1.0)": {
			rule:             Boolean(),
			value:            1.0,
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*float64(1.0)": {
			rule:             Boolean(),
			value:            ptr(1.0),
			expectedNewValue: true,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"float64(2.0)": {
			rule:             Boolean(),
			value:            2.0,
			expectedNewValue: 2.0,
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},

		// unsupported values
		"complex": {
			rule:             Boolean(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},
		"non-boolean string value": {
			rule:             Boolean(),
			value:            "foo",
			expectedNewValue: "foo",
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},
		"slice": {
			rule:             Boolean(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},
		"array": {
			rule:             Boolean(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},
		"map": {
			rule:             Boolean(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},
		"struct": {
			rule:             Boolean(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewBooleanValidationError(),
			expectedToBail:   true,
		},
	}
}
