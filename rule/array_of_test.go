package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ArrayOfRule(t *testing.T) {
	runRuleTestCases(t, arrayOfRuleDataProvider)
}

func Test_ArrayOfValidationError(t *testing.T) {
	// given
	var (
		expectedTypeDummy = fakerInstance.Lorem().Word()
		actualTypeDummy   = fakerInstance.Lorem().Word()
	)

	// when
	err := NewArrayOfValidationError(expectedTypeDummy, actualTypeDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must be an array of %q, but is %q",
		expectedTypeDummy,
		actualTypeDummy,
	))
}

func BenchmarkArrayOfRule(b *testing.B) {
	runRuleBenchmarks(b, arrayOfRuleDataProvider)
}

func arrayOfRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		arrayOfIntsDummy        = [3]int{1, 2, 3}
		arrayOfIntPointersDummy = [3]*int{ptr(1), ptr(2), ptr(3)}
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             ArrayOf[int](),
			value:            nil,
			expectedNewValue: (*[0]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"pointer to array of ints nil pointer": {
			rule:             ArrayOf[int](),
			value:            (*[0]int)(nil),
			expectedNewValue: (*[0]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to array of strings nil pointer, array of ints wanted": {
			rule:             ArrayOf[int](),
			value:            (*[]string)(nil),
			expectedNewValue: (*[0]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"array of ints": {
			rule:             ArrayOf[int](),
			value:            arrayOfIntsDummy,
			expectedNewValue: arrayOfIntsDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to array of ints": {
			rule:             ArrayOf[int](),
			value:            &arrayOfIntsDummy,
			expectedNewValue: &arrayOfIntsDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"array of int pointers, array of ints wanted": {
			rule:             ArrayOf[int](),
			value:            arrayOfIntPointersDummy,
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("int", "[3]*int"),
			expectedToBail:   true,
		},
		"pointer to array of int pointers, array of ints wanted": {
			rule:             ArrayOf[int](),
			value:            &arrayOfIntPointersDummy,
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("int", "[3]*int"),
			expectedToBail:   true,
		},

		"array of ints, array of int pointers wanted": {
			rule:             ArrayOf[*int](),
			value:            arrayOfIntsDummy,
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("*int", "[3]int"),
			expectedToBail:   true,
		},
		"pointer to array of ints, array of int pointers wanted": {
			rule:             ArrayOf[*int](),
			value:            &arrayOfIntsDummy,
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("*int", "[3]int"),
			expectedToBail:   true,
		},
		"array of int pointers, array of int pointers wanted": {
			rule:             ArrayOf[*int](),
			value:            arrayOfIntPointersDummy,
			expectedNewValue: arrayOfIntPointersDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to array of int pointers, array of int pointers wanted": {
			rule:             ArrayOf[*int](),
			value:            &arrayOfIntPointersDummy,
			expectedNewValue: &arrayOfIntPointersDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		"int": {
			rule:             ArrayOf[int](),
			value:            1,
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("int", "int"),
			expectedToBail:   true,
		},
		"float": {
			rule:             ArrayOf[int](),
			value:            1.0,
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("int", "float64"),
			expectedToBail:   true,
		},
		"complex": {
			rule:             ArrayOf[int](),
			value:            1 + 2i,
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("int", "complex128"),
			expectedToBail:   true,
		},
		"string": {
			rule:             ArrayOf[int](),
			value:            "foo",
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("int", "string"),
			expectedToBail:   true,
		},
		"bool": {
			rule:             ArrayOf[int](),
			value:            true,
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("int", "bool"),
			expectedToBail:   true,
		},
		"slice": {
			rule:             ArrayOf[int](),
			value:            []int{},
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("int", "[]int"),
			expectedToBail:   true,
		},
		"map": {
			rule:             ArrayOf[int](),
			value:            map[any]any{},
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("int", "map[interface {}]interface {}"),
			expectedToBail:   true,
		},
		"struct": {
			rule:             ArrayOf[int](),
			value:            someStruct{},
			expectedNewValue: nil,
			expectedError:    NewArrayOfValidationError("int", "rule.someStruct"),
			expectedToBail:   true,
		},
	}
}
