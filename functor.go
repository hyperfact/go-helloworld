package main

func fun0() {

}

func fun1(int) {

}

func fun2(string) {

}

func fun3(int, string) {

}

func Currying(f func(...interface{}), args ...interface{}) func() {
	return func() { f(args...) }
}

var Functor func()

func testFunctor() {
	Println()
}
