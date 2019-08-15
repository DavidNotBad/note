package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

var db *sql.DB

type User struct {
	UID   int `json:"uid"`
	Name string `json:"name"`
	Age int `json:"age"`
}

//https://www.cnblogs.com/mmdsnb/p/6439267.html
//https://www.jianshu.com/p/c4ec92afeca8
func main() {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")

	sqlStr := "select * from users"

	var datas2 []User
	e := Parse(sqlStr, &datas2)
	fmt.Println(datas2, e)
}

func Parse(sqlStr string, m interface{}) error {
	mVal := reflect.ValueOf(m)
	//mTyp := reflect.TypeOf(m)
	mInd := reflect.Indirect(mVal)

	rows, err := db.Query(sqlStr)
	if err != nil{
		return err
	}
	defer func() {
		err = rows.Close()
		return
	}()

	columns, _ := rows.Columns()
	rs := mVal.Elem().Type().Elem()

	for rows.Next() {
		r := reflect.New(rs)

		tmp := make([]interface{}, len(columns))
		rsNumField := rs.NumField()
		for i := 0; i < rsNumField; i++ {
			tmp[i] = r.Elem().Field(i).Addr().Interface()
		}

		err = rows.Scan(tmp...)
		if err != nil{
			return err
		}

		mInd.Set(reflect.Append(mInd, reflect.Indirect(r)))
	}
	return err
}








