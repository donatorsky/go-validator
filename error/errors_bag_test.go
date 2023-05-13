package error

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func Test_NewErrorsBag(t *testing.T) {
	var fakerInstance = faker.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	var (
		existingField1Dummy   = fmt.Sprintf("existing1.%s", fakerInstance.Lorem().Word())
		existingField2Dummy   = fmt.Sprintf("existing2.%s", fakerInstance.Lorem().Word())
		nonExistingFieldDummy = fmt.Sprintf("non-existent.%s", fakerInstance.Lorem().Word())
	)

	errorsBag := NewErrorsBag()

	// then
	require.False(t, errorsBag.Any())
	require.False(t, errorsBag.Has(nonExistingFieldDummy))
	require.Nil(t, errorsBag.Get(nonExistingFieldDummy))
	require.False(t, errorsBag.Has(existingField1Dummy))
	require.Nil(t, errorsBag.Get(existingField1Dummy))
	require.False(t, errorsBag.Has(existingField2Dummy))
	require.Nil(t, errorsBag.Get(existingField2Dummy))
	require.Empty(t, errorsBag.All())
	require.EqualError(t, errorsBag, "0 field(s) failed:")

	// and given
	var (
		error1Mock         = NewMockValidationError(ctrl)
		error2Mock         = NewMockValidationError(ctrl)
		error3Mock         = NewMockValidationError(ctrl)
		error1MessageDummy = fakerInstance.Lorem().Sentence(6)
		error2MessageDummy = fakerInstance.Lorem().Sentence(6)
		error3MessageDummy = fakerInstance.Lorem().Sentence(6)
	)

	error1Mock.EXPECT().
		Error().
		Times(1).
		Return(error1MessageDummy)

	error2Mock.EXPECT().
		Error().
		Times(1).
		Return(error2MessageDummy)

	error3Mock.EXPECT().
		Error().
		Times(1).
		Return(error3MessageDummy)

	// when
	errorsBag.Add(existingField1Dummy, error1Mock)
	errorsBag.Add(existingField2Dummy, error2Mock, error3Mock)

	// then
	require.True(t, errorsBag.Any())
	require.False(t, errorsBag.Has(nonExistingFieldDummy))
	require.Nil(t, errorsBag.Get(nonExistingFieldDummy))
	require.True(t, errorsBag.Has(existingField1Dummy))
	require.Equal(t, []ValidationError{error1Mock}, errorsBag.Get(existingField1Dummy))
	require.True(t, errorsBag.Has(existingField2Dummy))
	require.Equal(t, []ValidationError{error2Mock, error3Mock}, errorsBag.Get(existingField2Dummy))
	require.Equal(t, map[string][]ValidationError{
		existingField1Dummy: {error1Mock},
		existingField2Dummy: {error2Mock, error3Mock},
	}, errorsBag.All())

	errorMessage := errorsBag.Error()
	require.Contains(t, errorMessage, "2 field(s) failed:")
	require.Contains(t, errorMessage, fmt.Sprintf(
		"%s: [1]{%s}",
		existingField1Dummy,
		error1MessageDummy,
	))
	require.Contains(t, errorMessage, fmt.Sprintf(
		"%s: [2]{%s; %s}",
		existingField2Dummy,
		error2MessageDummy,
		error3MessageDummy,
	))

	// and given
	var (
		error4Mock         = NewMockValidationError(ctrl)
		error4MessageDummy = fakerInstance.Lorem().Sentence(6)
	)

	error1Mock.EXPECT().
		Error().
		Times(1).
		Return(error1MessageDummy)

	error2Mock.EXPECT().
		Error().
		Times(1).
		Return(error2MessageDummy)

	error3Mock.EXPECT().
		Error().
		Times(1).
		Return(error3MessageDummy)

	error4Mock.EXPECT().
		Error().
		Times(1).
		Return(error4MessageDummy)

	// when
	errorsBag.Add(existingField1Dummy, error4Mock)

	// then
	require.True(t, errorsBag.Any())
	require.False(t, errorsBag.Has(nonExistingFieldDummy))
	require.Nil(t, errorsBag.Get(nonExistingFieldDummy))
	require.True(t, errorsBag.Has(existingField1Dummy))
	require.Equal(t, []ValidationError{error1Mock, error4Mock}, errorsBag.Get(existingField1Dummy))
	require.True(t, errorsBag.Has(existingField2Dummy))
	require.Equal(t, []ValidationError{error2Mock, error3Mock}, errorsBag.Get(existingField2Dummy))
	require.Equal(t, map[string][]ValidationError{
		existingField1Dummy: {error1Mock, error4Mock},
		existingField2Dummy: {error2Mock, error3Mock},
	}, errorsBag.All())

	errorMessage = errorsBag.Error()
	require.Contains(t, errorMessage, "2 field(s) failed:")
	require.Contains(t, errorMessage, fmt.Sprintf(
		"%s: [2]{%s; %s}",
		existingField1Dummy,
		error1MessageDummy,
		error4MessageDummy,
	))
	require.Contains(t, errorMessage, fmt.Sprintf(
		"%s: [2]{%s; %s}",
		existingField2Dummy,
		error2MessageDummy,
		error3MessageDummy,
	))
}
