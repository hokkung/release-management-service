package gorem

import (
	"reflect"
	"strings"
)

func getPrimaryKey[T any]() string {
	t := reflect.TypeOf((*T)(nil)).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag, ok := field.Tag.Lookup("gorm"); ok && strings.Contains(tag, "primaryKey") {
			return field.Name
		}
	}
	return "id"
}
