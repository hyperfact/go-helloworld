package functor

import (
	_ "fmt"
)

type Func struct {
}

type Callable interface {
	Call0(args ...interface{})
	Call1(args ...interface{}) interface{}
	CallN(args ...interface{}) []interface{}

	Currying(args ...interface{}) Callable
}
