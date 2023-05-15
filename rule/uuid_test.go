package rule

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UuidRule(t *testing.T) {
	runRuleTestCases(t, uuidRuleDataProvider)
}

func Test_UuidValidationError(t *testing.T) {
	// when
	err := NewUuidValidationError()

	// then
	require.EqualError(t, err, "must be a valid UUID")
}

func BenchmarkUuidRule(b *testing.B) {
	runRuleBenchmarks(b, uuidRuleDataProvider)
}

func uuidRuleDataProvider() map[string]*ruleTestCaseData {
	var (
		uuidV1Dummy      = "239eca5c-f348-11ed-8832-325096b39f47"
		uuidV3Dummy      = "7a5ab747-9269-3d14-bc16-d843452c9442"
		uuidV4Dummy      = fakerInstance.UUID().V4()
		uuidV5Dummy      = "deb74e35-ea5f-535f-890f-5779b5d8e27f"
		invalidUuidDummy = "invalid-uuid"
	)

	return map[string]*ruleTestCaseData{
		"nil": {
			rule:             UUID(),
			value:            nil,
			expectedNewValue: nil,
			expectedError:    nil,
		},

		"string with nil UUID": {
			rule:             UUID(),
			value:            nilUUID,
			expectedNewValue: nilUUID,
			expectedError:    nil,
		},
		"*string with nil UUID": {
			rule:             UUID(),
			value:            ptr(nilUUID),
			expectedNewValue: ptr(nilUUID),
			expectedError:    nil,
		},
		"string with nil UUID, nil UUD not allowed": {
			rule:             UUID(UUIDRuleDisallowNilUUID()),
			value:            nilUUID,
			expectedNewValue: nilUUID,
			expectedError:    NewUuidValidationError(),
		},

		"string with valid UUIDv1": {
			rule:             UUID(),
			value:            uuidV1Dummy,
			expectedNewValue: uuidV1Dummy,
			expectedError:    nil,
		},
		"*string with valid UUIDv1": {
			rule:             UUID(),
			value:            &uuidV1Dummy,
			expectedNewValue: &uuidV1Dummy,
			expectedError:    nil,
		},
		"string with valid UUIDv1, want UUIDv1": {
			rule:             UUID(UUIDRuleVersion1()),
			value:            uuidV1Dummy,
			expectedNewValue: uuidV1Dummy,
			expectedError:    nil,
		},
		"string with valid UUIDv1, want UUIDv3": {
			rule:             UUID(UUIDRuleVersion3()),
			value:            uuidV1Dummy,
			expectedNewValue: uuidV1Dummy,
			expectedError:    NewUuidValidationError(),
		},

		"string with valid UUIDv3": {
			rule:             UUID(),
			value:            uuidV3Dummy,
			expectedNewValue: uuidV3Dummy,
			expectedError:    nil,
		},
		"*string with valid UUIDv3": {
			rule:             UUID(),
			value:            &uuidV3Dummy,
			expectedNewValue: &uuidV3Dummy,
			expectedError:    nil,
		},
		"string with valid UUIDv3, want UUIDv3": {
			rule:             UUID(UUIDRuleVersion3()),
			value:            uuidV3Dummy,
			expectedNewValue: uuidV3Dummy,
			expectedError:    nil,
		},
		"string with valid UUIDv3, want UUIDv4": {
			rule:             UUID(UUIDRuleVersion4()),
			value:            uuidV3Dummy,
			expectedNewValue: uuidV3Dummy,
			expectedError:    NewUuidValidationError(),
		},

		"string with valid UUIDv4": {
			rule:             UUID(),
			value:            uuidV4Dummy,
			expectedNewValue: uuidV4Dummy,
			expectedError:    nil,
		},
		"*string with valid UUIDv4": {
			rule:             UUID(),
			value:            &uuidV4Dummy,
			expectedNewValue: &uuidV4Dummy,
			expectedError:    nil,
		},
		"string with valid UUIDv4, want UUIDv4": {
			rule:             UUID(UUIDRuleVersion4()),
			value:            uuidV4Dummy,
			expectedNewValue: uuidV4Dummy,
			expectedError:    nil,
		},
		"string with valid UUIDv4, want UUIDv5": {
			rule:             UUID(UUIDRuleVersion5()),
			value:            uuidV4Dummy,
			expectedNewValue: uuidV4Dummy,
			expectedError:    NewUuidValidationError(),
		},

		"string with valid UUIDv5": {
			rule:             UUID(),
			value:            uuidV5Dummy,
			expectedNewValue: uuidV5Dummy,
			expectedError:    nil,
		},
		"*string with valid UUIDv5": {
			rule:             UUID(),
			value:            &uuidV5Dummy,
			expectedNewValue: &uuidV5Dummy,
			expectedError:    nil,
		},
		"string with valid UUIDv5, want UUIDv5": {
			rule:             UUID(UUIDRuleVersion5()),
			value:            uuidV5Dummy,
			expectedNewValue: uuidV5Dummy,
			expectedError:    nil,
		},
		"string with valid UUIDv5, want UUIDv1": {
			rule:             UUID(UUIDRuleVersion1()),
			value:            uuidV5Dummy,
			expectedNewValue: uuidV5Dummy,
			expectedError:    NewUuidValidationError(),
		},

		"string with invalid UUID": {
			rule:             UUID(),
			value:            invalidUuidDummy,
			expectedNewValue: invalidUuidDummy,
			expectedError:    NewUuidValidationError(),
		},
		"*string with invalid UUID": {
			rule:             UUID(),
			value:            &invalidUuidDummy,
			expectedNewValue: &invalidUuidDummy,
			expectedError:    NewUuidValidationError(),
		},

		// unsupported values
		"int": {
			rule:             UUID(),
			value:            0,
			expectedNewValue: 0,
			expectedError:    NewUuidValidationError(),
		},
		"float": {
			rule:             UUID(),
			value:            0.0,
			expectedNewValue: 0.0,
			expectedError:    NewUuidValidationError(),
		},
		"complex": {
			rule:             UUID(),
			value:            1 + 2i,
			expectedNewValue: 1 + 2i,
			expectedError:    NewUuidValidationError(),
		},
		"bool": {
			rule:             UUID(),
			value:            true,
			expectedNewValue: true,
			expectedError:    NewUuidValidationError(),
		},
		"slice": {
			rule:             UUID(),
			value:            []int{},
			expectedNewValue: []int{},
			expectedError:    NewUuidValidationError(),
		},
		"array": {
			rule:             UUID(),
			value:            [1]int{},
			expectedNewValue: [1]int{},
			expectedError:    NewUuidValidationError(),
		},
		"map": {
			rule:             UUID(),
			value:            map[any]any{},
			expectedNewValue: map[any]any{},
			expectedError:    NewUuidValidationError(),
		},
		"struct": {
			rule:             UUID(),
			value:            someStruct{},
			expectedNewValue: someStruct{},
			expectedError:    NewUuidValidationError(),
		},
	}
}
