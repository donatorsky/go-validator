package rule

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_DoesntStartWithRule(t *testing.T) {
	runRuleTestCases(t, doesntStartWithRuleDataProvider)
}

func Test_DoesntStartWithValidationError(t *testing.T) {
	// given
	var (
		prefix1Dummy = fakerInstance.Lorem().Sentence(6)
		prefix2Dummy = fakerInstance.Lorem().Sentence(6)
	)

	// when
	err := NewDoesntStartWithValidationError([]string{prefix1Dummy, prefix2Dummy})

	// then
	require.EqualError(t, err, fmt.Sprintf(
		"must not start with any of the following: %s",
		strings.Join([]string{prefix1Dummy, prefix2Dummy}, "; "),
	))
}

func BenchmarkDoesntStartWithRule(b *testing.B) {
	runRuleBenchmarks(b, doesntStartWithRuleDataProvider)
}

func doesntStartWithRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		prefix1Dummy                = fakerInstance.Lorem().Sentence(3)
		prefix2Dummy                = fakerInstance.Lorem().Sentence(3)
		stringWithPrefix1Dummy      = fmt.Sprintf("%s_%s", prefix1Dummy, fakerInstance.Lorem().Sentence(6))
		stringWithPrefix2Dummy      = fmt.Sprintf("%s_%s", prefix2Dummy, fakerInstance.Lorem().Sentence(6))
		stringWithoutAnyPrefixDummy = fakerInstance.UUID().V4()
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             DoesntStartWith(fakerInstance.Lorem().Sentence(6)),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string with prefix #1, search for prefix #1": {
			rule:             DoesntStartWith(prefix1Dummy),
			value:            stringWithPrefix1Dummy,
			expectedNewValue: stringWithPrefix1Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy}), err)
			},
		},
		"*string with prefix #1, search for prefix #1": {
			rule:             DoesntStartWith(prefix1Dummy),
			value:            &stringWithPrefix1Dummy,
			expectedNewValue: &stringWithPrefix1Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy}), err)
			},
		},
		"string with prefix #1, search for prefix #1 and #2": {
			rule:             DoesntStartWith(prefix1Dummy, prefix2Dummy),
			value:            stringWithPrefix1Dummy,
			expectedNewValue: stringWithPrefix1Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy, prefix2Dummy}), err)
			},
		},

		"string with prefix #2, search for prefix #2": {
			rule:             DoesntStartWith(prefix2Dummy),
			value:            stringWithPrefix2Dummy,
			expectedNewValue: stringWithPrefix2Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix2Dummy}), err)
			},
		},
		"*string with prefix #2, search for prefix #2": {
			rule:             DoesntStartWith(prefix2Dummy),
			value:            &stringWithPrefix2Dummy,
			expectedNewValue: &stringWithPrefix2Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix2Dummy}), err)
			},
		},
		"string with prefix #2, search for prefix #1 and #2": {
			rule:             DoesntStartWith(prefix1Dummy, prefix2Dummy),
			value:            stringWithPrefix2Dummy,
			expectedNewValue: stringWithPrefix2Dummy,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy, prefix2Dummy}), err)
			},
		},

		"string without any prefix, search for prefix #1": {
			rule:              DoesntStartWith(prefix1Dummy),
			value:             stringWithoutAnyPrefixDummy,
			expectedNewValue:  stringWithoutAnyPrefixDummy,
			expectedErrorFunc: nil,
		},
		"string without any prefix, search for prefix #2": {
			rule:              DoesntStartWith(prefix2Dummy),
			value:             stringWithoutAnyPrefixDummy,
			expectedNewValue:  stringWithoutAnyPrefixDummy,
			expectedErrorFunc: nil,
		},
		"string without any prefix, search for prefix #1 and #2": {
			rule:              DoesntStartWith(prefix1Dummy, prefix2Dummy),
			value:             stringWithoutAnyPrefixDummy,
			expectedNewValue:  stringWithoutAnyPrefixDummy,
			expectedErrorFunc: nil,
		},
		"*string without any prefix, search for prefix #1 and #2": {
			rule:              DoesntStartWith(prefix1Dummy, prefix2Dummy),
			value:             &stringWithoutAnyPrefixDummy,
			expectedNewValue:  &stringWithoutAnyPrefixDummy,
			expectedErrorFunc: nil,
		},

		// unsupported values
		"int": {
			rule:             DoesntStartWith(prefix1Dummy),
			value:            0,
			expectedNewValue: 0,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy}), err)
			},
		},
		"float": {
			rule:             DoesntStartWith(prefix1Dummy),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy}), err)
			},
		},
		"complex": {
			rule:             DoesntStartWith(prefix1Dummy),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy}), err)
			},
		},
		"bool": {
			rule:             DoesntStartWith(prefix1Dummy),
			value:            true,
			expectedNewValue: true,
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy}), err)
			},
		},
		"slice": {
			rule:             DoesntStartWith(prefix1Dummy),
			value:            []int{},
			expectedNewValue: []int{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy}), err)
			},
		},
		"array": {
			rule:             DoesntStartWith(prefix1Dummy),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy}), err)
			},
		},
		"map": {
			rule:             DoesntStartWith(prefix1Dummy),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy}), err)
			},
		},
		"struct": {
			rule:             DoesntStartWith(prefix1Dummy),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedErrorFunc: func(t *testing.T, err ve.ValidationError) bool {
				return assert.Equal(t, NewDoesntStartWithValidationError([]string{prefix1Dummy}), err)
			},
		},
	}
}
