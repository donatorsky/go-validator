package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SliceOfRule(t *testing.T) {
	runRuleTestCases(t, sliceOfRuleDataProvider)
}

func Test_SliceOfValidationError(t *testing.T) {
	// given
	var (
		expectedTypeDummy = fakerInstance.Lorem().Word()
		actualTypeDummy   = fakerInstance.Lorem().Word()
	)

	// when
	err := NewSliceOfValidationError(expectedTypeDummy, actualTypeDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must be a slice of %q, but is %q",
		expectedTypeDummy,
		actualTypeDummy,
	))
}

func BenchmarkSliceOfRule(b *testing.B) {
	runRuleBenchmarks(b, sliceOfRuleDataProvider)
}

func sliceOfRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		sliceOfIntsDummy        = []int{1, 2, 3}
		sliceOfIntPointersDummy = []*int{ptr(1), ptr(2), ptr(3)}
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             SliceOf[int](),
			value:            nil,
			expectedNewValue: ([]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"slice of ints nil pointer": {
			rule:             SliceOf[int](),
			value:            ([]int)(nil),
			expectedNewValue: ([]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to slice of ints nil pointer": {
			rule:             SliceOf[int](),
			value:            (*[]int)(nil),
			expectedNewValue: ([]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to slice of strings nil pointer, slice of ints wanted": {
			rule:             SliceOf[int](),
			value:            (*[]string)(nil),
			expectedNewValue: ([]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"slice of ints": {
			rule:             SliceOf[int](),
			value:            sliceOfIntsDummy,
			expectedNewValue: sliceOfIntsDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to slice of ints": {
			rule:             SliceOf[int](),
			value:            &sliceOfIntsDummy,
			expectedNewValue: &sliceOfIntsDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"slice of int pointers, slice of int wanted": {
			rule:             SliceOf[int](),
			value:            sliceOfIntPointersDummy,
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("int", "[]*int"),
			expectedToBail:   true,
		},
		"pointer to slice of int pointers": {
			rule:             SliceOf[int](),
			value:            &sliceOfIntPointersDummy,
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("int", "[]*int"),
			expectedToBail:   true,
		},

		"slice of ints, slice of int pointers wanted": {
			rule:             SliceOf[*int](),
			value:            sliceOfIntsDummy,
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("*int", "[]int"),
			expectedToBail:   true,
		},
		"pointer to slice of ints, slice of int pointers wanted": {
			rule:             SliceOf[*int](),
			value:            &sliceOfIntsDummy,
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("*int", "[]int"),
			expectedToBail:   true,
		},
		"slice of int pointers, slice of int pointers wanted": {
			rule:             SliceOf[*int](),
			value:            sliceOfIntPointersDummy,
			expectedNewValue: sliceOfIntPointersDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to slice of int pointers, slice of int pointers wanted": {
			rule:             SliceOf[*int](),
			value:            &sliceOfIntPointersDummy,
			expectedNewValue: &sliceOfIntPointersDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		"int": {
			rule:             SliceOf[int](),
			value:            1,
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("int", "int"),
			expectedToBail:   true,
		},
		"float": {
			rule:             SliceOf[int](),
			value:            1.0,
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("int", "float64"),
			expectedToBail:   true,
		},
		"complex": {
			rule:             SliceOf[int](),
			value:            1 + 2i,
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("int", "complex128"),
			expectedToBail:   true,
		},
		"string": {
			rule:             SliceOf[int](),
			value:            "foo",
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("int", "string"),
			expectedToBail:   true,
		},
		"bool": {
			rule:             SliceOf[int](),
			value:            true,
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("int", "bool"),
			expectedToBail:   true,
		},
		"array": {
			rule:             SliceOf[int](),
			value:            [1]int{},
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("int", "[1]int"),
			expectedToBail:   true,
		},
		"map": {
			rule:             SliceOf[int](),
			value:            map[any]any{},
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("int", "map[interface {}]interface {}"),
			expectedToBail:   true,
		},
		"struct": {
			rule:             SliceOf[int](),
			value:            someStruct{},
			expectedNewValue: nil,
			expectedError:    NewSliceOfValidationError("int", "rule.someStruct"),
			expectedToBail:   true,
		},
	}
}
