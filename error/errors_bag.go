package error

import "fmt"

func NewErrorsBag() ErrorsBag {
	return make(ErrorsBag)
}

type ErrorsBag map[string][]ValidationError

func (b ErrorsBag) Add(field string, message ValidationError) {
	if _, exists := b[field]; !exists {
		b[field] = []ValidationError{message}
	} else {
		b[field] = append(b[field], message)
	}
}

func (b ErrorsBag) Any() bool {
	return len(b) > 0
}

func (b ErrorsBag) All() map[string][]ValidationError {
	return b
}

func (b ErrorsBag) Has(field string) bool {
	_, exists := b[field]

	return exists
}

func (b ErrorsBag) Get(field string) []ValidationError {
	if errors, exists := b[field]; exists {
		return errors
	}

	return nil
}

func (b ErrorsBag) Error() string {
	message := fmt.Sprintf("%d field(s) failed:", len(b))

	for field, errors := range b {
		message += fmt.Sprintf("\n%s: [%d]%s", field, len(errors), errors)
	}

	return message
}
