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

func Test_ForMap(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.Background()
		data = map[string]any{
			"value": fakerInstance.Int(),
			"slice": []int{
				fakerInstance.IntBetween(-100, -1),
				fakerInstance.IntBetween(0, 100),
				fakerInstance.IntBetween(1000, 9999),
			},
		}

		valueRuleMocks, valueValidationErrorsBag                 = setupValueMocks(ctx, ctrl, "value", data["value"], data)
		sliceRuleMocks, sliceValidationErrorsBag                 = setupValueMocks(ctx, ctrl, "slice", data["slice"], data)
		sliceElementsRuleMocks, sliceElementsValidationErrorsBag = setupSliceElementsMocks(ctx, ctrl, "slice", data["slice"], data)
	)

	// when
	errorsBag, err := ForMap(data, mergeMaps(
		RulesMap{
			"value": valueRuleMocks,
			"slice": sliceRuleMocks,

			"non-existing value": {
				vr.Required(),
			},
			"value.*": {
				vr.Required(),
			},
			"value.0": {
				vr.Required(),
			},
			"value.foo": {
				vr.Required(),
			},
			"slice.foo": {
				vr.Required(),
			},
		},
		sliceElementsRuleMocks,
	))

	// then
	require.NoError(t, err)
	require.Len(t, errorsBag, 10)

	// "value" assertions
	require.True(t, assertErrorsBagContainsErrorsFromBagForField(t, errorsBag, valueValidationErrorsBag, "value"))

	// "slice" assertions
	require.True(t, assertErrorsBagContainsErrorsFromBagForField(t, errorsBag, sliceValidationErrorsBag, "slice"))
	require.True(t, assertErrorsBagContainsErrorsFromBagForField(t, errorsBag, sliceElementsValidationErrorsBag, "slice.0"))
	require.True(t, assertErrorsBagContainsErrorsFromBagForField(t, errorsBag, sliceElementsValidationErrorsBag, "slice.1"))
	require.True(t, assertErrorsBagContainsErrorsFromBagForField(t, errorsBag, sliceElementsValidationErrorsBag, "slice.2"))

	// "non-existing value" assertions
	require.True(t, assertErrorsBagContainsErrorsForField(t, errorsBag, []ve.ValidationError{vr.NewRequiredValidationError()}, "non-existing value"))

	// "value.*" assertions
	require.True(t, assertErrorsBagContainsErrorsForField(t, errorsBag, []ve.ValidationError{vr.NewRequiredValidationError()}, "value.*"))

	// "value.0" assertions
	require.True(t, assertErrorsBagContainsErrorsForField(t, errorsBag, []ve.ValidationError{vr.NewRequiredValidationError()}, "value.0"))

	// "value.foo" assertions
	require.True(t, assertErrorsBagContainsErrorsForField(t, errorsBag, []ve.ValidationError{vr.NewRequiredValidationError()}, "value.foo"))

	// "slice.foo" assertions
	require.True(t, assertErrorsBagContainsErrorsForField(t, errorsBag, []ve.ValidationError{vr.NewRequiredValidationError()}, "slice.foo"))
}

func Test_ForMapWithContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.TODO()
		data = map[string]any{
			"value": fakerInstance.Int(),
			"slice": []int{
				fakerInstance.IntBetween(-100, -1),
				fakerInstance.IntBetween(0, 100),
				fakerInstance.IntBetween(1000, 9999),
			},
		}

		valueRuleMocks, valueValidationErrorsBag                 = setupValueMocks(ctx, ctrl, "value", data["value"], data)
		sliceRuleMocks, sliceValidationErrorsBag                 = setupValueMocks(ctx, ctrl, "slice", data["slice"], data)
		sliceElementsRuleMocks, sliceElementsValidationErrorsBag = setupSliceElementsMocks(ctx, ctrl, "slice", data["slice"], data)
	)

	// when
	errorsBag, err := ForMapWithContext(ctx, data, mergeMaps(
		RulesMap{
			"value": valueRuleMocks,
			"slice": sliceRuleMocks,

			"non-existing value": {
				vr.Required(),
			},
			"value.*": {
				vr.Required(),
			},
			"value.0": {
				vr.Required(),
			},
			"value.foo": {
				vr.Required(),
			},
			"slice.foo": {
				vr.Required(),
			},
		},
		sliceElementsRuleMocks,
	))

	// then
	require.NoError(t, err)
	require.Len(t, errorsBag, 10)

	// "value" assertions
	require.True(t, assertErrorsBagContainsErrorsFromBagForField(t, errorsBag, valueValidationErrorsBag, "value"))

	// "slice" assertions
	require.True(t, assertErrorsBagContainsErrorsFromBagForField(t, errorsBag, sliceValidationErrorsBag, "slice"))
	require.True(t, assertErrorsBagContainsErrorsFromBagForField(t, errorsBag, sliceElementsValidationErrorsBag, "slice.0"))
	require.True(t, assertErrorsBagContainsErrorsFromBagForField(t, errorsBag, sliceElementsValidationErrorsBag, "slice.1"))
	require.True(t, assertErrorsBagContainsErrorsFromBagForField(t, errorsBag, sliceElementsValidationErrorsBag, "slice.2"))

	// "non-existing value" assertions
	require.True(t, assertErrorsBagContainsErrorsForField(t, errorsBag, []ve.ValidationError{vr.NewRequiredValidationError()}, "non-existing value"))

	// "value.*" assertions
	require.True(t, assertErrorsBagContainsErrorsForField(t, errorsBag, []ve.ValidationError{vr.NewRequiredValidationError()}, "value.*"))

	// "value.0" assertions
	require.True(t, assertErrorsBagContainsErrorsForField(t, errorsBag, []ve.ValidationError{vr.NewRequiredValidationError()}, "value.0"))

	// "value.foo" assertions
	require.True(t, assertErrorsBagContainsErrorsForField(t, errorsBag, []ve.ValidationError{vr.NewRequiredValidationError()}, "value.foo"))

	// "slice.foo" assertions
	require.True(t, assertErrorsBagContainsErrorsForField(t, errorsBag, []ve.ValidationError{vr.NewRequiredValidationError()}, "slice.foo"))
}

func Test_ForMapWithContext_ReturnsErrorFromOption(t *testing.T) {
	// given
	var (
		ctx  = context.TODO()
		data = map[string]any{
			"value": fakerInstance.Int(),
			"slice": []int{
				fakerInstance.IntBetween(-100, -1),
				fakerInstance.IntBetween(0, 100),
				fakerInstance.IntBetween(1000, 9999),
			},
		}

		errFromOption = errors.New(fakerInstance.Lorem().Sentence(6))
	)

	// when
	errorsBag, err := ForMapWithContext(ctx, data, nil, func(_ *validatorOptions) error {
		return errFromOption
	})

	// then
	require.ErrorIs(t, err, errFromOption)
	require.Nil(t, errorsBag)
}

func Test_ForMapWithContext_WithDataCollector(t *testing.T) {
	// given
	var (
		ctx  = context.TODO()
		data = map[string]any{
			"value": fakerInstance.Int(),
			"slice": []int{
				fakerInstance.IntBetween(-100, -1),
				fakerInstance.IntBetween(0, 100),
				fakerInstance.IntBetween(1000, 9999),
			},
		}
	)

	t.Run("value is not assigned when validation error occurs", func(t *testing.T) {
		collector := NewMapDataCollector()

		// when
		errorsBag, err := ForMapWithContext(ctx, data, RulesMap{
			"value": {
				vr.Required(),
				vr.Custom(func(_ context.Context, _ any, _ any) (any, error) {
					return nil, errors.New("validation failed")
				}),
			},
		}, ForMapWithDataCollector(collector))

		// then
		require.NoError(t, err)
		require.True(t, errorsBag.Has("value"))
		require.True(t, assertCollectorDoesNotHaveKey(t, collector, "value"))
		require.True(t, assertCollectorDoesNotHaveKey(t, collector, "slice"))
	})

	t.Run("changed value is assigned after successful validation", func(t *testing.T) {
		var newValue = fakerInstance.Int()

		collector := NewMapDataCollector()

		// when
		errorsBag, err := ForMapWithContext(ctx, data, RulesMap{
			"value": {
				vr.Required(),
				vr.Custom(func(_ context.Context, _ int, _ any) (int, error) {
					return newValue, nil
				}),
			},
			"slice.*": {
				vr.Required(),
			},
		}, ForMapWithDataCollector(collector))

		// then
		require.NoError(t, err)
		require.False(t, errorsBag.Has("value"))
		require.True(t, assertCollectorHasValue(t, collector, "value", newValue))
		require.True(t, assertCollectorDoesNotHaveKey(t, collector, "slice"))
		require.True(t, assertCollectorHasValue(t, collector, "slice.0", data["slice"].([]int)[0]))
		require.True(t, assertCollectorHasValue(t, collector, "slice.1", data["slice"].([]int)[1]))
		require.True(t, assertCollectorHasValue(t, collector, "slice.2", data["slice"].([]int)[2]))
	})
}
