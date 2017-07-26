package main

import (
	"fmt"

	"github.com/smallnest/rpcx"
)

type Entity struct {
}

type Req struct {
	Damage int
}

type Rsp struct {
	RealDamage int
}

func (e *Entity) RecvDamage(req *Req, rsp *Rsp) error {
	rsp.RealDamage = int(float32(req.Damage) * 0.8)
	return nil
}

func main() {
	fmt.Print()

	//c := codec.NewGobServerCode

	svr := rpcx.NewServer()
	svr.RegisterName("Entity", new(Entity))
	if err := svr.Serve("tcp", ":6666"); err != nil {
		fmt.Println(err)
	}
}
