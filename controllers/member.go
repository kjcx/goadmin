package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)
type MemberController struct {
	beego.Controller
}


func (c *MemberController)List(){
	c.TplName = "admin/member/member_list.html"
}

