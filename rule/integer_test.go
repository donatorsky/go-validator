package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_integerRule_Apply_ValidType(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		rule  *integerRule[int]
		value any
	}{
		{
			rule:  Integer[int](),
			value: 123,
		},
	} {
		t.Run(fmt.Sprintf("Test Data #%d: %T", ttIdx, tt.value), func(t *testing.T) {
			// when
			newValue, err := tt.rule.Apply(nil, tt.value, nil)

			// then
			require.NoError(t, err)
			require.Equal(t, tt.value, newValue)
		})
	}
}

func Test_integerRule_Apply_Nil(t *testing.T) {
	// given
	rule := Integer[int]()

	// when
	newValue, err := rule.Apply(nil, nil, nil)

	// then
	require.NoError(t, err)
	require.Nil(t, newValue)
}

func Test_integerRule_Apply_InvalidType(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		rule         *integerRule[int]
		value        any
		expectedType string
		actualType   string
	}{
		{
			rule:         Integer[int](),
			value:        "123",
			expectedType: "int",
			actualType:   "string",
		},
	} {
		t.Run(fmt.Sprintf("Test Data #%d: %T", ttIdx, tt.value), func(t *testing.T) {
			// when
			newValue, err := tt.rule.Apply(nil, tt.value, nil)

			// then
			require.Equal(t, tt.value, newValue)

			var ruleErr IntegerValidationError
			require.ErrorAs(t, err, &ruleErr)
			require.Equal(t, ve.TypeInt, ruleErr.Rule)
			require.Equal(t, tt.expectedType, ruleErr.ExpectedType)
			require.Equal(t, tt.actualType, ruleErr.ActualType)
		})
	}
}
