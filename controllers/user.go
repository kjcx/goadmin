package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)
type UserController struct {
	beego.Controller
}


func (c *UserController)List(){
	c.TplName = "admin/user/user_list.html"
}

