package rule

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_EndsWithRule(t *testing.T) {
	runRuleTestCases(t, endsWithRuleDataProvider)
}

func Test_EndsWithValidationError(t *testing.T) {
	// given
	var (
		suffix1Dummy = fakerInstance.Lorem().Sentence(6)
		suffix2Dummy = fakerInstance.Lorem().Sentence(6)
	)

	// when
	err := NewEndsWithValidationError([]string{suffix1Dummy, suffix2Dummy})

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must end with one of the following: %s",
		strings.Join([]string{suffix1Dummy, suffix2Dummy}, "; "),
	))
}

func BenchmarkEndsWithRule(b *testing.B) {
	runRuleBenchmarks(b, endsWithRuleDataProvider)
}

func endsWithRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		suffix1Dummy                = fakerInstance.Lorem().Sentence(3)
		suffix2Dummy                = fakerInstance.Lorem().Sentence(3)
		stringWithSuffix1Dummy      = fmt.Sprintf("%s_%s", fakerInstance.Lorem().Sentence(6), suffix1Dummy)
		stringWithSuffix2Dummy      = fmt.Sprintf("%s_%s", fakerInstance.Lorem().Sentence(6), suffix2Dummy)
		stringWithoutAnySuffixDummy = fakerInstance.UUID().V4()
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             EndsWith(fakerInstance.Lorem().Sentence(6)),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string with suffix #1, search for suffix #1": {
			rule:             EndsWith(suffix1Dummy),
			value:            stringWithSuffix1Dummy,
			expectedNewValue: stringWithSuffix1Dummy,
			expectedError:    nil,
		},
		"*string with suffix #1, search for suffix #1": {
			rule:             EndsWith(suffix1Dummy),
			value:            &stringWithSuffix1Dummy,
			expectedNewValue: &stringWithSuffix1Dummy,
			expectedError:    nil,
		},
		"string with suffix #1, search for suffix #1 and #2": {
			rule:             EndsWith(suffix1Dummy, suffix2Dummy),
			value:            stringWithSuffix1Dummy,
			expectedNewValue: stringWithSuffix1Dummy,
			expectedError:    nil,
		},

		"string with suffix #2, search for suffix #2": {
			rule:             EndsWith(suffix2Dummy),
			value:            stringWithSuffix2Dummy,
			expectedNewValue: stringWithSuffix2Dummy,
			expectedError:    nil,
		},
		"*string with suffix #2, search for suffix #2": {
			rule:             EndsWith(suffix2Dummy),
			value:            &stringWithSuffix2Dummy,
			expectedNewValue: &stringWithSuffix2Dummy,
			expectedError:    nil,
		},
		"string with suffix #2, search for suffix #1 and #2": {
			rule:             EndsWith(suffix1Dummy, suffix2Dummy),
			value:            stringWithSuffix2Dummy,
			expectedNewValue: stringWithSuffix2Dummy,
			expectedError:    nil,
		},

		"string without any suffix, search for suffix #1": {
			rule:             EndsWith(suffix1Dummy),
			value:            stringWithoutAnySuffixDummy,
			expectedNewValue: stringWithoutAnySuffixDummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"string without any suffix, search for suffix #2": {
			rule:             EndsWith(suffix2Dummy),
			value:            stringWithoutAnySuffixDummy,
			expectedNewValue: stringWithoutAnySuffixDummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix2Dummy}), err)
			},
		},
		"string without any suffix, search for suffix #1 and #2": {
			rule:             EndsWith(suffix1Dummy, suffix2Dummy),
			value:            stringWithoutAnySuffixDummy,
			expectedNewValue: stringWithoutAnySuffixDummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy, suffix2Dummy}), err)
			},
		},
		"*string without any suffix, search for suffix #1 and #2": {
			rule:             EndsWith(suffix1Dummy, suffix2Dummy),
			value:            &stringWithoutAnySuffixDummy,
			expectedNewValue: &stringWithoutAnySuffixDummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy, suffix2Dummy}), err)
			},
		},

		// unsupported values
		"int": {
			rule:             EndsWith(suffix1Dummy),
			value:            0,
			expectedNewValue: 0,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"float": {
			rule:             EndsWith(suffix1Dummy),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"complex": {
			rule:             EndsWith(suffix1Dummy),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"bool": {
			rule:             EndsWith(suffix1Dummy),
			value:            true,
			expectedNewValue: true,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"slice": {
			rule:             EndsWith(suffix1Dummy),
			value:            []int{},
			expectedNewValue: []int{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"array": {
			rule:             EndsWith(suffix1Dummy),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"map": {
			rule:             EndsWith(suffix1Dummy),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"struct": {
			rule:             EndsWith(suffix1Dummy),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewEndsWithValidationError([]string{suffix1Dummy}), err)
			},
		},
	}
}
