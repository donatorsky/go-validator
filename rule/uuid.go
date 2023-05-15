package rule

import (
	"context"
	"regexp"

	ve "github.com/donatorsky/go-validator/error"
)

const nilUUID = "00000000-0000-0000-0000-000000000000"

var anyUUIDRegex = regexp.MustCompile(`^[[:xdigit:]]{8}-[[:xdigit:]]{4}-([1345])[[:xdigit:]]{3}-[[:xdigit:]]{4}-[[:xdigit:]]{12}$`)

type uuidRuleOption func(options *uuidRuleOptions)

func UUID(options ...uuidRuleOption) *uuidRule {
	opts := uuidRuleOptions{
		any:      true,
		allowNil: true,
	}

	for _, option := range options {
		option(&opts)
	}

	return &uuidRule{
		options: opts,
	}
}

type uuidRule struct {
	options uuidRuleOptions
}

func (r *uuidRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	stringValue, ok := v.(string)
	if !ok {
		return value, NewUuidValidationError()
	}

	if stringValue == nilUUID {
		if !r.options.allowNil {
			return value, NewUuidValidationError()
		}

		return value, nil
	}

	submatch := anyUUIDRegex.FindStringSubmatch(stringValue)
	if submatch == nil {
		return value, NewUuidValidationError()
	}

	if !r.options.any &&
		!(r.options.version1 && submatch[1] == "1") &&
		!(r.options.version3 && submatch[1] == "3") &&
		!(r.options.version4 && submatch[1] == "4") &&
		!(r.options.version5 && submatch[1] == "5") {
		return value, NewUuidValidationError()
	}

	return value, nil
}

func NewUuidValidationError() UuidValidationError {
	return UuidValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.TypeUUID,
		},
	}
}

type UuidValidationError struct {
	ve.BasicValidationError
}

func (e UuidValidationError) Error() string {
	return "must be a valid UUID"
}

type uuidRuleOptions struct {
	any      bool
	version1 bool
	version3 bool
	version4 bool
	version5 bool
	allowNil bool
}

func UUIDRuleVersion1() uuidRuleOption {
	return func(options *uuidRuleOptions) {
		options.any = false
		options.version1 = true
	}
}

func UUIDRuleVersion3() uuidRuleOption {
	return func(options *uuidRuleOptions) {
		options.any = false
		options.version3 = true
	}
}

func UUIDRuleVersion4() uuidRuleOption {
	return func(options *uuidRuleOptions) {
		options.any = false
		options.version4 = true
	}
}

func UUIDRuleVersion5() uuidRuleOption {
	return func(options *uuidRuleOptions) {
		options.any = false
		options.version5 = true
	}
}

func UUIDRuleDisallowNilUUID() uuidRuleOption {
	return func(options *uuidRuleOptions) {
		options.allowNil = false
	}
}
