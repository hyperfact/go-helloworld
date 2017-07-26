package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// M for bson
type M bson.M

const conURL = "mongodb://test:Pgk2HcVApHi307W7@aly-hz01.hjent.cn:2121/test"

func main() {
	session, err := mgo.Dial(conURL)
	if err != nil {
		return
	}
	defer session.Close()

	db := session.DB("test")

	fmt.Println(db.CollectionNames())

	c := db.C("test")

	//c.UpdateAll(M{}, M{"$set": M{"createTime": time.Now()}})

	q := c.Find(nil)
	rs := make([]interface{}, 10)
	q.All(&rs)
	fmt.Println(rs)

	var r interface{}
	q.One(&r)
	fmt.Println(r)
	fmt.Printf("%T\n", r)

	//fmt.Println(time.Now())

	tm := (r.(bson.M)["createTime"])
	fmt.Println("time:", tm)
}
