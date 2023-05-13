package error

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func Test_BasicValidationError_GetRule(t *testing.T) {
	var fakerInstance = faker.New()

	// given
	var ruleDummy = fakerInstance.Lorem().Word()

	bve := BasicValidationError{
		Rule: ruleDummy,
	}

	// when
	rule := bve.GetRule()

	// then
	require.Equal(t, ruleDummy, rule)
}

func TestNewCustomMessageValidationError(t *testing.T) {
	var fakerInstance = faker.New()

	// given
	var (
		ruleDummy    = fakerInstance.Lorem().Word()
		messageDummy = fakerInstance.Lorem().Sentence(6)
	)

	cve := NewCustomMessageValidationError(ruleDummy, messageDummy)

	// then
	require.EqualError(t, cve, messageDummy)
}
