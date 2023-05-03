package validator

import (
	"fmt"
	"reflect"
	"strings"

	vr "github.com/donatorsky/go-validator/rule"
)

func newFieldsIterator(field string, data any) <-chan fieldValue {
	fieldsValues := make(chan fieldValue, 1)

	go func() {
		defer close(fieldsValues)

		fieldParts := strings.Split(field, ".")

		iterateOverFieldPart(fieldsValues, "", fieldParts, data, false)
	}()

	return fieldsValues
}

type fieldValue struct {
	pattern bool
	field   string
	value   any
	isNil   bool
}

func iterateOverFieldPart(fieldsValues chan<- fieldValue, fieldName string, fieldParts []string, value any, pattern bool) {
	if len(fieldParts) == 0 {
		fv := fieldValue{
			pattern: pattern,
			field:   fieldName,
		}

		fv.value, fv.isNil = vr.Dereference(value)

		fieldsValues <- fv

		return
	}

	if fieldParts[0] == "*" {
		if valueOf := reflect.ValueOf(value); valueOf.Kind() == reflect.Slice || valueOf.Kind() == reflect.Array {
			for idx := 0; idx < valueOf.Len(); idx++ {
				var subFieldName string
				if len(fieldName) == 0 {
					subFieldName = fmt.Sprintf("%d", idx)
				} else {
					subFieldName = fmt.Sprintf("%s.%d", fieldName, idx)
				}

				iterateOverFieldPart(fieldsValues, subFieldName, fieldParts[1:], valueOf.Index(idx).Interface(), true)
			}
		} else {
			fv := fieldValue{
				pattern: true,
			}

			remainingFields := strings.Join(fieldParts[1:], ".")
			if len(remainingFields) != 0 {
				remainingFields = fmt.Sprintf(".%s", remainingFields)
			}

			if len(fieldName) == 0 {
				fv.field = fmt.Sprintf("*%s", remainingFields)
			} else {
				fv.field = fmt.Sprintf("%s.*%s", fieldName, remainingFields)
			}

			fv.value, fv.isNil = vr.Dereference(value)

			fieldsValues <- fv
		}

		return
	}

	switch valueOf := reflect.ValueOf(value); valueOf.Kind() {
	case reflect.Map:
		mapIndex := valueOf.MapIndex(reflect.ValueOf(fieldParts[0]))
		if mapIndex.IsValid() {
			value = mapIndex.Interface()
		} else {
			value = nil
		}

	case reflect.Struct:
		fieldByName := valueOf.FieldByName(fieldParts[0])
		if fieldByName.IsValid() {
			value = fieldByName.Interface()
		} else {
			typeOf := reflect.TypeOf(value)
			value = nil
			for idx := 0; idx < typeOf.NumField(); idx++ {
				structField := typeOf.Field(idx)
				nameFromTag := structField.Tag.Get("validation")
				if nameFromTag != "" && nameFromTag == fieldParts[0] {
					value = valueOf.FieldByName(structField.Name).Interface()
				}
			}
		}

	default:
		value = nil
	}

	if len(fieldName) == 0 {
		fieldName = fieldParts[0]
	} else if len(fieldName) > 0 && len(fieldParts) == 1 {
		fieldName += "." + fieldParts[0]
	}

	iterateOverFieldPart(fieldsValues, fieldName, fieldParts[1:], value, pattern)
}
