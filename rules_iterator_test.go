package validator

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

func Test_rulesIterator_CanIterate(t *testing.T) {
	// given
	var (
		ruleMock1 = ruleMock{1}
		ruleMock2 = ruleMock{2}
		ruleMock3 = ruleMock{3}
	)

	for ttName, tt := range map[string]struct {
		rules         []vr.Rule
		expectedRules []vr.Rule
	}{
		"empty slice": {
			rules:         nil,
			expectedRules: []vr.Rule{},
		},
		"non-empty slice": {
			rules:         []vr.Rule{ruleMock1, ruleMock2, ruleMock3},
			expectedRules: []vr.Rule{ruleMock1, ruleMock2, ruleMock3},
		},
	} {
		t.Run(ttName, func(t *testing.T) {
			var rules = make([]vr.Rule, 0, len(tt.expectedRules))

			i := newRulesIterator(tt.rules)

			// when
			for i.Valid() {
				rules = append(rules, i.Current())

				i.Next(context.TODO(), nil, nil)
			}

			// then
			require.Equal(t, tt.expectedRules, rules)
		})
	}
}

func Test_rulesIterator_ReturnsNilForEmptyIterator(t *testing.T) {
	// given
	i := newRulesIterator(nil)

	// then
	require.Nil(t, i.Current())
	require.False(t, i.Valid())
	require.Equal(t, 0, i.index)

	// when
	i.Next(context.TODO(), nil, nil)

	// then
	require.Nil(t, i.Current())
	require.False(t, i.Valid())
	require.Equal(t, 0, i.index)
}

type ruleMock struct {
	id int
}

func (ruleMock) Apply(_ context.Context, _ any, _ any) (newValue any, err ve.ValidationError) {
	panic("unexpected method call")
}

func (ruleMock) Bails() bool {
	panic("unexpected method call")
}

func (r ruleMock) String() string {
	return fmt.Sprintf("ruleMock{id=%d}", r.id)
}
