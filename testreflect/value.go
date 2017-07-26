package main

import (
	"fmt"
	"reflect"
	"sort"
)

func testValue() {
	testFuncCall()
	testSetValue()
	testMakeFunc()
	testMethodValue()
}

func printValues(v []reflect.Value) {
	for _, i := range v {
		if i.CanInterface() {
			fmt.Printf("%v,", i.Interface())
		}
	}
}

func testFuncCall() {
	fmt.Println("----------------------------------------")
	fmt.Println("relect value func : foo")

	foo := func(i int, s string) error {
		fmt.Printf("i:%v,s:%v\n", i, s)
		return fmt.Errorf("error")
	}

	vFoo := reflect.ValueOf(foo)
	i := 123
	s := "abc"
	r := vFoo.Call([]reflect.Value{reflect.ValueOf(i), reflect.ValueOf(s)})
	printValues(r)
	fmt.Println()
}

func testSetValue() {
	fmt.Println("----------------------------------------")
	var v int
	fmt.Printf("v can set:%v | can address:%v\n", reflect.ValueOf(v).CanSet(), reflect.ValueOf(v).CanAddr())

	pv := &v
	fmt.Printf("pv can set:%v | can address:%v\n", reflect.ValueOf(pv).CanSet(), reflect.ValueOf(pv).CanAddr())
	fmt.Printf("*pv can set:%v | can address:%v\n", reflect.ValueOf(*pv).CanSet(), reflect.ValueOf(*pv).CanAddr())
	fmt.Printf("pv.Elem() can set:%v | can address:%v\n", reflect.ValueOf(pv).Elem().CanSet(), reflect.ValueOf(pv).Elem().CanAddr())

	var ar = [...]int{1, 2, 3}
	fmt.Printf("ar[0] can set: %v | can address:%v\n", reflect.ValueOf(ar[0]).CanSet(), reflect.ValueOf(ar[0]).CanAddr())
	fmt.Printf("ar.Index(0) can set: %v | can address:%v\n", reflect.ValueOf(ar).Index(0).CanSet(), reflect.ValueOf(ar).Index(0).CanAddr())

	par := &ar
	fmt.Printf("par[0] can set: %v | can address:%v\n", reflect.ValueOf((*par)[0]).CanSet(), reflect.ValueOf((*par)[0]).CanAddr())
	fmt.Printf("par.Index(0) can set: %v | can address:%v\n", reflect.ValueOf(par).Elem().Index(0).CanSet(), reflect.ValueOf(par).Elem().Index(0).CanAddr())

	var sli = []int{1, 2, 3}
	fmt.Printf("sli[0] can set: %v | can address:%v\n", reflect.ValueOf(sli[0]).CanSet(), reflect.ValueOf(sli).CanAddr())
	fmt.Printf("sli.Index(0) can set: %v | can address:%v\n", reflect.ValueOf(sli).Index(0).CanSet(), reflect.ValueOf(sli).Index(0).CanAddr())

	var st = struct {
		F1 int
		f2 string
	}{}
	fmt.Printf("st can set: %v | can address:%v\n", reflect.ValueOf(st).CanSet(), reflect.ValueOf(st).CanAddr())
	fmt.Printf("st.Field(0) can set: %v | can address:%v\n", reflect.ValueOf(st).Field(0).CanSet(), reflect.ValueOf(st).Field(0).CanAddr())
	fmt.Printf("st.Field(1) can set: %v | can address:%v\n", reflect.ValueOf(st).Field(1).CanSet(), reflect.ValueOf(st).Field(1).CanAddr())

	pst := &st
	fmt.Printf("pst[0] can set: %v | can address:%v\n", reflect.ValueOf(*pst).CanSet(), reflect.ValueOf(*pst).CanAddr())
	fmt.Printf("pst.Field(0) can set: %v | can address:%v\n", reflect.ValueOf(pst).Elem().Field(0).CanSet(), reflect.ValueOf(pst).Elem().Field(0).CanAddr())
	fmt.Printf("pst.Field(1) can set: %v | can address:%v\n", reflect.ValueOf(pst).Elem().Field(1).CanSet(), reflect.ValueOf(pst).Elem().Field(1).CanAddr())
}

type ValueSlice []reflect.Value

func (s ValueSlice) Len() int           { return len(s) }
func (s ValueSlice) Less(i, j int) bool { return i > j }
func (s ValueSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func testMakeFunc() {

	sort.Reverse(nil)
	reverse := func(in []reflect.Value) []reflect.Value {
		sort.Sort(ValueSlice(in))
		return in
	}

	makeFunc := func(fptr interface{}, fimpl func([]reflect.Value) []reflect.Value) {
		fmt.Println("----------------------------------------")
		fmt.Println("relect value MakeFunc : reverse")
		// fptr is a pointer to a function.
		// Obtain the function value itself (likely nil) as a reflect.Value
		// so that we can query its type and then set the value.
		fn := reflect.ValueOf(fptr).Elem()

		// Make a function of the right type.
		v := reflect.MakeFunc(fn.Type(), fimpl)

		// Assign it to the value fn represents.
		fn.Set(v)
	}

	var fi func(int, int, int) (int, int, int)
	makeFunc(&fi, reverse)

	fmt.Println(fi(1, 2, 3))
}

type ST1 struct {
}

func (*ST1) Method(i int, s string) error {
	return nil
}

func testMethodValue() {
	fmt.Println("----------------------------------------")
	fmt.Println("relect method value : st.Method")
	var st ST1
	valM := reflect.ValueOf(st.Method)
	tyM := valM.Type()

	inspectType(tyM, 0)
	_ = valM
	_ = tyM
	fmt.Println()
}
