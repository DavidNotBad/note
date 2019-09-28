package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"mvc2/app/defs"
	"mvc2/app/models"
	"net/http"
)

var (
	model models.User
)

type UserController struct {
	Controller
}

func (u *UserController) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {

}

func (u *UserController) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	var user defs.User
	model.First(&user)
	fmt.Println(user)


	//var users []defs.User
	//model.Get(&users)
	//
	//fmt.Println(users)
	//fmt.Println(user.Data)
}



