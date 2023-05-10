//go:generate mockgen -destination=validator_gomock_test.go -package=validator -source ./validator_test.go -mock_names=bailingRule=MockBailingRule bailingRule
//go:generate mockgen -destination=rule_gomock_test.go -package=validator github.com/donatorsky/go-validator/rule Rule

package validator

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

var fakerInstance = faker.New()

func TestForMap(t *testing.T) {
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

func TestForMapWithContext(t *testing.T) {
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

type someRequest struct {
	Value int   `validation:"value"`
	Slice []int `validation:"slice"`
}

func TestForStruct(t *testing.T) {
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

func TestForStruct_FailsWhenInvalidDataProvided(t *testing.T) {
	// when
	_, err := ForStruct(1, nil)

	// then
	require.ErrorIs(t, err, ErrNotStructType)
}

func TestForStructWithContext(t *testing.T) {
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

func TestForStructWithContext_FailsWhenInvalidDataProvided(t *testing.T) {
	// when
	_, err := ForStructWithContext(context.TODO(), 1, nil)

	// then
	require.ErrorIs(t, err, ErrNotStructType)
}

func TestForSlice(t *testing.T) {
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
	validationErrors, err := ForSlice(data, valueRuleMocks["_.*"]...)

	// then
	require.NoError(t, err)
	require.Len(t, validationErrors, 3)
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["0"], "_.0"))
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["1"], "_.1"))
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["2"], "_.2"))
}

func TestForSlice_FailsWhenInvalidDataProvided(t *testing.T) {
	// when
	_, err := ForSlice(1)

	// then
	require.ErrorIs(t, err, ErrNotListType)
}

func TestForSliceWithContext(t *testing.T) {
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
	validationErrors, err := ForSliceWithContext(ctx, data, valueRuleMocks["_.*"]...)

	// then
	require.NoError(t, err)
	require.Len(t, validationErrors, 3)
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["0"], "_.0"))
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["1"], "_.1"))
	require.True(t, assertErrorsBagContainsErrorsForField(t, sliceElementsValidationErrorsBag, validationErrors["2"], "_.2"))
}

func TestForSliceWithContext_FailsWhenInvalidDataProvided(t *testing.T) {
	// when
	_, err := ForSliceWithContext(context.TODO(), 1)

	// then
	require.ErrorIs(t, err, ErrNotListType)
}

func TestForValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.Background()
		data = fakerInstance.Int()

		valueRuleMocks, valueValidationErrorsBag = setupValueMocks(ctx, ctrl, "*", data, data)
	)

	// when
	validationErrors, err := ForValue(data, valueRuleMocks...)

	// then
	require.NoError(t, err)
	require.Len(t, validationErrors, 2)
	require.Equal(t, valueValidationErrorsBag.Get("*"), validationErrors)
}

func TestForValueWithContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.TODO()
		data = fakerInstance.Int()

		valueRuleMocks, valueValidationErrorsBag = setupValueMocks(ctx, ctrl, "*", data, data)
	)

	// when
	validationErrors, err := ForValueWithContext(ctx, data, valueRuleMocks...)

	// then
	require.NoError(t, err)
	require.Len(t, validationErrors, 2)
	require.Equal(t, valueValidationErrorsBag.Get("*"), validationErrors)
}

func setupValueMocks(ctx context.Context, ctrl *gomock.Controller, name string, value any, data any) ([]vr.Rule, ve.ErrorsBag) {
	var (
		errorsBag = ve.NewErrorsBag()

		valueRule1Mock     = NewMockRule(ctrl)
		valueRule1NewValue = fakerInstance.UInt()

		valueRule2Mock      = NewMockBailingRule(ctrl)
		valueRule2NewValue  = fakerInstance.Boolean().Bool()
		valueRule2MockError = newValidationErrorMock("valueRule2Mock", "error from valueRule2Mock")

		valueRule3Mock     = NewMockRule(ctrl)
		valueRule3NewValue = fakerInstance.Lorem().Sentence(5)

		valueRule4Mock      = NewMockBailingRule(ctrl)
		valueRule4NewValue  = fakerInstance.Float(5, -100, 100)
		valueRule4MockError = newValidationErrorMock("valueRule4Mock", "error from valueRule4Mock")

		valueRule5Mock = NewMockRule(ctrl)
	)

	valueRule1Mock.EXPECT().
		Apply(ctx, value, data).
		Times(1).
		Return(valueRule1NewValue, nil)

	valueRule2Mock.EXPECT().
		Apply(ctx, valueRule1NewValue, data).
		Times(1).
		Return(valueRule2NewValue, valueRule2MockError)
	valueRule2Mock.EXPECT().
		Bails().
		Times(1).
		Return(false)

	valueRule3Mock.EXPECT().
		Apply(ctx, valueRule2NewValue, data).
		Times(1).
		Return(valueRule3NewValue, nil)

	valueRule4Mock.EXPECT().
		Apply(ctx, valueRule3NewValue, data).
		Times(1).
		Return(valueRule4NewValue, valueRule4MockError)
	valueRule4Mock.EXPECT().
		Bails().
		Times(1).
		Return(true)

	errorsBag.Add(name, valueRule2MockError, valueRule4MockError)

	return []vr.Rule{
			valueRule1Mock,
			valueRule2Mock,
			valueRule3Mock,
			valueRule4Mock,
			valueRule5Mock,
		},
		errorsBag
}

func setupSliceElementsMocks(ctx context.Context, ctrl *gomock.Controller, name string, value any, data any) (RulesMap, ve.ErrorsBag) {
	eb := ve.NewErrorsBag()

	var (
		elementRule1Mock = NewMockRule(ctrl)
		elementRule2Mock = NewMockBailingRule(ctrl)
		elementRule3Mock = NewMockRule(ctrl)
		elementRule4Mock = NewMockBailingRule(ctrl)
		elementRule5Mock = NewMockRule(ctrl)
	)

	valueOf := reflect.ValueOf(value)
	for idx := 0; idx < valueOf.Len(); idx++ {
		var (
			elementRule1NewValue = fakerInstance.UInt()

			elementRule2NewValue  = fakerInstance.Boolean().Bool()
			elementRule2MockError = newValidationErrorMock(
				fmt.Sprintf("elementRule2Mock.%d", idx),
				fmt.Sprintf("error from elementRule2Mock.%d", idx),
			)

			elementRule3NewValue = fakerInstance.Lorem().Sentence(5)

			elementRule4NewValue  = fakerInstance.Float(5, -100, 100)
			elementRule4MockError = newValidationErrorMock(
				fmt.Sprintf("elementRule4Mock.%d", idx),
				fmt.Sprintf("error from elementRule4Mock.%d", idx),
			)
		)

		elementRule1Mock.EXPECT().
			Apply(ctx, valueOf.Index(idx).Interface(), data).
			Times(1).
			Return(elementRule1NewValue, nil)

		elementRule2Mock.EXPECT().
			Apply(ctx, elementRule1NewValue, data).
			Times(1).
			Return(elementRule2NewValue, elementRule2MockError)
		elementRule2Mock.EXPECT().
			Bails().
			Times(1).
			Return(false)

		elementRule3Mock.EXPECT().
			Apply(ctx, elementRule2NewValue, data).
			Times(1).
			Return(elementRule3NewValue, nil)

		elementRule4Mock.EXPECT().
			Apply(ctx, elementRule3NewValue, data).
			Times(1).
			Return(elementRule4NewValue, elementRule4MockError)
		elementRule4Mock.EXPECT().
			Bails().
			Times(1).
			Return(true)

		eb.Add(fmt.Sprintf("%s.%d", name, idx), elementRule2MockError, elementRule4MockError)
	}

	return RulesMap{
		fmt.Sprintf("%s.*", name): {
			elementRule1Mock,
			elementRule2Mock,
			elementRule3Mock,
			elementRule4Mock,
			elementRule5Mock,
		},
	}, eb
}

func mergeMaps(maps ...RulesMap) (merged RulesMap) {
	if len(maps) == 0 {
		return nil
	}

	merged = maps[0]

	for _, rulesMap := range maps[1:] {
		for k, v := range rulesMap {
			merged[k] = v
		}
	}

	return
}

func newValidationErrorMock(ruleName, errorValue string) *validationErrorMock {
	return &validationErrorMock{
		rule:  ruleName,
		error: errorValue,
	}
}

type bailingRule interface {
	vr.Rule
	vr.BailingRule
}

type validationErrorMock struct {
	rule  string
	error string
}

func (v validationErrorMock) GetRule() string {
	return v.rule
}

func (v validationErrorMock) Error() string {
	return v.error
}

func assertErrorsBagContainsErrorsForField(
	t *testing.T,
	errorsBag ve.ErrorsBag,
	errors []ve.ValidationError,
	field string,
) bool {
	return assert.True(t, errorsBag.Has(field), "Field is missing") &&
		assert.Equal(t, errors, errorsBag.Get(field), "Errors does not match")
}

func assertErrorsBagContainsErrorsFromBagForField(
	t *testing.T,
	errorsBag ve.ErrorsBag,
	valueValidationErrorsBag ve.ErrorsBag,
	field string,
) bool {
	return assertErrorsBagContainsErrorsForField(t, errorsBag, valueValidationErrorsBag.Get(field), field)
}
