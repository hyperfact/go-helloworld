package main

import (
	"flag"
	"fmt"
	"helloworld/pkg01"
	pkg02_01 "helloworld/pkg02/pkg01"
	"runtime/debug"

	_ "helloworld/pkg01"

	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

const APP_VERSION = "0.1"

type Interface interface{}

type Info struct {
	mem int
}

var Println = fmt.Println
var Printf = fmt.Printf

const Pi float32 = 3.1415926

func Fun1(i Interface) {
	fmt.Println(i)
}

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func testReflect() {
	//t := reflect.TypeOf(Fun1)

	t := reflect.TypeOf(Fun1)
	Println(t)

	b := 10
	Println(reflect.TypeOf(b))

	//Println(Info)
	a := new(int)
	_ = a
}

func testFuncSE() {
	defer func() func() {
		Println("func start")
		return func() {
			Println("func end")
		}
	}()()

	Println("func doing...")
	fmt.Printf("%d", time.Second)

	//time.Sleep(3 * time.Second)
}

func testBlockDefer() {
	for i := 0; i < 3; i++ {
		Printf("loop: %v\n", i)
		func() {
			defer func() {
				Printf("loop defer: %v\n", i)
			}()
		}()

		defer func() {
			Printf("loop defer func: %v\n", i)
		}()
	}
}

type ITest interface {
}

type Pointer *int

type TestMethod struct {
	PMethod  func()
	PMethod1 func(int)
}

func (t *TestMethod) Method() {

}

func (t TestMethod) Method2() {

}

func FunMethod(t *TestMethod) {

}

func testMethod() {
	t := &TestMethod{}
	f := (*TestMethod).Method
	f1 := t.Method
	f2 := TestMethod.Method2
	f3 := t.Method2
	f4 := FunMethod
	t.PMethod = t.Method

	Printf("%T\n", f)
	Printf("%T\n", f1)
	Printf("%T\n", f2)
	Printf("%T\n", f3)
	Printf("%T\n", f4)
	Printf("%T\n", t.PMethod)

	t.Method()
	f(t)
	f1()
	f2(*t)
	f3()
	f4(t)
	t.PMethod()
}

func testEmptyInterface() {
	var i interface{}
	i = 1
	Printf("type of i: %T\n", i)
	i = "123456"
	Printf("type of i: %T\n", i)
	i = new(int)
	Printf("type of i: %T\n", i)
	i = []int{}
	Printf("type of i: %T\n", i)
	i = func() {}
	Printf("type of i: %T\n", i)
	i = struct{}{}
	Printf("type of i: %T\n", i)
	i = Info{}
	Printf("type of i: %T\n", i)
	_ = i
}

func testBareReturn() (a int, err error) {
	return
}

func srcFileLine(skip int) (file string, line int) {
	_, file, line, _ = runtime.Caller(skip)
	return
}

func srcFileLineString() string {
	f, l := srcFileLine(2)
	return fmt.Sprintf("%v:%v", f, l)
}

func SprintStack() string {
	buf := make([]byte, 1024)
	nbytes := runtime.Stack(buf, false)
	Printf("stack size:%v\n", nbytes)
	return string(buf[:nbytes])
}

func testPanicStackTrace() {
	defer func() {
		if e := recover(); e != nil {
			Println(e)
			Println(string(debug.Stack()))
		}
	}()

	panic(fmt.Errorf(srcFileLineString()))
}

func testStackTrace() {
	skipframe := 2
	s := string(debug.Stack())
	const linesep = '\n'
	for i := 0; i < skipframe; i++ {
		idx := strings.IndexRune(s, linesep)
		if idx < 0 {
			break
		}
		s = s[idx+1:]
	}
}

func testChan() {
	ch := make(chan int)
	Printf("chan type: %T\n", ch)

	go func() {
		ch <- <-ch
	}()

	ch <- 12
	Printf("ch:%v\n", <-ch)
}

func Sync(l sync.Locker, f func()) {
	l.Lock()
	defer l.Unlock()
	f()
}

func testSync() {
	ctr := 0
	var mu sync.Mutex
	abort := make(chan bool)
	_ = mu
	_ = ctr

	Println("start testSync")

	go func() {
		Println("add routine...")
		run := true
		for run {
			Sync(&mu, func() {
				if ctr < 10 {
					ctr++
				} else {
					run = false
				}
			})
			time.Sleep(1 * time.Second)
		}
		close(abort)
		Println("end add routine")
	}()

	go func() {
		Println("print routine...")

		t := time.Tick(1 * time.Second)

	lbl_loop:
		for {
			select {
			case <-t:
				Sync(&mu, func() {
					Println("critical section:", ctr)
				})
			case <-abort:
				break lbl_loop
			}
		}

		Println("end print routine")
	}()

	<-abort

	Println("end testSync\n")
}

type expr fmt.Stringer

type compare interface {
	Less() bool
}

type CmpInt int

func (i CmpInt) Less() bool {
	return true
}

func testCompare(c *compare) bool {
	return true
}

func testInterface() {
	var c CmpInt = 1
	var ic compare = c
	//testCompare(&c)
	testCompare(&ic)

	if _, ok := ic.(compare); ok {

	}
}

func testNestedStruct() interface{} {
	type nested struct {
		i int
		s string
	}
	return nested{}
}

func testTypeDef() {
	type t int
	var a t = 0
	Printf("type of a %T\n", a)
	Printf("reflect.ValueOf a kind %v\n", reflect.ValueOf(a).Kind())

	var i interface{} = a
	switch i.(type) {
	case int:
		Println("int")
	default:
		Println("other")
	}

	switch i.(type) {
	case t:
		Println("t")
	case int:
		Println("int")
	}
}

func testSelect() {
	Println("start a empty select")
	//select {} // will cause sleep forever
	Println("end a empty select")

	Println("start a default select")
	select {
	default:
	}
	Println("end a default select")

	Println("start a 1 case select")
	ch := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		Println("write chan")
		ch <- true
	}()
	select {
	case <-ch:
	}
	Println("end a 1 case select")

	Println("start a 1 case and default select")
	ch1 := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		Println("write chan")
		ch1 <- true
	}()
	select {
	case <-ch1:
	default:
	}
	Println("end a 1 case and default select")
}

func testPointerIntefaceSlice() {
	var a int
	var b float32
	var c string

	is := []interface{}{
		&a, &b, &c,
	}
	for _, v := range is {
		Printf("%v:%T\n", v, v)
	}
}

func testTypeSwitch() {
	type myint int
	var val myint
	var i interface{} = val
	switch i.(type) {
	case int:
		Println("i.(type) : int")
	case myint:
		Println("i.(type) : myint")
	}

	Println("kind of i:%v", reflect.TypeOf(i))
}

func testNestedStructConvert() {
	type A struct {
		f int
	}
	type B struct {
		g string
	}
	type C struct {
		A
		b B
	}

	var c interface{} = C{}
	ca, ok := c.(A)
	Printf("c.(A): %v, %v\n", ca, ok)

	//var cc C
	//caa := A(c)
	//_ = caa
}

type IA interface {
	fun1()
}
type IB interface {
	fun2()
}
type ID interface {
	fun3()
}
type C struct {
	IA
	B  IB
	p1 int
}

func (c *C) fun3() {}

func testInterfaceQuery() {
	var c interface{} = &C{}

	ca, ok := c.(IA)
	Printf("c.(IA): %v, %v\n", ca, ok)
	cb, ok := c.(IB)
	Printf("c.(IB): %v, %v\n", cb, ok)
	cc, ok := c.(*C)
	Printf("c.(C): %v, %v\n", cc, ok)
	cd, ok := ca.(ID)
	Printf("ca.(ID): %v, %v\n", cd, ok)
}

func testMultiSelect() {
	chs := make([]chan struct{}, 10)
	for i := 0; i < len(chs); i++ {
		chs[0] = make(chan struct{}, 1)
	}

	type pair struct {
		ch  interface{}
		val struct{}
	}

	aggr := make(chan pair, 10)
	var wg sync.WaitGroup
	wg.Add(len(chs))
	for _, ch := range chs {
		go func(c chan struct{}) {
			aggr <- pair{c, <-c}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(aggr)
	}()

	for p := range aggr {
		Println(p.ch, p.val)
	}
}

func testForIter() {
	ar := []int{1, 2, 3, 4, 5}
	for i := range ar {
		// wrong
		defer func() { Println(i) }()

		a := i
		// correct
		defer func() { Println(a) }()
	}
}

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	var expr = 10
	_ = expr

	Printf("type of nil %T\n", nil)
	Printf("type of nil %v\n", reflect.TypeOf(nil))

	testForIter()

	testInterfaceQuery()

	testNestedStructConvert()

	testTypeSwitch()

	testStackTrace()

	testPointerIntefaceSlice()

	testSelect()

	testTypeDef()

	testNestedStruct()

	testInterface()

	testSync()

	testChan()

	testFunctor()

	testMethod()

	testPanicStackTrace()

	testBlockDefer()

	testEmptyInterface()

	testFuncSE()

	testReflect()

	const A = 10
	Println(A)

	Fun1(A)

	pkg01.PkgFun()
	pkg02_01.PkgFun()

	m := map[string]bool{"abc": true}
	fmt.Printf("map len: %d\n", len(m))

	a, _ := m["123"]
	fmt.Printf("map len: %d\n", len(m))

	a = m["123"]
	fmt.Printf("map len: %d\n", len(m))

	m["123"] = true
	fmt.Printf("map len: %d\n", len(m))

	_ = a
}
