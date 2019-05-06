package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "github.com/julienschmidt/httprouter"
    "github.com/avenssi/video_server/api/defs"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
    ubody := &defs.UserCredential{}

    if err := json.Unmarshal(res, ubody); err != nil {
        sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
        return 
    }

    fmt.Println(ubody)
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

