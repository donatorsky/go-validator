package rule

import (
	"context"
	"net"

	ve "github.com/donatorsky/go-validator/error"
)

func IP() *ipRule {
	return &ipRule{}
}

type ipRule struct {
}

func (*ipRule) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
	v, isNil := Dereference(value)
	if isNil {
		return value, nil
	}

	if stringValue, ok := v.(string); !ok || net.ParseIP(stringValue) == nil {
		return value, NewIpValidationError()
	}

	return value, nil
}

func NewIpValidationError() IpValidationError {
	return IpValidationError{
		BasicValidationError: ve.BasicValidationError{
			Rule: ve.RuleIP,
		},
	}
}

type IpValidationError struct {
	ve.BasicValidationError
}

func (e IpValidationError) Error() string {
	return "must be a valid IP address"
}
