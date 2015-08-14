package mpower

import (
	"os"
	"reflect"
)

// Get an environment variable or return `def` string as the default
//
//    str := envOr("MP-Master-Key", "43545-54545-454532-3232")
func envOr(name, def string) string {
	s := os.Getenv(name)
	if s == "" {
		return def
	}
	return s
}

// Get a value from a struct by using its field value through reflection
//    Setup is a struct
//    func (setup *Setup) Get(fieldName string) string {
//        return get(setup, fieldName)
//    }
func get(structObj interface{}, fieldName string) string {
	reflectValue := reflect.ValueOf(structObj)
	field := reflect.Indirect(reflectValue).FieldByName(fieldName)
	return string(field.String())
}
