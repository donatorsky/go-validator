//go:generate mockgen -destination=helpers_gomock_test.go -package=validator -source ./helpers_test.go -mock_names=bailingRule=MockBailingRule bailingRule
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

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

var fakerInstance = faker.New()

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

func assertCollectorDoesNotHaveKey(t *testing.T, collector DataCollector, key string) bool {
	return assert.False(t, collector.Has(key), fmt.Sprintf("Data collector is expected to not have %q key", key))
}

func assertCollectorHasValue(t *testing.T, collector DataCollector, key string, value any) bool {
	return assert.True(t, collector.Has(key), fmt.Sprintf("Data collector is expected to have %q key", key)) &&
		assert.Equal(t, value, collector.Get(key), fmt.Sprintf("Data collector %q key is expected to be %v", key, value))
}
