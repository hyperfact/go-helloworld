package main

import "C"
import "fmt"

//export goCallBack
func goCallBack(a C.int) {
	fmt.Println("goCallBack:", int(a))
}

//export GoCallBack
func GoCallBack(a *C.char) {
	fmt.Println("GoCallBack:", C.GoString(a))
}
