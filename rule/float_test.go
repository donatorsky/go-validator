package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FloatRule(t *testing.T) {
	// given
	for ttIdx, tt := range floatRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_FloatValidationError(t *testing.T) {
	// given
	var (
		expectedTypeDummy = fakerInstance.Lorem().Word()
		actualTypeDummy   = fakerInstance.Lorem().Word()
	)

	// when
	err := NewFloatValidationError(expectedTypeDummy, actualTypeDummy)

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must be a %s but is %s",
		expectedTypeDummy,
		actualTypeDummy,
	))
}

func BenchmarkFloatRule(b *testing.B) {
	for ttIdx, tt := range floatRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func floatRuleDataProvider() []*ruleTestCaseData {
	var (
		float32Dummy = fakerInstance.Float32(5, -1000, 1000)
		float64Dummy = fakerInstance.Float64(5, -1000, 1000)
	)

	return []*ruleTestCaseData{
		{
			rule:             Float[float32](),
			value:            nil,
			expectedNewValue: (*float32)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Float[float64](),
			value:            nil,
			expectedNewValue: (*float64)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		{
			rule:             Float[float32](),
			value:            float32Dummy,
			expectedNewValue: float32Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Float[float32](),
			value:            &float32Dummy,
			expectedNewValue: &float32Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Float[float32](),
			value:            float64Dummy,
			expectedNewValue: float64Dummy,
			expectedError:    NewFloatValidationError("float32", "float64"),
			expectedToBail:   true,
		},
		{
			rule:             Float[float32](),
			value:            &float64Dummy,
			expectedNewValue: &float64Dummy,
			expectedError:    NewFloatValidationError("float32", "float64"),
			expectedToBail:   true,
		},

		{
			rule:             Float[float64](),
			value:            float64Dummy,
			expectedNewValue: float64Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Float[float64](),
			value:            &float64Dummy,
			expectedNewValue: &float64Dummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Float[float64](),
			value:            float32Dummy,
			expectedNewValue: float32Dummy,
			expectedError:    NewFloatValidationError("float64", "float32"),
			expectedToBail:   true,
		},
		{
			rule:             Float[float64](),
			value:            &float32Dummy,
			expectedNewValue: &float32Dummy,
			expectedError:    NewFloatValidationError("float64", "float32"),
			expectedToBail:   true,
		},

		// unsupported values
		{
			rule:             Float[float64](),
			value:            1,
			expectedNewValue: 1,
			expectedError:    NewFloatValidationError("float64", "int"),
			expectedToBail:   true,
		},
		{
			rule:             Float[float64](),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewFloatValidationError("float64", "complex128"),
			expectedToBail:   true,
		},
		{
			rule:             Float[float64](),
			value:            "foo",
			expectedNewValue: "foo",
			expectedError:    NewFloatValidationError("float64", "string"),
			expectedToBail:   true,
		},
		{
			rule:             Float[float64](),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewFloatValidationError("float64", "bool"),
			expectedToBail:   true,
		},
		{
			rule:             Float[float64](),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewFloatValidationError("float64", "map[interface {}]interface {}"),
			expectedToBail:   true,
		},
		{
			rule:             Float[float64](),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewFloatValidationError("float64", "rule.someStruct"),
			expectedToBail:   true,
		},
	}
}
