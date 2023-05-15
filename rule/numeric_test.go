package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NumericRule(t *testing.T) {
	runRuleTestCases(t, NumericRuleDataProvider)
}

func Test_NumericValidationError(t *testing.T) {
	// when
	err := NewNumericValidationError()

	// then
	require.EqualError(t, err, "must be a number")
}

func BenchmarkNumericRule(b *testing.B) {
	runRuleBenchmarks(b, NumericRuleDataProvider)
}

func NumericRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		intDummy             = fakerInstance.Int()
		int8Dummy            = fakerInstance.Int8()
		int16Dummy           = fakerInstance.Int16()
		int32Dummy           = fakerInstance.Int32()
		int64Dummy           = fakerInstance.Int64()
		uintDummy            = fakerInstance.UInt()
		uint8Dummy           = fakerInstance.UInt8()
		uint16Dummy          = fakerInstance.UInt16()
		uint32Dummy          = fakerInstance.UInt32()
		uint64Dummy          = fakerInstance.UInt64()
		float32Dummy         = fakerInstance.Float32(5, -1000, 1000)
		float64Dummy         = fakerInstance.Float64(5, -1000, 1000)
		complex64Dummy       = complex(fakerInstance.Float32(5, -1000, 1000), fakerInstance.Float32(5, -1000, 1000))
		complex128Dummy      = complex(fakerInstance.Float64(5, -1000, 1000), fakerInstance.Float64(5, -1000, 1000))
		complexAsStringDummy = "1+2i"
		floatAsStringDummy   = "1.23"
		uintAsStringDummy    = fmt.Sprintf("%d", uint64(1)<<63+1)
		intAsStringDummy     = "-123"
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Numeric(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"int": {
			rule:             Numeric(),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    nil,
		},
		"*int": {
			rule:             Numeric(),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    nil,
		},

		"int8": {
			rule:             Numeric(),
			value:            int8Dummy,
			expectedNewValue: int8Dummy,
			expectedError:    nil,
		},
		"*int8": {
			rule:             Numeric(),
			value:            &int8Dummy,
			expectedNewValue: &int8Dummy,
			expectedError:    nil,
		},

		"int16": {
			rule:             Numeric(),
			value:            int16Dummy,
			expectedNewValue: int16Dummy,
			expectedError:    nil,
		},
		"*int16": {
			rule:             Numeric(),
			value:            &int16Dummy,
			expectedNewValue: &int16Dummy,
			expectedError:    nil,
		},

		"int32": {
			rule:             Numeric(),
			value:            int32Dummy,
			expectedNewValue: int32Dummy,
			expectedError:    nil,
		},
		"*int32": {
			rule:             Numeric(),
			value:            &int32Dummy,
			expectedNewValue: &int32Dummy,
			expectedError:    nil,
		},

		"int64": {
			rule:             Numeric(),
			value:            int64Dummy,
			expectedNewValue: int64Dummy,
			expectedError:    nil,
		},
		"*int64": {
			rule:             Numeric(),
			value:            &int64Dummy,
			expectedNewValue: &int64Dummy,
			expectedError:    nil,
		},

		"uint": {
			rule:             Numeric(),
			value:            uintDummy,
			expectedNewValue: uintDummy,
			expectedError:    nil,
		},
		"*uint": {
			rule:             Numeric(),
			value:            &uintDummy,
			expectedNewValue: &uintDummy,
			expectedError:    nil,
		},

		"uint8": {
			rule:             Numeric(),
			value:            uint8Dummy,
			expectedNewValue: uint8Dummy,
			expectedError:    nil,
		},
		"*uint8": {
			rule:             Numeric(),
			value:            &uint8Dummy,
			expectedNewValue: &uint8Dummy,
			expectedError:    nil,
		},

		"uint16": {
			rule:             Numeric(),
			value:            uint16Dummy,
			expectedNewValue: uint16Dummy,
			expectedError:    nil,
		},
		"*uint16": {
			rule:             Numeric(),
			value:            &uint16Dummy,
			expectedNewValue: &uint16Dummy,
			expectedError:    nil,
		},

		"uint32": {
			rule:             Numeric(),
			value:            uint32Dummy,
			expectedNewValue: uint32Dummy,
			expectedError:    nil,
		},
		"*uint32": {
			rule:             Numeric(),
			value:            &uint32Dummy,
			expectedNewValue: &uint32Dummy,
			expectedError:    nil,
		},

		"uint64": {
			rule:             Numeric(),
			value:            uint64Dummy,
			expectedNewValue: uint64Dummy,
			expectedError:    nil,
		},
		"*uint64": {
			rule:             Numeric(),
			value:            &uint64Dummy,
			expectedNewValue: &uint64Dummy,
			expectedError:    nil,
		},

		"float32": {
			rule:             Numeric(),
			value:            float32Dummy,
			expectedNewValue: float32Dummy,
			expectedError:    nil,
		},
		"*float32": {
			rule:             Numeric(),
			value:            &float32Dummy,
			expectedNewValue: &float32Dummy,
			expectedError:    nil,
		},

		"float64": {
			rule:             Numeric(),
			value:            float64Dummy,
			expectedNewValue: float64Dummy,
			expectedError:    nil,
		},
		"*float64": {
			rule:             Numeric(),
			value:            &float64Dummy,
			expectedNewValue: &float64Dummy,
			expectedError:    nil,
		},

		"complex64": {
			rule:             Numeric(),
			value:            complex64Dummy,
			expectedNewValue: complex64Dummy,
			expectedError:    nil,
		},
		"*complex64": {
			rule:             Numeric(),
			value:            &complex64Dummy,
			expectedNewValue: &complex64Dummy,
			expectedError:    nil,
		},

		"complex128": {
			rule:             Numeric(),
			value:            complex128Dummy,
			expectedNewValue: complex128Dummy,
			expectedError:    nil,
		},
		"*complex128": {
			rule:             Numeric(),
			value:            &complex128Dummy,
			expectedNewValue: &complex128Dummy,
			expectedError:    nil,
		},

		"complex as string": {
			rule:             Numeric(),
			value:            complexAsStringDummy,
			expectedNewValue: 1 + 2i,
			expectedError:    nil,
		},
		"complex as *string": {
			rule:             Numeric(),
			value:            &complexAsStringDummy,
			expectedNewValue: 1 + 2i,
			expectedError:    nil,
		},

		"float as string": {
			rule:             Numeric(),
			value:            floatAsStringDummy,
			expectedNewValue: 1.23,
			expectedError:    nil,
		},
		"float as *string": {
			rule:             Numeric(),
			value:            &floatAsStringDummy,
			expectedNewValue: 1.23,
			expectedError:    nil,
		},

		"uint as string": {
			rule:             Numeric(),
			value:            uintAsStringDummy,
			expectedNewValue: uint64(1)<<63 + 1,
			expectedError:    nil,
		},
		"uint as *string": {
			rule:             Numeric(),
			value:            &uintAsStringDummy,
			expectedNewValue: uint64(1)<<63 + 1,
			expectedError:    nil,
		},

		"int as string": {
			rule:             Numeric(),
			value:            intAsStringDummy,
			expectedNewValue: int64(-123),
			expectedError:    nil,
		},
		"int as *string": {
			rule:             Numeric(),
			value:            &intAsStringDummy,
			expectedNewValue: int64(-123),
			expectedError:    nil,
		},

		// unsupported values
		"bool": {
			rule:             Numeric(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewNumericValidationError(),
		},
		"slice": {
			rule:             Numeric(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewNumericValidationError(),
		},
		"array": {
			rule:             Numeric(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewNumericValidationError(),
		},
		"map": {
			rule:             Numeric(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewNumericValidationError(),
		},
		"struct": {
			rule:             Numeric(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewNumericValidationError(),
		},
	}
}
