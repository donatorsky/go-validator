package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SliceRule(t *testing.T) {
	runRuleTestCases(t, sliceRuleDataProvider)
}

func Test_SliceValidationError(t *testing.T) {
	// when
	err := NewSliceValidationError()

	// then
	require.EqualError(t, err, "must be a slice")
}

func BenchmarkSliceRule(b *testing.B) {
	runRuleBenchmarks(b, sliceRuleDataProvider)
}

func sliceRuleDataProvider() map[string]*ruleTestCaseData {
	var sliceDummy = []int{1, 2, 3}

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Slice(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
			expectedToBail:   false,
		},

		"pointer to slice nil pointer": {
			rule:             Slice(),
			value:            (*[]any)(nil),
			expectedNewValue: (*[]any)(nil),
			expectedError:    nil,
			expectedToBail:   false,
		},

		"slice": {
			rule:             Slice(),
			value:            sliceDummy,
			expectedNewValue: sliceDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},
		"pointer to slice": {
			rule:             Slice(),
			value:            &sliceDummy,
			expectedNewValue: &sliceDummy,
			expectedError:    nil,
			expectedToBail:   false,
		},

		// unsupported values
		"int": {
			rule:             Slice(),
			value:            1,
			expectedNewValue: nil,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		"float": {
			rule:             Slice(),
			value:            1.0,
			expectedNewValue: nil,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		"complex": {
			rule:             Slice(),
			value:            1 + 2i,
			expectedNewValue: nil,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		"string": {
			rule:             Slice(),
			value:            "foo",
			expectedNewValue: nil,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		"bool": {
			rule:             Slice(),
			value:            true,
			expectedNewValue: nil,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		"array": {
			rule:             Slice(),
			value:            [1]int{},
			expectedNewValue: nil,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		"map": {
			rule:             Slice(),
			value:            map[any]any{},
			expectedNewValue: nil,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
		"struct": {
			rule:             Slice(),
			value:            someStruct{},
			expectedNewValue: nil,
			expectedError:    NewSliceValidationError(),
			expectedToBail:   true,
		},
	}
}
