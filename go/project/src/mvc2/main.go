package main

import (
	"mvc2/app"
	"mvc2/configs"
	"mvc2/routers"
	"net/http"
)

func main() {
	router := routers.Router
	mh := app.NewMiddleWareHandler(router)
	err := http.ListenAndServe(configs.ListenAndServeAddr, mh)
	if err != nil{
		panic(err)
	}
}


