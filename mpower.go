package mpowergo

import (
	"os"
	"reflect"
)

func envOr(name, def string) string {
	s := os.Getenv(name)
	if s == "" {
		return def
	}
	return s
}

func get(structObj interface{}, fieldName string) string {
	reflectValue := reflect.ValueOf(structObj)
	field := reflect.Indirect(reflectValue.FieldByName(fieldName))
	return string(field.String())
}
