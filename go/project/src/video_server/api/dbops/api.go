package dbops

import "log"

func AddUserCredential(loginName string, pwd string) (err error) {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return
	}

	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return
}

func GetUserCredential(loginName string) (pwd string, err error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return
	}

	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return
}

func DeleteUser(loginName string, pwd string) (err error) {
	stmtDel, err := dbConn.Prepare("DELETE FROM users where login_name= ? and pwd = ?")
	if err != nil {
		log.Printf("%s", err)
		return
	}

	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return
}