package main

/*
	# include <stdlib.h>
	# include <stdio.h>

	extern void goCallBack(int);
	extern void GoCallBack(const char *);
	void Print(const char *str) {
		printf("%s\n", str);
		goCallBack(1);
		GoCallBack("abc");
	}
*/
import "C"

import (
	"unsafe"
)

func main() {
	s := "hello world"
	cs := C.CString(s)
	C.Print(cs)

	C.free(unsafe.Pointer(cs))
}
