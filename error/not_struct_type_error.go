package error

type NotStructTypeError struct {
}

func (NotStructTypeError) Error() string {
	return "not a struct type"
}
