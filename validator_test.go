//go:generate mockgen -destination=validator_gomock_test.go -package=validator -source ./validator_test.go -mock_names=bailingRule=MockBailingRule bailingRule
//go:generate mockgen -destination=rule_gomock_test.go -package=validator github.com/donatorsky/go-validator/rule Rule

package validator

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	ve "github.com/donatorsky/go-validator/error"
	vr "github.com/donatorsky/go-validator/rule"
)

func TestForMapWithContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.Background()
		data = map[string]any{
			"value": 100,
		}
		valueRule1Mock      = NewMockRule(ctrl)
		valueRule2Mock      = NewMockBailingRule(ctrl)
		valueRule2MockError = NewValidationErrorMock("valueRule2Mock", "error from valueRule2Mock")
		valueRule3Mock      = NewMockRule(ctrl)
		valueRule4Mock      = NewMockBailingRule(ctrl)
		valueRule4MockError = NewValidationErrorMock("valueRule4Mock", "error from valueRule4Mock")
		valueRule5Mock      = NewMockRule(ctrl)

		rulesMap = RulesMap{
			"value": {
				valueRule1Mock,
				valueRule2Mock,
				valueRule3Mock,
				valueRule4Mock,
				valueRule5Mock,
			},
		}
	)

	valueRule1Mock.EXPECT().
		Apply(ctx, 100, data).
		Times(1).
		Return(200, nil)

	valueRule2Mock.EXPECT().
		Apply(ctx, 200, data).
		Times(1).
		Return(300, valueRule2MockError)
	valueRule2Mock.EXPECT().
		Bails().
		Times(1).
		Return(false)

	valueRule3Mock.EXPECT().
		Apply(ctx, 300, data).
		Times(1).
		Return(400, nil)

	valueRule4Mock.EXPECT().
		Apply(ctx, 400, data).
		Times(1).
		Return(500, valueRule4MockError)
	valueRule4Mock.EXPECT().
		Bails().
		Times(1).
		Return(true)

	// when
	errorsBag := ForMapWithContext(ctx, data, rulesMap)

	// then
	fmt.Println(errorsBag)
	require.Len(t, errorsBag, 1)

	// "value" assertions
	require.True(t, errorsBag.Has("value"))
	require.Equal(t, []ve.ValidationError{valueRule2MockError, valueRule4MockError}, errorsBag.Get("value"))
}

func NewValidationErrorMock(ruleName, errorValue string) *validationErrorMock {
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
