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

type CompositeValidationError struct {
	errors []ValidationError
}

func (e CompositeValidationError) GetRule() string {
	return ""
}

func (e CompositeValidationError) Error() string {
	return ""
}

func (e *CompositeValidationError) Add(error ValidationError) {
	e.errors = append(e.errors, error)
}

func (e *CompositeValidationError) Errors() []ValidationError {
	return e.errors
}

func (e *CompositeValidationError) Empty() bool {
	return len(e.errors) == 0
}
