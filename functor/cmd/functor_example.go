package main

import (
	"fmt"
	"reflect"
)

func fun1(i int) {
	fmt.Printf("fun1(%v)\n", i)
}

func fun2(s string) {
	fmt.Printf("fun2(%q)\n", s)

}

func fun3(i int, s string) {
	fmt.Printf("fun3(%v,%q)\n", i, s)
}

type funT func(a interface{})

func Solution1() {
	var f interface{}
	var a1 interface{}
	var a2 interface{}

	f = fun1
	a1 = 1
	reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(a1)})

	f = fun2
	a1 = "abc"
	reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(a1)})

	f = fun3
	a1 = 1
	a2 = "abc"
	reflect.ValueOf(f).Call([]reflect.Value{
		reflect.ValueOf(a1),
		reflect.ValueOf(a2),
	})

	_ = f
}

type callable interface {
}

func Solution2() {
	var f callable
	_ = f
}

func Solution3() {
	f1 := func(args ...interface{}) {

	}
	f2 := func(args ...interface{}) {

	}
	f3 := func(args ...interface{}) {

	}

	var f func(...interface{})
	f = f1
	f = f2
	f = f3

	f(1)
	f("string")
	f(1, "string")
	_ = f
}

func main() {
	Solution1()

	Solution2()

	Solution3()
}
