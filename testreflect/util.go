package main

import (
	"fmt"
	"reflect"
)

func Print(indent int, a ...interface{}) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Println(a...)
}

func Printf(indent int, format string, a ...interface{}) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Printf(format, a...)
	fmt.Println()
}

func inspectType(t reflect.Type, indent int) {
	Print(indent, "String:", t.String())
	Print(indent, "Name:", t.Name())
	Print(indent, "PkgPath:", t.PkgPath())
	Print(indent, "Kind:", t.Kind())
	Print(indent, "Size:", t.Size())
	Print(indent, "Align:", t.Align())
	Print(indent, "FieldAlign:", t.FieldAlign())

	if reflect.Int <= t.Kind() && t.Kind() <= reflect.Complex128 {
		Print(indent, "Bits:", t.Bits())
	}

	Print(indent, "NumMethod:", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		Print(indent, "Method:", t.Method(i))
	}
	//fmt.Println("MethodByName:", t.MethodByName())

	Print(indent, "Comparable:", t.Comparable())
	//Print(indent, "Implements:", t.Implements(nil))
	//Print(indent, "AssignableTo:", t.AssignableTo(nil))
	//Print(indent, "ConvertibleTo:", t.ConvertibleTo(nil))

	switch t.Kind() {
	case reflect.Chan:
		Print(indent, "Elem:", t.Elem())
		Print(indent, "ChanDir:", t.ChanDir())
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			Print(indent, "Field:", t.Field(i))
		}
	case reflect.Func:
		Print(indent, "IsVariadic:", t.IsVariadic())
		for i := 0; i < t.NumIn(); i++ {
			Print(indent, "In:", t.In(i))
			//inspectType(t.In(i), indent+1)
		}
		for i := 0; i < t.NumOut(); i++ {
			Print(indent, "Out:", t.Out(i))
			//inspectType(t.Out(i), indent+1)
		}
	case reflect.Map:
		Print(indent, "Elem:", t.Elem())
		Print(indent, "Key:", t.Key())
	case reflect.Array:
		Print(indent, "Elem:", t.Elem())
		Print(indent, "Len:", t.Len())
	case reflect.Ptr:
		Print(indent, "Elem:", t.Elem())
	case reflect.Slice:
		Print(indent, "Elem:", t.Elem())
	}
}

func inspectValue(t reflect.Value, indent int) {

}
