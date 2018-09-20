package controllers

import (
	"github.com/astaxie/beego"
)
type LoginController struct {
	beego.Controller
}

// @router /login [get]
func (lg *LoginController)Login(){
	lg.Data["Website"] = "beego.me"
	lg.Data["Email"] = "astaxie@gmail.com"
	lg.TplName = "admin/login.html"
}
