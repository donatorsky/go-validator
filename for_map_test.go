package validator

import (
	"context"
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
