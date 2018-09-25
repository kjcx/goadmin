package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"goadmin/common"
	"goadmin/models"
)
type RoleController struct {
	beego.Controller
}

func (c *RoleController)List(){
	c.TplName = "admin/role/role_list.html"
}


func (c *RoleController)Ajax(){
	AG := models.AuthGroup{}
	bool,List := AG.List(1)
	if bool {
		r := common.Record{}
		r.Rows = List
		r.Total = len(List)
		c.Data["json"] = r
		c.ServeJSON()
	}

}

