package error

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotStructTypeError_Error(t *testing.T) {
	// given
	var err = NotStructTypeError{}

	// then
	require.EqualError(t, err, "not a struct type")
}
