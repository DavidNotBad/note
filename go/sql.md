## a

```go
package models

import (
	"database/sql"
	"errors"
	"fmt"
	"mvc2/core/driver"
	"os"
	"reflect"
)

var (
	DB = driver.DB
	DBTagName = "db"
)


type Model struct {
	Data interface{}
}


func (model *Model) Get(m interface{}) *Model  {
	model.Data = m

	//prepare???
	rows, err := DB.Query("select age,uid,name from users")
	if err != nil{
		panic(err)
	}
	err = model.ParseRows(rows, m)
	if err != nil{
		panic(err)
	}

	return model
}

func (model *Model) First(m interface{}) *Model {
	model.Data = m

	//prepare???
	rows := DB.QueryRow("select age,uid,name from users")
	err := model.ParseRow(rows, m)
	if err != nil{
		panic(err)
	}

	return model
}



func (model *Model) ParseMapRows(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil{
			return nil, err
		}

		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]uint8))
			}
		}
		records = append(records, record)
	}
	return records, nil
}


func (model *Model) ParseRow(row *sql.Row, m interface{}) error {

	field := reflect.ValueOf(row).Elem()

	fmt.Println(reflect.TypeOf(field.FieldByName("rows").Elem()))
	r := reflect.New(field.FieldByName("rows").Elem().Type())
	fmt.Println(reflect.TypeOf(r))
	fmt.Println(r)
	os.Exit(0)



	//mVal := reflect.ValueOf(m)
	//mInd := reflect.Indirect(mVal)
	//
	//rsNumField := 3
	//tmp := make([]interface{}, rsNumField)
	//r := reflect.New(mVal.Elem().Type())
	//
	//for i := 0; i < rsNumField; i++ {
	//	tmp[i] = r.Elem().Field(i).Addr().Interface()
	//}
	//
	//err := row.Scan(tmp...)
	//if err != nil{
	//	return err
	//}
	//
	//mInd.Set(reflect.Indirect(r))
	return nil
}




func (model *Model) ParseRows(rows *sql.Rows, m interface{}) error {
	mVal := reflect.ValueOf(m)
	mTyp := reflect.TypeOf(m)
	mInd := reflect.Indirect(mVal)
	columns, _ := rows.Columns()
	lenColumns := len(columns)
	rs := mVal.Elem().Type().Elem()
	t := mTyp.Elem().Elem()

	for rows.Next() {
		r := reflect.New(rs)
		tmp := make([]interface{}, lenColumns)

		for i := 0; i < lenColumns; i++ {
			field, b := t.FieldByNameFunc(func(s string) bool {
				field, bf := t.FieldByName(s)
				if ! bf {
					return false
				}

				value, lok := field.Tag.Lookup(DBTagName)
				if ! lok {
					lok2 := false
					value, lok2 = field.Tag.Lookup("json")
					if ! lok2 {
						return false
					}
				}

				return value == columns[i]
			})
			if !b {
				return errors.New("FieldByNameFunc: field \"" + columns[i] + "\" not found" )
			}else{
				tmp[i] = r.Elem().FieldByIndex(field.Index).Addr().Interface()
			}
		}

		err := rows.Scan(tmp...)
		if err != nil{
			return err
		}

		mInd.Set(reflect.Append(mInd, reflect.Indirect(r)))
	}
	return nil
}
```

## gorose的model

```go
// model.go
package models

import (
	"github.com/gohouse/gorose"
	"test/db"
)

var conn gorose.IOrm
func init()  {
	conn = db.DB()
}


type Model struct {
}

func (model *Model) Table(m interface{}) gorose.IOrm {
	return conn.Table(m).Reset()
}

func (model *Model) panic(err interface{}) {
	if err != nil{
		panic(err)
	}
}

//user.go
package models

type Users struct {
	Model
}
func (model *Users) Lists(u interface{}) {
	e := model.Table(u).Where("uid", "1").Select()
	model.panic(e)
}

//controller
usersModel := models.Users{}

var u defs.Users
usersModel.Lists(&u) //models.Users.Lists(&u)
fmt.Println(u)

//补充 conn.go
package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

var err error
var engin *gorose.Engin

func init() {
	engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: "root:root@tcp(127.0.0.1:3306)/test?charset=utf8"})
	engin.TagName("db")
}

func DB() gorose.IOrm {
	return engin.NewOrm()
}
```

