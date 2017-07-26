package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	conf := &mysql.Config{
		User:   "test",
		Passwd: "N0FCzKew3M",
		Net:    "tcp",
		Addr:   "120.27.235.185:2101",
		DBName: "test",
	}
	fmt.Println(conf.FormatDSN())
	db, err := sql.Open("mysql", "test:N0FCzKew3M@tcp(120.27.235.185:2101)/test")
	if err != nil {
		fmt.Printf("open mysql failed, err:%v\n", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Printf("connect failed:%v\n", err)
		return
	}

	_, err = db.Exec("insert into t_test (id, name, description) values (?, ?, ?)", 11, "abc", "abc")
	if err != nil {
		fmt.Println(err)
		return
	}

	//db.Exec("insert into t_test (id, name, description) values ($1, $2, $3)", 2, "def", "def")
	//db.Exec("insert into t_test (id, name, description) values ($1, $2, $3)", 3, "ghi", "ghi")

	rs, err := db.Query("select id,name,description as 'desc' from t_test")
	if err != nil {
		fmt.Printf("query failed:%v\n", err)
		return
	}
	cols, _ := rs.Columns()
	fmt.Printf("result:%v\n", cols)
	//fmt.Println(r.Next())
	for rs.Next() {
		var v1, v2, v3 interface{}
		rs.Scan(&v1, &v2, &v3)
		fmt.Println(v1, v2, v3)
	}
	fmt.Printf("end db")
}
