package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IntegerRule(t *testing.T) {
	runRuleTestCases(t, integerRuleDataProvider)
}

func Test_IntegerValidationError(t *testing.T) {
	// given
	var (
		expectedTypeDummy = fakerInstance.Lorem().Word()
		actualTypeDummy   = fakerInstance.Lorem().Word()
	)

	// when
	err := NewIntegerValidationError(expectedTypeDummy, actualTypeDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must be an %s but is %s",
		expectedTypeDummy,
		actualTypeDummy,
	))
}

func BenchmarkIntegerRule(b *testing.B) {
	runRuleBenchmarks(b, integerRuleDataProvider)
}

func integerRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		intDummy    = fakerInstance.Int()
		int8Dummy   = fakerInstance.Int8()
		int16Dummy  = fakerInstance.Int16()
		int32Dummy  = fakerInstance.Int32()
		int64Dummy  = fakerInstance.Int64()
		uintDummy   = fakerInstance.UInt()
		uint8Dummy  = fakerInstance.UInt8()
		uint16Dummy = fakerInstance.UInt16()
		uint32Dummy = fakerInstance.UInt32()
		uint64Dummy = fakerInstance.UInt64()
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Integer[int](),
			value:            nil,
			expectedNewValue: (*int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"int": {
			rule:             Integer[int](),
			value:            intDummy,
			expectedNewValue: intDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int": {
			rule:             Integer[int](),
			value:            &intDummy,
			expectedNewValue: &intDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"int8": {
			rule:             Integer[int8](),
			value:            int8Dummy,
			expectedNewValue: int8Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int8": {
			rule:             Integer[int8](),
			value:            &int8Dummy,
			expectedNewValue: &int8Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"int16": {
			rule:             Integer[int16](),
			value:            int16Dummy,
			expectedNewValue: int16Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int16": {
			rule:             Integer[int16](),
			value:            &int16Dummy,
			expectedNewValue: &int16Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"int32": {
			rule:             Integer[int32](),
			value:            int32Dummy,
			expectedNewValue: int32Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int32": {
			rule:             Integer[int32](),
			value:            &int32Dummy,
			expectedNewValue: &int32Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"int64": {
			rule:             Integer[int64](),
			value:            int64Dummy,
			expectedNewValue: int64Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*int64": {
			rule:             Integer[int64](),
			value:            &int64Dummy,
			expectedNewValue: &int64Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"uint": {
			rule:             Integer[uint](),
			value:            uintDummy,
			expectedNewValue: uintDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint": {
			rule:             Integer[uint](),
			value:            &uintDummy,
			expectedNewValue: &uintDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"uint8": {
			rule:             Integer[uint8](),
			value:            uint8Dummy,
			expectedNewValue: uint8Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint8": {
			rule:             Integer[uint8](),
			value:            &uint8Dummy,
			expectedNewValue: &uint8Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"uint16": {
			rule:             Integer[uint16](),
			value:            uint16Dummy,
			expectedNewValue: uint16Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint16": {
			rule:             Integer[uint16](),
			value:            &uint16Dummy,
			expectedNewValue: &uint16Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"uint32": {
			rule:             Integer[uint32](),
			value:            uint32Dummy,
			expectedNewValue: uint32Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint32": {
			rule:             Integer[uint32](),
			value:            &uint32Dummy,
			expectedNewValue: &uint32Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"uint64": {
			rule:             Integer[uint64](),
			value:            uint64Dummy,
			expectedNewValue: uint64Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"*uint64": {
			rule:             Integer[uint64](),
			value:            &uint64Dummy,
			expectedNewValue: &uint64Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		"uint when int is wanted": {
			rule:             Integer[int](),
			value:            uint(1),
			expectedNewValue: uint(1),
			expectedError:    NewIntegerValidationError("int", "uint"),
			expectedToBail:   true,
		},
		"float": {
			rule:             Integer[int](),
			value:            1.0,
			expectedNewValue: 1.0,
			expectedError:    NewIntegerValidationError("int", "float64"),
			expectedToBail:   true,
		},
		"complex": {
			rule:             Integer[int](),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewIntegerValidationError("int", "complex128"),
			expectedToBail:   true,
		},
		"string": {
			rule:             Integer[int](),
			value:            "foo",
			expectedNewValue: "foo",
			expectedError:    NewIntegerValidationError("int", "string"),
			expectedToBail:   true,
		},
		"bool": {
			rule:             Integer[int](),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewIntegerValidationError("int", "bool"),
			expectedToBail:   true,
		},
		"map": {
			rule:             Integer[int](),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewIntegerValidationError("int", "map[interface {}]interface {}"),
			expectedToBail:   true,
		},
		"struct": {
			rule:             Integer[int](),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewIntegerValidationError("int", "rule.someStruct"),
			expectedToBail:   true,
		},
	}
}
