package reflectutil

import (
	"reflect"
)

type info struct {
	typ   string
	value reflect.Value
}

func Dir(i interface{}) []string {
	v := reflect.ValueOf(i)

	dir(v)
	return nil
}

func Vars(i interface{}) {

}

func dir(v reflect.Value) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	case reflect.String:
	}
}
