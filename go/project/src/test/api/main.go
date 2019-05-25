package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main()  {
	r := registerHandlers()

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}

func registerHandlers() (router *httprouter.Router) {
	router = httprouter.New()

	router.GET("/", Index)

	return
}

func Index(writer http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("aaaaa")
}

