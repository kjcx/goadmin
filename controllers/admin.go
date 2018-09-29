package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"goadmin/common"
	"goadmin/models"
	"strconv"
)
type AdminController struct {
	beego.Controller
}

// @router /login [get]
func (c *AdminController)List(){
	c.TplName = "admin/admin/admin_list.html"
}

func (c *AdminController)Ajax(){
	offset := c.Input().Get("offset")
	fmt.Println("当前:",offset)
	offset_int,_ := strconv.Atoi(offset)

	user := models.User{}
	Total,List := user.List(offset_int)
	r := common.Record{}
	r.Rows = List
	r.Total = int(Total)
	c.Data["json"] = r
	c.ServeJSON()
}

func (c *AdminController)Add(){

}
//编辑管理员
func (c *AdminController)Edit(){
	if c.Ctx.Input.IsPost() {
		var usermap = make(map[string]interface{})

		id := c.Input().Get("id")
		Id,_ := strconv.Atoi(id)
		usermap["Id"] = Id
		password := c.Input().Get("password")
		if password != "" {
			usermap["Password"] = common.Md5(password)
		}
		user_name := c.Input().Get("user_name")
		role_id := c.Input().Get("role_id")
		usermap["role_id"] = role_id
		usermap["User_name"] = user_name
		company := c.Input().Get("company")
		usermap["Company"] = company
		name := c.Input().Get("name")
		usermap["Name"] = name
		email := c.Input().Get("email")
		usermap["Email"] = email
		phone := c.Input().Get("phone")
		usermap["Phone"] = phone
		usermap["name"] = name
		usermap["email"] = email

		models.UserUpdate(Id,usermap)
		url := c.URLFor("AdminController.List")
		fmt.Println("url",url)
		c.Redirect(url, 302)
	}else{
		id  := c.Ctx.Input.Param(":id")
		Id,_ := strconv.Atoi(id)
		fmt.Println(id)
		//查询id对应记录
		user := models.User{}
		Data := user.ListOne(Id)
		ag := models.AuthGroup{}
		bool,RoleList := ag.List(1)
		fmt.Println(RoleList)
		if bool {
			c.Data["Role"] = RoleList
		}
		c.Data["data"] = Data
		c.TplName = "admin/admin/admin_edit.html"
	}

}