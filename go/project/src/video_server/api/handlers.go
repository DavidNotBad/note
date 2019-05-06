package main 

import (
	"io"
	"encoding/json"
	"net/http"
<<<<<<< HEAD
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"video_server/api/defs"
	"video_server/api/dbops"
	"video_server/api/session"
=======
    "fmt"
    "encoding/json"
    "io/ioutil"
    "github.com/julienschmidt/httprouter"
    "github.com/avenssi/video_server/api/defs"
>>>>>>> 3e344f5bc7b4376984a06829d2c44d4b5c7b022f
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
<<<<<<< HEAD
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return 
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
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
=======
    ubody := &defs.UserCredential{}

    if err := json.Unmarshal(res, ubody); err != nil {
        sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
        return 
    }

    fmt.Println(ubody)
>>>>>>> 3e344f5bc7b4376984a06829d2c44d4b5c7b022f
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}