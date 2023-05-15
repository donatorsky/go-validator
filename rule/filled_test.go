package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FilledRule(t *testing.T) {
	runRuleTestCases(t, filledRuleDataProvider)
}

func Test_FilledValidationError(t *testing.T) {
	// when
	err := NewFilledValidationError()

	// then
	require.EqualError(t, err, "must not be empty")
}

func BenchmarkFilledRule(b *testing.B) {
	runRuleBenchmarks(b, filledRuleDataProvider)
}

func filledRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		intZeroValueDummy     int
		intValueDummy         int = fakerInstance.IntBetween(1, 1000)
		floatZeroValueDummy   float64
		floatValueDummy       float64 = fakerInstance.Float64(5, 1, 1000)
		complexZeroValueDummy complex128
		complexValueDummy     complex128 = complex(fakerInstance.Float64(5, 1, 1000), fakerInstance.Float64(5, 1, 1000))
		boolZeroValueDummy    bool
		boolValueDummy        bool = true
		stringZeroValueDummy  string
		stringValueDummy      string = fakerInstance.Lorem().Sentence(2)
		sliceZeroValueDummy   []any
		sliceValueDummy       []any = []any{1}
		arrayZeroValueDummy   [1]any
		arrayValueDummy       [1]any = [1]any{1}
		mapZeroValueDummy     map[any]any
		mapValueDummy         map[any]any = map[any]any{"foo": 123}
		structZeroValueDummy  someStruct
		structValueDummy      someStruct = someStruct{Foo: "bar"}
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             Filled(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"int zero value": {
			rule:             Filled(),
			value:            intZeroValueDummy,
			expectedNewValue: intZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"*int zero value": {
			rule:             Filled(),
			value:            &intZeroValueDummy,
			expectedNewValue: &intZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"int": {
			rule:             Filled(),
			value:            intValueDummy,
			expectedNewValue: intValueDummy,
			expectedError:    nil,
		},
		"*int": {
			rule:             Filled(),
			value:            &intValueDummy,
			expectedNewValue: &intValueDummy,
			expectedError:    nil,
		},
		"*int nil pointer": {
			rule:             Filled(),
			value:            (*int)(nil),
			expectedNewValue: (*int)(nil),
			expectedError:    nil,
		},

		"float zero value": {
			rule:             Filled(),
			value:            floatZeroValueDummy,
			expectedNewValue: floatZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"*float zero value": {
			rule:             Filled(),
			value:            &floatZeroValueDummy,
			expectedNewValue: &floatZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"float": {
			rule:             Filled(),
			value:            floatValueDummy,
			expectedNewValue: floatValueDummy,
			expectedError:    nil,
		},
		"*float": {
			rule:             Filled(),
			value:            &floatValueDummy,
			expectedNewValue: &floatValueDummy,
			expectedError:    nil,
		},
		"*float nil pointer": {
			rule:             Filled(),
			value:            (*float64)(nil),
			expectedNewValue: (*float64)(nil),
			expectedError:    nil,
		},

		"complex zero value": {
			rule:             Filled(),
			value:            complexZeroValueDummy,
			expectedNewValue: complexZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"*complex zero value": {
			rule:             Filled(),
			value:            &complexZeroValueDummy,
			expectedNewValue: &complexZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"complex": {
			rule:             Filled(),
			value:            complexValueDummy,
			expectedNewValue: complexValueDummy,
			expectedError:    nil,
		},
		"*complex": {
			rule:             Filled(),
			value:            &complexValueDummy,
			expectedNewValue: &complexValueDummy,
			expectedError:    nil,
		},
		"*complex nil pointer": {
			rule:             Filled(),
			value:            (*complex128)(nil),
			expectedNewValue: (*complex128)(nil),
			expectedError:    nil,
		},

		"bool": {
			rule:             Filled(),
			value:            boolZeroValueDummy,
			expectedNewValue: boolZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"*bool": {
			rule:             Filled(),
			value:            &boolZeroValueDummy,
			expectedNewValue: &boolZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"bool zero value": {
			rule:             Filled(),
			value:            boolValueDummy,
			expectedNewValue: boolValueDummy,
			expectedError:    nil,
		},
		"*bool zero value": {
			rule:             Filled(),
			value:            &boolValueDummy,
			expectedNewValue: &boolValueDummy,
			expectedError:    nil,
		},
		"*bool nil pointer": {
			rule:             Filled(),
			value:            (*bool)(nil),
			expectedNewValue: (*bool)(nil),
			expectedError:    nil,
		},

		"string zero value": {
			rule:             Filled(),
			value:            stringZeroValueDummy,
			expectedNewValue: stringZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"*string zero value": {
			rule:             Filled(),
			value:            &stringZeroValueDummy,
			expectedNewValue: &stringZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"string": {
			rule:             Filled(),
			value:            stringValueDummy,
			expectedNewValue: stringValueDummy,
			expectedError:    nil,
		},
		"*string": {
			rule:             Filled(),
			value:            &stringValueDummy,
			expectedNewValue: &stringValueDummy,
			expectedError:    nil,
		},
		"*string nil pointer": {
			rule:             Filled(),
			value:            (*string)(nil),
			expectedNewValue: (*string)(nil),
			expectedError:    nil,
		},

		"slice zero value": {
			rule:             Filled(),
			value:            sliceZeroValueDummy,
			expectedNewValue: sliceZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"*slice zero value": {
			rule:             Filled(),
			value:            &sliceZeroValueDummy,
			expectedNewValue: &sliceZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"slice": {
			rule:             Filled(),
			value:            sliceValueDummy,
			expectedNewValue: sliceValueDummy,
			expectedError:    nil,
		},
		"*slice": {
			rule:             Filled(),
			value:            &sliceValueDummy,
			expectedNewValue: &sliceValueDummy,
			expectedError:    nil,
		},
		"*slice nil pointer": {
			rule:             Filled(),
			value:            (*[]int)(nil),
			expectedNewValue: (*[]int)(nil),
			expectedError:    nil,
		},

		"array zero value": {
			rule:             Filled(),
			value:            arrayZeroValueDummy,
			expectedNewValue: arrayZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"*array zero value": {
			rule:             Filled(),
			value:            &arrayZeroValueDummy,
			expectedNewValue: &arrayZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"array": {
			rule:             Filled(),
			value:            arrayValueDummy,
			expectedNewValue: arrayValueDummy,
			expectedError:    nil,
		},
		"*array": {
			rule:             Filled(),
			value:            &arrayValueDummy,
			expectedNewValue: &arrayValueDummy,
			expectedError:    nil,
		},
		"*array nil pointer": {
			rule:             Filled(),
			value:            (*[1]any)(nil),
			expectedNewValue: (*[1]any)(nil),
			expectedError:    nil,
		},

		"map zero value": {
			rule:             Filled(),
			value:            mapZeroValueDummy,
			expectedNewValue: mapZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"*map zero value": {
			rule:             Filled(),
			value:            &mapZeroValueDummy,
			expectedNewValue: &mapZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"map": {
			rule:             Filled(),
			value:            mapValueDummy,
			expectedNewValue: mapValueDummy,
			expectedError:    nil,
		},
		"*map": {
			rule:             Filled(),
			value:            &mapValueDummy,
			expectedNewValue: &mapValueDummy,
			expectedError:    nil,
		},
		"map nil pointer": {
			rule:             Filled(),
			value:            (*map[any]any)(nil),
			expectedNewValue: (*map[any]any)(nil),
			expectedError:    nil,
		},

		"struct zero value": {
			rule:             Filled(),
			value:            structZeroValueDummy,
			expectedNewValue: structZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"*struct zero value": {
			rule:             Filled(),
			value:            &structZeroValueDummy,
			expectedNewValue: &structZeroValueDummy,
			expectedError:    NewFilledValidationError(),
		},
		"struct": {
			rule:             Filled(),
			value:            structValueDummy,
			expectedNewValue: structValueDummy,
			expectedError:    nil,
		},
		"*struct": {
			rule:             Filled(),
			value:            &structValueDummy,
			expectedNewValue: &structValueDummy,
			expectedError:    nil,
		},
		"*struct nil pointer": {
			rule:             Filled(),
			value:            (*someStruct)(nil),
			expectedNewValue: (*someStruct)(nil),
			expectedError:    nil,
		},
	}
}
