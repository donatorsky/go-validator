package error

import "fmt"

type ValueExporterTypeMismatchError struct {
	ValueType  string
	TargetType string
}

func (e ValueExporterTypeMismatchError) Error() string {
	return fmt.Sprintf("value of type %s is not assignable to the type of %s", e.ValueType, e.TargetType)
}
