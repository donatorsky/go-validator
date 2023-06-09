package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_LengthRule(t *testing.T) {
	runRuleTestCases(t, lengthRuleDataProvider)
}

func Test_LengthValidationError(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		error           error
		expectedMessage string
	}{
		{
			error:           NewLengthValidationError(ve.TypeString, 5),
			expectedMessage: "must be exactly 5 characters long",
		},
		{
			error:           NewLengthValidationError(ve.TypeSlice, 5),
			expectedMessage: "must have exactly 5 items",
		},
		{
			error:           NewLengthValidationError(ve.TypeArray, 5),
			expectedMessage: "must have exactly 5 items",
		},
		{
			error:           NewLengthValidationError(ve.TypeMap, 5),
			expectedMessage: "must have exactly 5 items",
		},
		{
			error:           NewLengthValidationError(ve.TypeInvalid, fakerInstance.Int()),
			expectedMessage: "length cannot be determined",
		},
	} {
		t.Run(fmt.Sprintf("#%d", ttIdx), func(t *testing.T) {
			// then
			require.EqualError(t, tt.error, tt.expectedMessage)
		})
	}
}

func BenchmarkLengthRule(b *testing.B) {
	runRuleBenchmarks(b, lengthRuleDataProvider)
}

func lengthRuleDataProvider() map[string]*ruleTestCaseData {
	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Length(3),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"too short string": {
			rule:             Length(3),
			value:            "ab",
			expectedNewValue: "ab",
			expectedError:    NewLengthValidationError(ve.TypeString, 3),
		},
		"exact length string": {
			rule:             Length(3),
			value:            "abc",
			expectedNewValue: "abc",
			expectedError:    nil,
		},
		"too long string": {
			rule:             Length(3),
			value:            "abcd",
			expectedNewValue: "abcd",
			expectedError:    NewLengthValidationError(ve.TypeString, 3),
		},

		"too short slice": {
			rule:             Length(3),
			value:            []int{1, 2},
			expectedNewValue: []int{1, 2},
			expectedError:    NewLengthValidationError(ve.TypeSlice, 3),
		},
		"slice of exact length": {
			rule:             Length(3),
			value:            []int{1, 2, 3},
			expectedNewValue: []int{1, 2, 3},
			expectedError:    nil,
		},
		"too long slice": {
			rule:             Length(3),
			value:            []int{1, 2, 3, 4},
			expectedNewValue: []int{1, 2, 3, 4},
			expectedError:    NewLengthValidationError(ve.TypeSlice, 3),
		},

		"too short array": {
			rule:             Length(3),
			value:            [2]int{1, 2},
			expectedNewValue: [2]int{1, 2},
			expectedError:    NewLengthValidationError(ve.TypeArray, 3),
		},
		"array of exact length": {
			rule:             Length(3),
			value:            [3]int{1, 2, 3},
			expectedNewValue: [3]int{1, 2, 3},
			expectedError:    nil,
		},
		"too long array": {
			rule:             Length(3),
			value:            [4]int{1, 2, 3, 4},
			expectedNewValue: [4]int{1, 2, 3, 4},
			expectedError:    NewLengthValidationError(ve.TypeArray, 3),
		},

		"map with not enough # of keys": {
			rule:             Length(3),
			value:            map[any]int{1: 1, "a": 2},
			expectedNewValue: map[any]int{1: 1, "a": 2},
			expectedError:    NewLengthValidationError(ve.TypeMap, 3),
		},
		"map with exact # of keys": {
			rule:             Length(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3},
			expectedError:    nil,
		},
		"map with too many # of keys": {
			rule:             Length(3),
			value:            map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedNewValue: map[any]int{1: 1, "a": 2, 3.0: 3, true: 4},
			expectedError:    NewLengthValidationError(ve.TypeMap, 3),
		},

		// unsupported values
		"int": {
			rule:             Length(3),
			value:            1,
			expectedNewValue: 1,
			expectedError:    NewLengthValidationError(ve.TypeInvalid, 3),
		},
		"float": {
			rule:             Length(3),
			value:            1.0,
			expectedNewValue: 1.0,
			expectedError:    NewLengthValidationError(ve.TypeInvalid, 3),
		},
		"complex": {
			rule:             Length(3),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewLengthValidationError(ve.TypeInvalid, 3),
		},
		"bool": {
			rule:             Length(3),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewLengthValidationError(ve.TypeInvalid, 3),
		},
		"struct": {
			rule:             Length(3),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewLengthValidationError(ve.TypeInvalid, 3),
		},
	}
}
