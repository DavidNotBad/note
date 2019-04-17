package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	//router.

	return router
}

func main() {
	r := RegisterHandlers()

	serve := http.ListenAndServe(":8000", r)
	fmt.Println(serve)
}


