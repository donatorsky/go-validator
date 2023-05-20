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

func Test_ForValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.Background()
		data = fakerInstance.Int()

		valueRuleMocks, valueValidationErrorsBag = setupValueMocks(ctx, ctrl, "*", data, data)
	)

	// when
	validationErrors, err := ForValue(data, valueRuleMocks)

	// then
	require.NoError(t, err)
	require.Len(t, validationErrors, 2)
	require.Equal(t, valueValidationErrorsBag.Get("*"), validationErrors)
}

func Test_ForValueWithContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		ctx  = context.TODO()
		data = fakerInstance.Int()

		valueRuleMocks, valueValidationErrorsBag = setupValueMocks(ctx, ctrl, "*", data, data)
	)

	// when
	validationErrors, err := ForValueWithContext(ctx, data, valueRuleMocks)

	// then
	require.NoError(t, err)
	require.Len(t, validationErrors, 2)
	require.Equal(t, valueValidationErrorsBag.Get("*"), validationErrors)
}

func Test_ForValueWithContext_ReturnsErrorFromOption(t *testing.T) {
	// given
	var errFromOption = errors.New(fakerInstance.Lorem().Sentence(6))

	// when
	validationErrors, err := ForValueWithContext[any](context.TODO(), nil, nil, func(_ *validatorOptions) error {
		return errFromOption
	})

	// then
	require.ErrorIs(t, err, errFromOption)
	require.Nil(t, validationErrors)
}

func Test_ForValueWithContext_WithValueExporter(t *testing.T) {
	// given
	var (
		ctx  = context.TODO()
		data = fakerInstance.IntBetween(-9999, 9999)
	)

	t.Run("value is not assigned when validation error occurs", func(t *testing.T) {
		var out = 10_000

		// when
		validationErrors, err := ForValueWithContext(
			ctx,
			data,
			[]vr.Rule{
				vr.Required(),
				vr.Custom(func(_ context.Context, _ any, _ any) (any, error) {
					return nil, errors.New("validation failed")
				}),
			},
			ForValueWithValueExporter(&out),
		)

		// then
		require.NoError(t, err)
		require.NotEmpty(t, validationErrors)
		require.Equal(t, 10_000, out)
	})

	t.Run("value is not assigned when output value type differs", func(t *testing.T) {
		var out int

		// when
		validationErrors, err := ForValueWithContext(
			ctx,
			data,
			[]vr.Rule{
				vr.Required(),
				vr.Custom(func(_ context.Context, _ int, _ any) (string, error) {
					return "", nil
				}),
			},
			ForValueWithValueExporter(&out),
		)

		// then
		require.ErrorIs(t, err, ve.ValueExporterTypeMismatchError{
			ValueType:  "string",
			TargetType: "int",
		})
		require.Empty(t, validationErrors)
	})

	t.Run("changed value is assigned after successful validation", func(t *testing.T) {
		var (
			newValue = fakerInstance.Int()
			out      int
		)

		// when
		validationErrors, err := ForValueWithContext(
			ctx,
			data,
			[]vr.Rule{
				vr.Required(),
				vr.Custom(func(_ context.Context, _ int, _ any) (int, error) {
					return newValue, nil
				}),
			},
			ForValueWithValueExporter(&out),
		)

		// then
		require.NoError(t, err)
		require.Empty(t, validationErrors)
		require.Equal(t, newValue, out)
	})
}
