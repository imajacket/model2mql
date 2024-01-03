package model2mql

import (
	"fmt"
	"reflect"
	"strings"
)

type Convertor[T any] struct {
	model T
}

func NewConvertor[T any](model T) *Convertor[T] {
	return &Convertor[T]{
		model: model,
	}
}

func (c *Convertor[T]) Parse(model T) (string, error) {
	var query []string

	val := reflect.ValueOf(model)
	typ := reflect.TypeOf(model)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		name := field.Name
		tag := field.Tag.Get("mql")
		fieldValue := val.Field(i)
		value := fieldValue.Interface()
		valueType := fieldValue.Type().String()

		var valueStr string
		if valueType == "string" {
			valueStr = fmt.Sprintf("\"%v\"", value)
		} else {
			valueStr = fmt.Sprintf("%v", value)
		}

		if strings.HasSuffix(name, "Contains") {
			query = append(query, fmt.Sprintf("%s %% %s", tag, valueStr))
			continue
		}

		if strings.HasSuffix(name, "Gt") {
			query = append(query, fmt.Sprintf("%s > %s", tag, valueStr))
			continue
		}

		if strings.HasSuffix(name, "Gte") {
			query = append(query, fmt.Sprintf("%s >= %s", tag, valueStr))
			continue
		}

		if strings.HasSuffix(name, "Lt") {
			query = append(query, fmt.Sprintf("%s < %s", tag, valueStr))
			continue
		}

		if strings.HasSuffix(name, "Lte") {
			query = append(query, fmt.Sprintf("%s <= %s", tag, valueStr))
			continue
		}

		if strings.HasSuffix(name, "Ne") {
			query = append(query, fmt.Sprintf("%s != %s", tag, valueStr))
			continue
		}

		query = append(query, fmt.Sprintf("%s = %s", tag, valueStr))
	}
	return strings.Join(query, " and "), nil
}
