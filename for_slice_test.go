package validator

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
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
	require.ErrorIs(t, err, ve.NotListTypeError{})
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
	require.ErrorIs(t, err, ve.NotListTypeError{})
}

func Test_ForSliceWithContext_ReturnsErrorFromOption(t *testing.T) {
	// given
	var (
		ctx  = context.TODO()
		data = []int{
			fakerInstance.IntBetween(-100, -1),
			fakerInstance.IntBetween(0, 100),
			fakerInstance.IntBetween(1000, 9999),
		}

		errFromOption = errors.New(fakerInstance.Lorem().Sentence(6))
	)

	// when
	validationErrors, err := ForSliceWithContext(ctx, data, nil, func(_ *validatorOptions) error {
		return errFromOption
	})

	// then
	require.ErrorIs(t, err, errFromOption)
	require.Nil(t, validationErrors)
}

func Test_ForSliceWithContext_WithDataCollector(t *testing.T) {
	// given
	var (
		ctx  = context.TODO()
		data = []int{
			fakerInstance.IntBetween(-100, -1),
			fakerInstance.IntBetween(0, 100),
			fakerInstance.IntBetween(1000, 9999),
		}
	)

	t.Run("value is not assigned when validation error occurs", func(t *testing.T) {
		collector := NewMapDataCollector()

		// when
		errorsBag, err := ForSliceWithContext(ctx, data, []vr.Rule{
			vr.Required(),
			vr.Custom(func(_ context.Context, _ any, _ any) (any, error) {
				return nil, errors.New("validation failed")
			}),
		}, ForSliceWithDataCollector(collector))

		// then
		require.NoError(t, err)
		require.True(t, errorsBag.Has("0"))
		require.True(t, errorsBag.Has("1"))
		require.True(t, errorsBag.Has("2"))
		require.True(t, assertCollectorDoesNotHaveKey(t, collector, "0"))
		require.True(t, assertCollectorDoesNotHaveKey(t, collector, "1"))
		require.True(t, assertCollectorDoesNotHaveKey(t, collector, "2"))
	})

	t.Run("changed value is assigned after successful validation", func(t *testing.T) {
		var newValue = fakerInstance.Int()

		collector := NewMapDataCollector()

		// when
		errorsBag, err := ForSliceWithContext(ctx, data, []vr.Rule{
			vr.Required(),
			vr.Custom(func(_ context.Context, _ int, _ any) (int, error) {
				return newValue, nil
			}),
		}, ForSliceWithDataCollector(collector))

		// then
		require.NoError(t, err)
		require.False(t, errorsBag.Has("0"))
		require.False(t, errorsBag.Has("1"))
		require.False(t, errorsBag.Has("2"))
		require.True(t, assertCollectorHasValue(t, collector, "0", newValue))
		require.True(t, assertCollectorHasValue(t, collector, "1", newValue))
		require.True(t, assertCollectorHasValue(t, collector, "2", newValue))
	})
}
