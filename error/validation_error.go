//go:generate mockgen -destination=validation_error_gomock_test.go -package=error -source ./validation_error.go ValidationError

package error

type ValidationError interface {
	GetRule() string
	Error() string
}

type BasicValidationError struct {
	Rule string `json:"rule"`
}

func (e BasicValidationError) GetRule() string {
	return e.Rule
}

func NewCustomMessageValidationError(rule, message string) *CustomMessageValidationError {
	return &CustomMessageValidationError{
		BasicValidationError: BasicValidationError{
			Rule: rule,
		},
		Message: message,
	}
}

type CustomMessageValidationError struct {
	BasicValidationError

	Message string `json:"message"`
}

func (c CustomMessageValidationError) Error() string {
	return c.Message
}
