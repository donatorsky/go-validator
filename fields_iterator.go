package validator

import (
	"reflect"
	"strconv"
	"strings"

	vr "github.com/donatorsky/go-validator/rule"
)

func newFieldsIterator(field string, data any) <-chan fieldValue {
	fieldsValues := make(chan fieldValue, 1)

	go func() {
		defer close(fieldsValues)

		fieldParts := strings.Split(field, ".")

		iterateOverFieldPart(fieldsValues, make([]string, len(fieldParts)), fieldParts, 0, data)
	}()

	return fieldsValues
}

type fieldValue struct {
	field string
	value any
}

func iterateOverFieldPart(fieldsValues chan<- fieldValue, fieldName []string, fieldParts []string, position int, value any) {
	isNil := false
	if value == nil {
		isNil = true
	} else {
		value, isNil = vr.Dereference(value)
	}

	if isNil {
		for idx := position; idx < len(fieldParts); idx++ {
			fieldName[idx] = fieldParts[idx]
		}

		fieldsValues <- fieldValue{
			field: strings.Join(fieldName, "."),
			value: nil,
		}

		return
	}

	if len(fieldParts) == position {
		fieldsValues <- fieldValue{
			field: strings.Join(fieldName, "."),
			value: value,
		}

		return
	}

	valueOf := reflect.ValueOf(value)

	if fieldParts[position] == "*" {
		if valueOf.Kind() == reflect.Slice || valueOf.Kind() == reflect.Array {
			for idx := 0; idx < valueOf.Len(); idx++ {
				fieldName[position] = strconv.Itoa(idx)

				iterateOverFieldPart(fieldsValues, fieldName, fieldParts, position+1, valueOf.Index(idx).Interface())
			}

			return
		}

		for idx := position; idx < len(fieldParts); idx++ {
			fieldName[idx] = fieldParts[idx]
		}

		fieldsValues <- fieldValue{
			value: nil,
			field: strings.Join(fieldName, "."),
		}

		return
	}

	switch valueOf.Kind() {
	case reflect.Map:
		mapIndex := valueOf.MapIndex(reflect.ValueOf(fieldParts[position]))
		if mapIndex.IsValid() {
			value = mapIndex.Interface()
		} else {
			value = nil
		}

	case reflect.Struct:
		fieldByName := valueOf.FieldByName(fieldParts[position])
		if fieldByName.IsValid() {
			value = fieldByName.Interface()
		} else {
			typeOf := reflect.TypeOf(value)
			value = nil
			for idx := 0; idx < typeOf.NumField(); idx++ {
				structField := typeOf.Field(idx)
				nameFromTag := structField.Tag.Get("validation")
				if nameFromTag != "" && nameFromTag == fieldParts[position] {
					value = valueOf.FieldByName(structField.Name).Interface()
				}
			}
		}

	case reflect.Slice, reflect.Array:
		idx, err := strconv.Atoi(fieldParts[position])
		if err != nil || idx < 0 || idx >= valueOf.Len() {
			value = nil

			break
		}

		value = valueOf.Index(idx).Interface()

	default:
		value = nil
	}

	fieldName[position] = fieldParts[position]

	iterateOverFieldPart(fieldsValues, fieldName, fieldParts, position+1, value)
}
