package rule

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_CustomRule(t *testing.T) {
	runRuleTestCases(t, customRuleDataProvider)
}

func Test_CustomValidationError(t *testing.T) {
	// given
	var (
		errorMessage = fakerInstance.Lorem().Sentence(6)
		errorDummy   = errors.New(errorMessage)
	)

	// when
	err := NewCustomValidationError(errorDummy)

	// then
	require.EqualError(t, err, errorMessage)
}

func BenchmarkCustomRule(b *testing.B) {
	runRuleBenchmarks(b, customRuleDataProvider)
}

func customRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		inputDummy        = fakerInstance.Lorem().Sentence(6)
		outputDummy       = fakerInstance.Int()
		errorMessageDummy = fakerInstance.Lorem().Sentence(5)
	)

	return map[string]*ruleTestCaseData{
		"string input, int output": {
			rule: Custom(func(_ context.Context, value string, _ any) (int, error) {
				return outputDummy, nil
			}),
			value:            inputDummy,
			expectedNewValue: outputDummy,
			expectedError:    nil,
		},

		"string input, int output, but int input wanted": {
			rule: Custom(func(_ context.Context, value int, _ any) (int, error) {
				return outputDummy, nil
			}),
			value:            inputDummy,
			expectedNewValue: inputDummy,
			expectedError:    NewCustomValidationError(errors.New("invalid data type provided: string, expected int")),
		},

		"error": {
			rule: Custom(func(_ context.Context, value string, _ any) (int, error) {
				return outputDummy, errors.New(errorMessageDummy)
			}),
			value:            inputDummy,
			expectedNewValue: inputDummy,
			expectedError:    NewCustomValidationError(errors.New(errorMessageDummy)),
		},

		"custom error": {
			rule: Custom(func(_ context.Context, value string, _ any) (int, error) {
				return outputDummy, customError{Value: value}
			}),
			value:            inputDummy,
			expectedNewValue: inputDummy,
			expectedError:    customError{Value: inputDummy},
		},
	}
}

type customError struct {
	Value string
}

func (c customError) GetRule() string {
	return ve.TypeCustom
}

func (c customError) Error() string {
	return c.Value
}
