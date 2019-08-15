package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gpmgo/gopm/modules/log"
	"mvc/config"
)

var (
	DB *sql.DB
	err error
)

func init() {
	driver := config.Get("db_driver")
	userName := config.Get("db_username")
	pwd := config.Get("db_password")
	url := config.Get("db_url")
	port := config.Get("db_port")
	dbName := config.Get("db_name")
	charset := config.Get("db_charset")

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

type Model struct {

}


func (m *Model) Parse(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err.Error())
			panic(err)
		}

		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		records = append(records, record)
	}
	return records
}