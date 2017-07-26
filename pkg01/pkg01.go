package pkg01

import (
	"fmt"
)

const (
	Mon = 1 + iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

func init() {

}

func PkgFun() {
	fmt.Println("PkgFun")

}

func PkgFun1() {

	a := [...]int{1, 2}
	//a = []int{1, 2, 3}
	fmt.Print(a[0])

	m := map[string]int{
		"Mon": 1,
		"Tue": 2,
		"Wed": 3,
		"Thu": 4,
		"Fri": 5,
		"Sat": 6,
		"Sun": 7,
	}
	fmt.Print(m)
}

func test() {

}

func init() {

}
