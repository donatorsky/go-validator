package error

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotListTypeError_Error(t *testing.T) {
	// given
	err := NotListTypeError{}

	// then
	require.EqualError(t, err, "not an array or a slice type")
}
