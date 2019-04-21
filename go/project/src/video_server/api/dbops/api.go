package dbops

import (
	"log"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

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


func AddNewVideo(aid int, name string)(*defs.VideoInfo, error) {
	//create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("insert into video_info(id, author_id, name, display_ctime) values (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtime:ctime}
	defer stmtIns.Close()
	return res, nil
}



