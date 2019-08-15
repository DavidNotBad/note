package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    //https://www.kancloud.cn/fizz/gorose-2/1135835
    "github.com/gohouse/gorose"
)


type Users struct {
	Uid  int64       `gorose:"uid"`
	Name string      `gorose:"name"`
	Age  int64       `gorose:"age"`
	Xxx  interface{} `gorose:"-"` // 这个字段在orm中会忽略
}

func (u *Users) TableName() string {
	return "users"
}

var engin *gorose.Engin

func init() {
	engin, _ = gorose.Open(&gorose.Config{
		Driver: "mysql",
		Dsn:    "root:root@tcp(127.0.0.1:3306)/test?charset=utf8",
	})
}

var (
    db = engin.NewOrm()
    u1 Users
    Users1 = db.Table(&u1)
)

func main() {
    var db = engin.NewOrm()

    var u1 Users
    _ = db.Table(&u1).Fields("uid,name,age").Where("age", ">", 0).OrderBy("uid desc").Select()
    db.Reset()
    fmt.Println(u1)

    var u2 []Users
    _ = db.Table(&u2).Fields("uid,name,age").Select()
    fmt.Println(u2)
}

