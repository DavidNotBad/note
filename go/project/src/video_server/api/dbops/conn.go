package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/my_videos?timeout=30s&charset=utf8&parseTime=true")
	if err != nil {
		panic(err.Error())
	}
}
