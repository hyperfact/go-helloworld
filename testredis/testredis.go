package main

import (
	"context"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
)

const (
	addr   = "aly-hz01.hjent.cn:2111"
	passwd = "lm6s4Z63K5n5CrPowKQtlAAijJGeoncF"
)

type redisRes struct {
	redis.Conn
}

func (c redisRes) Close() {
	c.Conn.Close()
}

func testRedisPool() {
	p := redis.Pool{
		Dial: func() (redis.Conn, error) {
			r, err := redis.Dial("tcp", addr, redis.DialPassword(passwd))
			return r, err
		},

		MaxIdle:   2,
		MaxActive: 8,
	}

	defer p.Close()

	c := p.Get()
	c.Send("GET", "abc")
	c.Flush()
	val, err := redis.String(c.Receive())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("result:", val)
}

func testVitnessPool() {
	p := pools.NewResourcePool(func() (pools.Resource, error) {
		r, err := redis.Dial("tcp", addr, redis.DialPassword(passwd))
		return redisRes{r}, err
	}, 1, 2, time.Minute)
	defer p.Close()

	ctx := context.TODO()
	r, err := p.Get(ctx)
	if err != nil {
		return
	}
	defer p.Put(r)

	c := r.(redisRes)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Send("SET", "abc", "123456")
	c.Send("GET", "abc")
	c.Flush()
	c.Receive()
	val, err := redis.String(c.Receive())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("result:", val)
}

func main() {
	testRedisPool()
}
