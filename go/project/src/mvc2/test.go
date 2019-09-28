package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

type Users struct {
	Uid int64 `db:"uid"`
	Name string `db:"name"`
	Age int64 `db:"age"`
	Xxx interface{} `db:"-"` // 这个字段在orm中会忽略
}

func (u *Users) TableName() string {
	return "users"
}

var err error
var engin *gorose.Engin
var db gorose.IOrm

func init() {
	// 全局初始化数据库,并复用
	// 这里的engin需要全局保存,可以用全局变量,也可以用单例
	// 配置&gorose.Config{}是单一数据库配置
	// 如果配置读写分离集群,则使用&gorose.ConfigCluster{}
	engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: "root:root@tcp(127.0.0.1:3306)/test?charset=utf8"})
	engin.TagName("db")

	db = engin.NewOrm()
}

type Model struct {
	gorose.IOrm
}
type UsersModel struct {
	Model
}
func (um *UsersModel) Lists(u *Users) UsersModel {
	db.Table(u).Reset()
}



func main() {
	var u Users
	e := UserModel(&u).Where("uid", "1").Select()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(u)

	var u1 Users
	e1 := newUsers(&u1).Where("uid", 3).Select()
	if e1 != nil {
		fmt.Println(e1)
	}
	fmt.Println(u1)









	//// 这里定义一个变量db, 是为了复用db对象, 可以在最后使用 db.LastSql() 获取最后执行的sql
	//// 如果不复用 db, 而是直接使用 DB(), 则会新建一个orm对象, 每一次都是全新的对象
	//// 所以复用 db, 一定要在当前会话周期内
	//db := DB()
	//
	//// 查询一条
	//var u Users
	//// 查询数据并绑定到 user{} 上
	//_, err = db.Table(&u).Fields("uid,name,age").Where("age",">",0).OrderBy("uid desc").Get()
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(u)
	//fmt.Println(db.LastSql())
	//db.Reset()
	//
	//// 查询多条
	//// 查询数据并绑定到 []Users 上, 这里复用了 db 及上下文条件参数
	//// 如果不想复用,则可以使用DB()就会开启全新会话,或者使用db.Reset()
	//// db.Reset()只会清除上下文参数干扰,不会更换链接,DB()则会更换链接
	////var u2 []Users
	//err = db.Limit(9).Offset(1).Select()
	//fmt.Println(u)
	//fmt.Println(db.LastSql())
	//db.Reset()
	//
	//// 统计数据
	//var count int64
	//// 这里reset清除上边查询的参数干扰, 可以统计所有数据, 如果不清楚, 则条件为上边查询的条件
	//// 同时, 可以新调用 DB(), 也不会产生干扰
	//count,err = db.Reset().Count()
	//// 或
	//count, err = DB().Table(&u).Count()
	//fmt.Println(count, err)
	//db.Reset()
}