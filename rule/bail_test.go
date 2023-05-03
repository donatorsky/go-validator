package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bailRule(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		value any
	}{
		{value: nil},
		{value: 1},
		{value: 1.0},
		{value: true},
		{value: []int{}},
		{value: map[string]int{}},
	} {
		t.Run(fmt.Sprintf("Test data #%d", ttIdx), func(t *testing.T) {
			rule := Bail()

			// when
			newValue, err := rule.Apply(nil, tt.value, nil)

			require.True(t, rule.Bails())

			// then
			require.NoError(t, err)
			require.Nil(t, newValue)
			require.True(t, rule.Bails())
		})
	}
}
