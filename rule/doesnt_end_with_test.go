package rule

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_DoesntEndWithRule(t *testing.T) {
	runRuleTestCases(t, doesntEndWithRuleDataProvider)
}

func Test_DoesntEndWithValidationError(t *testing.T) {
	// given
	var (
		suffix1Dummy = fakerInstance.Lorem().Sentence(6)
		suffix2Dummy = fakerInstance.Lorem().Sentence(6)
	)

	// when
	err := NewDoesntEndWithValidationError([]string{suffix1Dummy, suffix2Dummy})

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must not end with any of the following: %s",
		strings.Join([]string{suffix1Dummy, suffix2Dummy}, "; "),
	))
}

func BenchmarkDoesntEndWithRule(b *testing.B) {
	runRuleBenchmarks(b, doesntEndWithRuleDataProvider)
}

func doesntEndWithRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		suffix1Dummy                = fakerInstance.Lorem().Sentence(3)
		suffix2Dummy                = fakerInstance.Lorem().Sentence(3)
		stringWithSuffix1Dummy      = fmt.Sprintf("%s_%s", fakerInstance.Lorem().Sentence(6), suffix1Dummy)
		stringWithSuffix2Dummy      = fmt.Sprintf("%s_%s", fakerInstance.Lorem().Sentence(6), suffix2Dummy)
		stringWithoutAnySuffixDummy = fakerInstance.UUID().V4()
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             DoesntEndWith(fakerInstance.Lorem().Sentence(6)),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string with suffix #1, search for suffix #1": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            stringWithSuffix1Dummy,
			expectedNewValue: stringWithSuffix1Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"*string with suffix #1, search for suffix #1": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            &stringWithSuffix1Dummy,
			expectedNewValue: &stringWithSuffix1Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"string with suffix #1, search for suffix #1 and #2": {
			rule:             DoesntEndWith(suffix1Dummy, suffix2Dummy),
			value:            stringWithSuffix1Dummy,
			expectedNewValue: stringWithSuffix1Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy, suffix2Dummy}), err)
			},
		},

		"string with suffix #2, search for suffix #2": {
			rule:             DoesntEndWith(suffix2Dummy),
			value:            stringWithSuffix2Dummy,
			expectedNewValue: stringWithSuffix2Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix2Dummy}), err)
			},
		},
		"*string with suffix #2, search for suffix #2": {
			rule:             DoesntEndWith(suffix2Dummy),
			value:            &stringWithSuffix2Dummy,
			expectedNewValue: &stringWithSuffix2Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix2Dummy}), err)
			},
		},
		"string with suffix #2, search for suffix #1 and #2": {
			rule:             DoesntEndWith(suffix1Dummy, suffix2Dummy),
			value:            stringWithSuffix2Dummy,
			expectedNewValue: stringWithSuffix2Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy, suffix2Dummy}), err)
			},
		},

		"string without any suffix, search for suffix #1": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            stringWithoutAnySuffixDummy,
			expectedNewValue: stringWithoutAnySuffixDummy,
		},
		"string without any suffix, search for suffix #2": {
			rule:             DoesntEndWith(suffix2Dummy),
			value:            stringWithoutAnySuffixDummy,
			expectedNewValue: stringWithoutAnySuffixDummy,
		},
		"string without any suffix, search for suffix #1 and #2": {
			rule:             DoesntEndWith(suffix1Dummy, suffix2Dummy),
			value:            stringWithoutAnySuffixDummy,
			expectedNewValue: stringWithoutAnySuffixDummy,
		},
		"*string without any suffix, search for suffix #1 and #2": {
			rule:             DoesntEndWith(suffix1Dummy, suffix2Dummy),
			value:            &stringWithoutAnySuffixDummy,
			expectedNewValue: &stringWithoutAnySuffixDummy,
		},

		// unsupported values
		"int": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            0,
			expectedNewValue: 0,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"float": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"complex": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"bool": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            true,
			expectedNewValue: true,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"slice": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            []int{},
			expectedNewValue: []int{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"array": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"map": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy}), err)
			},
		},
		"struct": {
			rule:             DoesntEndWith(suffix1Dummy),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntEndWithValidationError([]string{suffix1Dummy}), err)
			},
		},
	}
}
