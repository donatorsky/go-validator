package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewMapDataCollector(t *testing.T) {
	// given
	var (
		keyDummy    = fakerInstance.Lorem().Sentence(3)
		value1Dummy = fakerInstance.Int()
		value2Dummy = fakerInstance.Float(5, -9999, 9999)
	)

	collector := NewMapDataCollector()

	// then
	require.False(t, collector.Has(keyDummy))
	require.Nil(t, collector.Get(keyDummy))

	// when
	collector.Set(keyDummy, value1Dummy)

	// then
	require.True(t, collector.Has(keyDummy))
	require.Equal(t, value1Dummy, collector.Get(keyDummy))

	// and when
	collector.Set(keyDummy, value2Dummy)

	// then
	require.True(t, collector.Has(keyDummy))
	require.Equal(t, value2Dummy, collector.Get(keyDummy))
}
