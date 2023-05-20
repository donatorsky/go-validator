package validator

import "reflect"

type validatorOptions struct {
	dataCollector DataCollector
	valueExporter *reflect.Value
}
