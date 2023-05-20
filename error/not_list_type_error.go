package error

type NotListTypeError struct {
}

func (NotListTypeError) Error() string {
	return "not an array or a slice type"
}
