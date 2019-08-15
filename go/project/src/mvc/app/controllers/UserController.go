package controllers

import (
	"github.com/julienschmidt/httprouter"
	"mvc/app/models"
	"net/http"
)

type UserController struct {
	Controller
}

func (u *UserController) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	user := models.User{}
	user.Insert()
	//user1 := models.User{}
	//user1.Insert()
}

func (u *UserController) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {

}



