package rule

import (
	"context"
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

var fakerInstance = faker.New()

type someStruct struct {
	Foo string
}

type ruleTestCaseData struct {
	rule                 Rule
	value                any
	data                 any
	expectedNewValue     any
	expectedNewValueFunc func(value any) bool
	expectedError        ve.ValidationError
	expectedToBail       bool
}

func ptr[T any](v T) *T {
	return &v
}

func runRuleTestCase(t *testing.T, ttIdx int, tt *ruleTestCaseData) {
	t.Run(fmt.Sprintf("#%[1]d: for value %[2]T(%[2]v)", ttIdx, tt.value), func(t *testing.T) {
		// when
		newValue, err := tt.rule.Apply(context.Background(), tt.value, tt.data)

		// then
		if tt.expectedError == nil {
			require.NoError(t, err, "Rule is expected to not return error")
		} else {
			require.ErrorIs(t, tt.expectedError, err, "Rule returned unexpected error")
		}

		if tt.expectedNewValueFunc == nil {
			require.Equal(t, tt.expectedNewValue, newValue, "Rule returned unexpected value")
		} else {
			require.True(t, tt.expectedNewValueFunc(newValue), "Rule returned unexpected value")
		}

		if bailingRule, ok := tt.rule.(BailingRule); ok {
			if tt.expectedToBail {
				require.True(t, bailingRule.Bails(), "Rule is expected to bail")
			} else {
				require.False(t, bailingRule.Bails(), "Rule is expected to not bail")
			}
		} else {
			require.False(t, tt.expectedToBail, "Rule is expected to be bailing")
		}
	})
}

func runRuleBenchmark(b *testing.B, ttIdx int, tt *ruleTestCaseData) {
	b.Run(fmt.Sprintf("#%[1]d: for value %[2]T(%[2]v)", ttIdx, tt.value), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = tt.rule.Apply(context.Background(), tt.value, tt.data)
		}
	})
}

func newRuleMock(idx int) *ruleMock {
	return &ruleMock{
		Idx: idx,
	}
}

type ruleMock struct {
	Idx int
}

func (r *ruleMock) Apply(_ context.Context, _ any, _ any) (any, ve.ValidationError) {
	//TODO implement me
	panic("implement me")
}
