package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video_server_1_5/api/dbops"
	"video_server_1_5/api/defs"
	"video_server_1_5/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	fmt.Println(string(res))

	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(err)
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return 
	}

	fmt.Println(ubody)
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		fmt.Println(err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}