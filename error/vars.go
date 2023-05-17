package error

import (
	"errors"
	"fmt"
)

var (
	ErrNotStructType                     = errors.New("not a struct type")
	ErrNotListType                       = errors.New("not an array or a slice type")
	ErrValueExporterPointerNotAssignable = errors.New("cannot assign value to the pointer")
)

type ValueExporterTypeMismatchError struct {
	ValueType  string
	TargetType string
}

func (e ValueExporterTypeMismatchError) Error() string {
	return fmt.Sprintf("value of type %s is not assignable to the type of %s", e.ValueType, e.TargetType)
}
