package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SliceRule(t *testing.T) {
	// given
	for ttIdx, tt := range sliceRuleDataProvider() {
		runRuleTestCase(t, ttIdx, tt)
	}
}

func Test_SliceValidationError(t *testing.T) {
	// when
	err := NewSliceValidationError()

	// then
	require.EqualError(t, err, "must be a slice")
}

func BenchmarkSliceRule(b *testing.B) {
	for ttIdx, tt := range sliceRuleDataProvider() {
		runRuleBenchmark(b, ttIdx, tt)
	}
}

func sliceRuleDataProvider() []*ruleTestCaseData {
	var sliceDummy = []int{1, 2, 3}

	return []*ruleTestCaseData{
		{
			rule:             Slice(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   false,
		},

		{
			rule:             Slice(),
			value:            (*[]any)(nil),
			expectedNewValue: (*[]any)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		{
			rule:             Slice(),
			value:            sliceDummy,
			expectedNewValue: sliceDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		{
			rule:             Slice(),
			value:            &sliceDummy,
			expectedNewValue: &sliceDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		{
			rule:             Slice(),
			value:            1,
			expectedNewValue: 1,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Slice(),
			value:            1.0,
			expectedNewValue: 1.0,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Slice(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Slice(),
			value:            "foo",
			expectedNewValue: "foo",
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Slice(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Slice(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Slice(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		{
			rule:             Slice(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
	}
}
