package main

import (
	"fmt"
)

type IA interface {
	methodA()
}

type AImpl struct{}

func (a *AImpl) methodA() {
	var va interface{} = a
	b, ok := va.(IB)
	fmt.Printf("va.(IB) = %v,%v \n", b, ok)
}

type IB interface {
	methodB()
}

type BImpl struct{}

func (*BImpl) methodB() {}

type C struct {
	IA
	*BImpl
}

func (*C) methodC() {}

func main() {
	var c interface{} = &C{
		IA:    new(AImpl),
		BImpl: new(BImpl),
	}

	pc, ok := c.(*C)
	fmt.Printf("c.(*C) = %v,%v \n", pc, ok)
	ibimpl, ok := c.(*BImpl)
	fmt.Printf("c.(*BImpl) = %v,%v \n", ibimpl, ok)
	ib, ok := c.(IB)
	fmt.Printf("c.(IB) = %v,%v \n", ib, ok)
	ibimpl, ok = ib.(*BImpl)
	fmt.Printf("ib.(*BImpl) = %v,%v \n", ibimpl, ok)
	ia, ok := c.(IA)
	fmt.Printf("c.(IA) = %v,%v \n", ia, ok)
	iaimpl, ok := c.(*AImpl)
	fmt.Printf("c.(*AImpl) = %v,%v \n", iaimpl, ok)
	pc.methodA()
}
