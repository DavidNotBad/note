package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mvc/app/def"
	"os"
	"reflect"
)

type User struct {
	Model
}

type Mod interface {}


func parse(rows *sql.Rows, mod interface{})(records []interface{}, err error)  {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
fmt.Println(reflect.TypeOf(mod).String())
	t := reflect.ValueOf(mod).Type()
fmt.Println(t)
	v := reflect.New(t).Elem()
fmt.Println(v)
//os.Exit(0)
	record := make(map[string]interface{})
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		if err != nil{
			return
		}

		u := mod
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}

		bytes, err := json.Marshal(record)
		if err != nil{
			return nil, err
		}
		err = json.Unmarshal(bytes, &u)

		if err != nil{
			return nil, err
		}
		records = append(records, u)
	}
	return
}


func (u *User) Insert(){
	rows, _ := DB.Query("select * from users")

	t := def.User{}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	reflectStruct := reflect.ValueOf(&t).Elem()

	for i, v := range columns {
		fmt.Println(reflectStruct.FieldByName(v).Addr())
		os.Exit(0)
		values[i] = reflectStruct.FieldByName(v).Addr().Interface()
		fmt.Println(values[i])
		os.Exit(0)
	}
	//for rows.Next() {
	//	rows.Scan(values...)
	//}
	//fmt.Println(values)

	//records, _ := parse(rows, &user)
	//
	//fmt.Println(records)
}



