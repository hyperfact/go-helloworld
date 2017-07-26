package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func testTypeinfo() {
	testFunc()
	testStruct()
	testStructByIndex()
	testStructPtr()
	testArray()
	testString()
	testSlice()
	testChan()
	testInterface()
	testMap()
	testStructOf()
	testFuncOf()
}

func Func(p1 int, p2 string) error {
	return nil
}

func funcFoo(p1 float32, p2 string) error {
	return nil
}

func testFunc() {
	fmt.Println("----------------------------------------")
	fmt.Println("inner func relect: Func")
	inspectType(reflect.TypeOf(Func), 0)

	fmt.Println("----------------------------------------")
	fmt.Println("inner nonexport func relect: funcFoo")
	inspectType(reflect.TypeOf(funcFoo), 0)

	fmt.Println("----------------------------------------")
	fmt.Println("packaged func relect: fmt.Printf")
	inspectType(reflect.TypeOf(fmt.Printf), 0)

	val := reflect.ValueOf(Func)

	_ = val
}

type structFoo struct {
	f1 int
	f2 string
}

func (s *structFoo) Method1(a1 int, a2 string) error {
	return nil
}
func (s *structFoo) method1(a1 int, a2 string) error {
	return nil
}
func (s structFoo) Method2(a1 int, a2 string) error {
	return nil
}
func (s structFoo) method2(a1 int, a2 string) error {
	return nil
}

func testStruct() {
	fmt.Println("----------------------------------------")
	fmt.Println("inner struct relect: structFoo")
	var s structFoo
	inspectType(reflect.TypeOf(s), 0)

	fmt.Println("----------------------------------------")
	fmt.Println("packaged struct relect: bytes.Buffer")
	var b bytes.Buffer
	inspectType(reflect.TypeOf(b), 0)
}

func testStructByIndex() {
	type ST struct {
		A struct {
			Aa int
			Ab struct {
				Aba int
				Abb string
			}
			Ac string
		}
		B float32
		C string
		D []byte
	}
	fmt.Println("----------------------------------------")
	fmt.Println("struct byindex{0, 1, 1}} relect: ST")
	var s ST
	inspectType(reflect.TypeOf(s).FieldByIndex([]int{0, 1, 1}).Type, 0)
}

func testStructPtr() {
	fmt.Println("----------------------------------------")
	fmt.Println("inner struct ptr relect: structFoo")
	var s structFoo
	inspectType(reflect.TypeOf(&s), 0)

	fmt.Println("----------------------------------------")
	fmt.Println("packaged struct ptr relect: bytes.Buffer")
	var b bytes.Buffer
	inspectType(reflect.TypeOf(&b), 0)
}

func testArray() {
	fmt.Println("----------------------------------------")
	fmt.Println("array relect: arr")
	var arr [10]int
	inspectType(reflect.TypeOf(arr), 0)
}

func testString() {
	fmt.Println("----------------------------------------")
	fmt.Println("string relect: str")
	var str string
	inspectType(reflect.TypeOf(str), 0)
}

func testSlice() {
	fmt.Println("----------------------------------------")
	fmt.Println("slice relect: s")
	var s []int
	inspectType(reflect.TypeOf(s), 0)
}

func testChan() {
	fmt.Println("----------------------------------------")
	fmt.Println("chan relect: ch")
	var ch chan int
	inspectType(reflect.TypeOf(ch), 0)

	fmt.Println("----------------------------------------")
	fmt.Println("chan send relect: ch1")
	var ch1 <-chan int
	inspectType(reflect.TypeOf(ch1), 0)

	fmt.Println("----------------------------------------")
	fmt.Println("chan recv relect: ch2")
	var ch2 chan<- int
	inspectType(reflect.TypeOf(ch2), 0)
}

type Itfc interface {
	Method1(a1 int, a2 string) error
	Method2(a1 int, a2 string)
}

type ItfcImpl struct{}

func (i *ItfcImpl) Method1(a1 int, a2 string) error {
	return nil
}
func (i *ItfcImpl) Method2(a1 int, a2 string) {
}

func testInterface() {
	fmt.Println("----------------------------------------")
	fmt.Println("interface relect: i1")
	var i1 Itfc
	// 下面这段代码返回的type信息是interface的动态类型的信息
	// inspectType(reflect.TypeOf(i1), 0)
	inspectType(reflect.TypeOf(&i1).Elem(), 0)

	fmt.Println("----------------------------------------")
	fmt.Println("interface relect: (*Itfc)(nil)")
	inspectType(reflect.TypeOf((*Itfc)(nil)).Elem(), 0)

	fmt.Println("----------------------------------------")
	fmt.Println("interface relect: i2")
	var i2 io.ReadWriter
	inspectType(reflect.TypeOf(&i2).Elem(), 0)
}

func testMap() {
	fmt.Println("----------------------------------------")
	fmt.Println("map relect: m")
	var m map[int]string
	inspectType(reflect.TypeOf(m), 0)
}

func testStructOf() {
	fmt.Println("----------------------------------------")
	fmt.Println("StructOf: st")
	sfs := []reflect.StructField{
		{
			Name: "F1",
			Type: reflect.TypeOf(0),
		},
		{
			Name: "F2",
			Type: reflect.TypeOf(""),
		},
	}
	st := reflect.StructOf(sfs)
	inspectType(st, 0)
}

func testFuncOf() {
	fmt.Println("----------------------------------------")
	fmt.Println("FuncOf: f")

	in := []reflect.Type{
		reflect.TypeOf(0),
		reflect.TypeOf(""),
	}
	out := []reflect.Type{
		reflect.TypeOf(1),
		reflect.TypeOf(false),
	}
	f := reflect.FuncOf(in, out, false)
	inspectType(f, 0)
}
