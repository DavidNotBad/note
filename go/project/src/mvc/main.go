package main

import (
	appHttp "mvc/app/http"
	"mvc/config"
	"mvc/routers"
	"net/http"
)

func main() {
	router := routers.Router
	mh := appHttp.NewMiddleWareHandler(router)
	err := http.ListenAndServe(config.Get("ListenAndServeAddr"), mh)
	if err != nil{
		panic(err)
	}
}


