package validator

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
)

func Test_ForSlice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.Background()
		data = []int{
			fakerInstance.IntBetween(-100, -1),
			fakerInstance.IntBetween(0, 100),
			fakerInstance.IntBetween(1000, 9999),
		}

		valueRuleMocks, sliceElementsValidationErrorsBag = setupSliceElementsMocks(ctx, ctrl, "_", data, data)
	)

	// when
	validationErrors, err := ForSlice(data, valueRuleMocks["_.*"])

	// then
	require.NoError(t, err)
	require.Len(t, validationErrors, 3)
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["0"], "_.0"))
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["1"], "_.1"))
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["2"], "_.2"))
}

func Test_ForSlice_FailsWhenInvalidDataProvided(t *testing.T) {
	// when
	_, err := ForSlice(1, nil)

	// then
	require.ErrorIs(t, err, ve.ErrNotListType)
}

func Test_ForSliceWithContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.TODO()
		data = []int{
			fakerInstance.IntBetween(-100, -1),
			fakerInstance.IntBetween(0, 100),
			fakerInstance.IntBetween(1000, 9999),
		}

		valueRuleMocks, sliceElementsValidationErrorsBag = setupSliceElementsMocks(ctx, ctrl, "_", data, data)
	)

	// when
	validationErrors, err := ForSliceWithContext(ctx, data, valueRuleMocks["_.*"])

	// then
	require.NoError(t, err)
	require.Len(t, validationErrors, 3)
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["0"], "_.0"))
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["1"], "_.1"))
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["2"], "_.2"))
}

func Test_ForSliceWithContext_FailsWhenInvalidDataProvided(t *testing.T) {
	// when
	_, err := ForSliceWithContext(context.TODO(), 1, nil)

	// then
	require.ErrorIs(t, err, ve.ErrNotListType)
}
