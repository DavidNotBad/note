package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	var user User
	user.Name = "slene"
	user.IsActive = true

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}


}
