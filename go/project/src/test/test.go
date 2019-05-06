package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}


func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.

}



func main() {

	router := httprouter.New()

	router.POST("/user", CreateUser)

	// router.POST("/user/:user_name", Login)

	http.ListenAndServe(":8000", router)
}

