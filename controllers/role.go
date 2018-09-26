package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"goadmin/common"
	"goadmin/models"
	"strconv"
)
type RoleController struct {
	beego.Controller
}
// @router /admin/role/role_list/ [get]
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
func (c *RoleController)Edit()  {
	
}

func (c *RoleController)Access()  {
	group_id := c.Ctx.Input.Param(":group_id")
	ag  := models.AuthGroup{}
	bool,authgroup := ag.List(1)
	if !bool {
		fmt.Println("error access")
	}
	//menu := models.MenuList(3)
	fmt.Println("common.UserMenu",common.UserMenu)
	c.Data["node_list"] = common.UserMenu
	c.Data["auth_group"] = authgroup
	id,_ := strconv.Atoi(group_id)
	bool,auth_group := ag.ListOne(id,1)
	if bool {
		fmt.Println("auth_group:",auth_group)
		c.Data["this_group"] = auth_group
	}
	c.TplName = "admin/role/managergroup.html"
}

func (c *RoleController)WriteGroup()  {
	form := c.Ctx.Request.Form
	fmt.Println(form["rules[]"],len(form["rules[]"]))

	group_id := c.Input().Get("id")
	fmt.Println(group_id,len(group_id))
	AG := models.AuthGroup{}
	id,_ := strconv.Atoi(group_id)
	fmt.Println(id,AG)
	AG.Save(id,form["rules[]"])
	c.TplName = "admin/role/role_list.html"
}