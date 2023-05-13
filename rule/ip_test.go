package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IpRule(t *testing.T) {
	runRuleTestCases(t, ipRuleDataProvider)
}

func Test_IpValidationError(t *testing.T) {
	// when
	err := NewIpValidationError()

	// then
	require.EqualError(t, err, "must be a valid IP address")
}

func BenchmarkIpRule(b *testing.B) {
	runRuleBenchmarks(b, ipRuleDataProvider)
}

func ipRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		ipv4Dummy      = fakerInstance.Internet().Ipv4()
		ipv6Dummy      = fakerInstance.Internet().Ipv6()
		invalidIpDummy = fakerInstance.Lorem().Sentence(6)
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             IP(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"valid IPv4 string": {
			rule:             IP(),
			value:            ipv4Dummy,
			expectedNewValue: ipv4Dummy,
			expectedError:    nil,
		},
		"valid IPv4 *string": {
			rule:             IP(),
			value:            ptr(ipv4Dummy),
			expectedNewValue: ptr(ipv4Dummy),
			expectedError:    nil,
		},

		"valid IPv6 string": {
			rule:             IP(),
			value:            ipv6Dummy,
			expectedNewValue: ipv6Dummy,
			expectedError:    nil,
		},
		"valid IPv6 *string": {
			rule:             IP(),
			value:            ptr(ipv6Dummy),
			expectedNewValue: ptr(ipv6Dummy),
			expectedError:    nil,
		},

		"invalid IP address": {
			rule:             IP(),
			value:            invalidIpDummy,
			expectedNewValue: invalidIpDummy,
			expectedError:    NewIpValidationError(),
		},

		// unsupported values
		"int": {
			rule:             IP(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewIpValidationError(),
		},
		"float": {
			rule:             IP(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewIpValidationError(),
		},
		"complex": {
			rule:             IP(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewIpValidationError(),
		},
		"bool": {
			rule:             IP(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewIpValidationError(),
		},
		"slice": {
			rule:             IP(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewIpValidationError(),
		},
		"array": {
			rule:             IP(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewIpValidationError(),
		},
		"map": {
			rule:             IP(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewIpValidationError(),
		},
		"struct": {
			rule:             IP(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewIpValidationError(),
		},
	}
}
