package error

import (
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func Test_ValueExporterTypeMismatchError_Error(t *testing.T) {
	fakerInstance := faker.New()

	// given
	var (
		valueTypeDummy  = fakerInstance.Lorem().Sentence(3)
		targetTypeDummy = fakerInstance.Lorem().Sentence(3)

		err = ValueExporterTypeMismatchError{
			ValueType:  valueTypeDummy,
			TargetType: targetTypeDummy,
		}
	)

	// then
	require.EqualError(t, err, fmt.Sprintf("value of type %s is not assignable to the type of %s", valueTypeDummy, targetTypeDummy))
}
