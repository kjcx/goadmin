package controllers

import (
	"github.com/astaxie/beego"
)
type AdminController struct {
	beego.Controller
}

// @router /login [get]
func (c *AdminController)List(){
	c.TplName = "admin/admin_list.html"
}
