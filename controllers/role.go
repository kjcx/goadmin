package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

func (c *RoleController)Del(){
	id := c.Input().Get("id")
	Id,_ := strconv.Atoi(id)
	ag := models.AuthGroup{}
	ag.Del(Id)
	//c.TplName = "admin/role/role_list.html"
	resmap := map[string]bool{}
	resmap["status"] = true
	c.Data["json"] = resmap
	c.ServeJSON()
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
//添加角色
func (c *RoleController)Add(){
	if c.Ctx.Input.IsPost() {
		Title := c.Input().Get("title")
		Description := c.Input().Get("description")
		ag := models.AuthGroup{}
		ag.Add(Title,Description)
		//c.TplName = "admin/role/role_list.html"
		url := c.URLFor("RoleController.List")
		fmt.Println("url",url)
		c.Redirect(url, 302)
	}else {
		c.TplName = "admin/role/role_add.html"
	}
}
//编辑角色
func (c *RoleController)Edit()  {
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		Id,_ := strconv.Atoi(id)
		Title := c.Input().Get("title")
		Description := c.Input().Get("description")
		ag := models.AuthGroup{}
		ag.Edit(Id,Title,Description)
		c.TplName = "admin/role/role_list.html"

	}else{
		id := c.Ctx.Input.Param(":id")
		fmt.Println("id:",id)
		Id,_ := strconv.Atoi(id)
		ag := models.AuthGroup{Id:Id}
		o := orm.NewOrm()
		err := o.Read(&ag)
		if err != nil {
			fmt.Println(err)
		}
		c.Data["data"] = ag
		c.TplName = "admin/role/role_edit.html"

	}

}
//查询权限
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
//编辑权限
func (c *RoleController)WriteGroup()  {
	form := c.Ctx.Request.Form
	fmt.Println(form["rules[]"],len(form["rules[]"]))

	group_id := c.Input().Get("id")
	fmt.Println("g:",group_id,len(group_id))
	AG := models.AuthGroup{}
	id,_ := strconv.Atoi(group_id)
	fmt.Println("id:ag:",id,AG)
	AG.Save(id,form["rules[]"])
	c.TplName = "admin/role/role_list.html"
}