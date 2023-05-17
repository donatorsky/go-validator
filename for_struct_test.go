package validator

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

type someRequest struct {
	Value int   `validation:"value"`
	Slice []int `validation:"slice"`
}

func Test_ForStruct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.Background()
		data = someRequest{
			Value: fakerInstance.Int(),
			Slice: []int{
				fakerInstance.IntBetween(-100, -1),
				fakerInstance.IntBetween(0, 100),
				fakerInstance.IntBetween(1000, 9999),
			},
		}

		valueRuleMocks, valueValidationErrorsBag                 = setupValueMocks(ctx, ctrl, "value", data.Value, data)
		sliceRuleMocks, sliceValidationErrorsBag                 = setupValueMocks(ctx, ctrl, "slice", data.Slice, data)
		sliceElementsRuleMocks, sliceElementsValidationErrorsBag = setupSliceElementsMocks(ctx, ctrl, "slice", data.Slice, data)
	)

	// when
	errorsBag, err := ForStruct(data, mergeMaps(
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

func Test_ForStruct_FailsWhenInvalidDataProvided(t *testing.T) {
	// when
	_, err := ForStruct(1, nil)

	// then
	require.ErrorIs(t, err, ve.ErrNotStructType)
}

func Test_ForStructWithContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.TODO()
		data = someRequest{
			Value: fakerInstance.Int(),
			Slice: []int{
				fakerInstance.IntBetween(-100, -1),
				fakerInstance.IntBetween(0, 100),
				fakerInstance.IntBetween(1000, 9999),
			},
		}

		valueRuleMocks, valueValidationErrorsBag                 = setupValueMocks(ctx, ctrl, "value", data.Value, data)
		sliceRuleMocks, sliceValidationErrorsBag                 = setupValueMocks(ctx, ctrl, "slice", data.Slice, data)
		sliceElementsRuleMocks, sliceElementsValidationErrorsBag = setupSliceElementsMocks(ctx, ctrl, "slice", data.Slice, data)
	)

	// when
	errorsBag, err := ForStructWithContext(ctx, data, mergeMaps(
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

func Test_ForStructWithContext_FailsWhenInvalidDataProvided(t *testing.T) {
	// when
	_, err := ForStructWithContext(context.TODO(), 1, nil)

	// then
	require.ErrorIs(t, err, ve.ErrNotStructType)
}
