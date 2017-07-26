package main

import (
	"fmt"
	"time"

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

func AsyncCall(cli *rpcx.Client) {
	req := &Req{100}
	rsp := Rsp{}
	chCall := cli.Go("Entity.RecvDamage", req, &rsp, nil)
	replyCall := <-chCall.Done
	if replyCall.Error != nil {
		fmt.Println(replyCall.Error)
	} else {
		fmt.Printf("rsp:%v\n", rsp.RealDamage)
	}
}

func SyncCall(cli *rpcx.Client) {
	req := &Req{100}
	rsp := Rsp{}
	err := cli.Call("Entity.RecvDamage", req, &rsp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("rsp:%v\n", rsp.RealDamage)
	}
}

func main() {
	fmt.Print()

	sel := &rpcx.DirectClientSelector{
		Network:     "tcp",
		Address:     "127.0.0.1:6666",
		DialTimeout: 10 * time.Second,
	}
	cli := rpcx.NewClient(sel)
	defer cli.Close()

	SyncCall(cli)
	AsyncCall(cli)
}
