package routers

import (
	"github.com/julienschmidt/httprouter"
	"mvc/app/controllers"
)

var Router = httprouter.New()

var (
	userController = &controllers.UserController{}
)

func init() {
	Router.GET("/user/:id", userController.CreateUser)
	Router.GET("/user", userController.Index)
}









