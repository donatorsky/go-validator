package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ArrayOfRule(t *testing.T) {
	// given
	for ttIdx, tt := range arrayOfRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
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
	for ttIdx, tt := range arrayOfRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func arrayOfRuleDataProvider() []*ruleTestCaseData {
	var (
		arrayOfIntsDummy        = [3]int{1, 2, 3}
		arrayOfIntPointersDummy = [3]*int{ptr(1), ptr(2), ptr(3)}
	)

	return []*ruleTestCaseData{
		{
			rule:             ArrayOf[int](),
			value:            nil,
			expectedNewValue: (*[0]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		{
			rule:             ArrayOf[int](),
			value:            (*[0]int)(nil),
			expectedNewValue: (*[0]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             ArrayOf[int](),
			value:            (*[]string)(nil),
			expectedNewValue: (*[0]int)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		{
			rule:             ArrayOf[int](),
			value:            arrayOfIntsDummy,
			expectedNewValue: arrayOfIntsDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             ArrayOf[int](),
			value:            &arrayOfIntsDummy,
			expectedNewValue: &arrayOfIntsDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             ArrayOf[int](),
			value:            arrayOfIntPointersDummy,
			expectedNewValue: arrayOfIntPointersDummy,
			expectedError:    NewArrayOfValidationError("int", "[3]*int"),
			expectedToBail:   true,
		},
		{
			rule:             ArrayOf[int](),
			value:            &arrayOfIntPointersDummy,
			expectedNewValue: &arrayOfIntPointersDummy,
			expectedError:    NewArrayOfValidationError("int", "[3]*int"),
			expectedToBail:   true,
		},

		{
			rule:             ArrayOf[*int](),
			value:            arrayOfIntsDummy,
			expectedNewValue: arrayOfIntsDummy,
			expectedError:    NewArrayOfValidationError("*int", "[3]int"),
			expectedToBail:   true,
		},
		{
			rule:             ArrayOf[*int](),
			value:            &arrayOfIntsDummy,
			expectedNewValue: &arrayOfIntsDummy,
			expectedError:    NewArrayOfValidationError("*int", "[3]int"),
			expectedToBail:   true,
		},
		{
			rule:             ArrayOf[*int](),
			value:            arrayOfIntPointersDummy,
			expectedNewValue: arrayOfIntPointersDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             ArrayOf[*int](),
			value:            &arrayOfIntPointersDummy,
			expectedNewValue: &arrayOfIntPointersDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		{
			rule:             ArrayOf[int](),
			value:            1,
			expectedNewValue: 1,
			expectedError:    NewArrayOfValidationError("int", "int"),
			expectedToBail:   true,
		},
		{
			rule:             ArrayOf[int](),
			value:            1.0,
			expectedNewValue: 1.0,
			expectedError:    NewArrayOfValidationError("int", "float64"),
			expectedToBail:   true,
		},
		{
			rule:             ArrayOf[int](),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewArrayOfValidationError("int", "complex128"),
			expectedToBail:   true,
		},
		{
			rule:             ArrayOf[int](),
			value:            "foo",
			expectedNewValue: "foo",
			expectedError:    NewArrayOfValidationError("int", "string"),
			expectedToBail:   true,
		},
		{
			rule:             ArrayOf[int](),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewArrayOfValidationError("int", "bool"),
			expectedToBail:   true,
		},
		{
			rule:             ArrayOf[int](),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewArrayOfValidationError("int", "[]int"),
			expectedToBail:   true,
		},
		{
			rule:             ArrayOf[int](),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewArrayOfValidationError("int", "map[interface {}]interface {}"),
			expectedToBail:   true,
		},
		{
			rule:             ArrayOf[int](),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewArrayOfValidationError("int", "rule.someStruct"),
			expectedToBail:   true,
		},
	}
}
