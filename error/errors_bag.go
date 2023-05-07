package error

import (
	"fmt"
	"strings"
)

func NewErrorsBag() ErrorsBag {
	return make(ErrorsBag)
}

type ErrorsBag map[string][]ValidationError

func (b ErrorsBag) Add(field string, errors ...ValidationError) {
	if _, exists := b[field]; !exists {
		b[field] = errors
	} else {
		b[field] = append(b[field], errors...)
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
		messages := make([]string, len(errors))
		for idx, validationError := range errors {
			messages[idx] = validationError.Error()
		}

		message += fmt.Sprintf("\n%s: [%d]%s", field, len(errors), strings.Join(messages, "; "))
	}

	return message
}
