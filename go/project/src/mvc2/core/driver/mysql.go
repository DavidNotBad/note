package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mvc2/configs"
)


var (
	DB *sql.DB
	err error
)

func init()  {
	driver := configs.DbDriver
	userName := configs.DbUsername
	pwd := configs.DbPassword
	url := configs.DbUrl
	port := configs.DbPort
	dbName := configs.DbName
	charset := configs.DbCharset

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", userName, pwd, url, port, dbName, charset)
	DB, err = sql.Open(driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	DB.SetMaxOpenConns(1000)

	if err := DB.Ping(); err != nil {
		panic("Failed to connect to mysql, err:" + err.Error())
	}
}


//func First(row *sql.Row, m interface{}) error {
//	mVal := reflect.ValueOf(m)
//	mInd := reflect.Indirect(mVal)
//
//	rsNumField := reflect.TypeOf(m).Elem().NumField()
//	tmp := make([]interface{}, rsNumField)
//	r := reflect.New(mVal.Elem().Type())
//
//	for i := 0; i < rsNumField; i++ {
//		tmp[i] = r.Elem().Field(i).Addr().Interface()
//	}
//
//	err = row.Scan(tmp...)
//	if err != nil{
//		return err
//	}
//
//	mInd.Set(reflect.Indirect(r))
//
//	return nil
//}
//
//
//
//
//
//func ParseRows1(rows *sql.Rows) []map[string]interface{} {
//	columns, _ := rows.Columns()
//	scanArgs := make([]interface{}, len(columns))
//	values := make([]interface{}, len(columns))
//	//for j := range values {
//	//	scanArgs[j] = &values[j]
//	//}
//	for k,_ := range values {
//		scanArgs[k] = &values[k]
//	}
//
//	record := make(map[string]interface{})
//	records := make([]map[string]interface{}, 0)
//	for rows.Next() {
//		//将行数据保存到record字典
//		_ = rows.Scan(scanArgs...)
//fmt.Println(values)
//		os.Exit(0)
//		for i, col := range values {
//			if col != nil {
//				record[columns[i]] = col
//			}
//		}
//		records = append(records, record)
//	}
//	return records
//}