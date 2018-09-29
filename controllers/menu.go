package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"goadmin/common"
	"goadmin/models"
	"strconv"
)
type MenuController struct {
	beego.Controller
}

func (c *MenuController) Prepare() {
	//Url := c.Ctx.Request.URL.Path
	//bool := common.Check(Url,3,1)
	//if !bool && Url != "/admin/menu/ajax" {
	//	c.Data["json"] = "无权访问"
	//	c.ServeJSON()
	//}
	//return
}
//菜单列表
func (c *MenuController)Index(){
	c.TplName = "admin/menu/index.html"
}

//菜单添加
func (c *MenuController)Add(){
	if c.Ctx.Input.IsPost() {
		title := c.Input().Get("title")
		url := c.Input().Get("url")
		sort := c.Input().Get("sort")
		group := c.Input().Get("group")
		pid := c.Input().Get("pid")
		M := models.Menu{}

		M.Title = title
		M.Url = url
		M.Sort = common.StringToInt(sort)
		M.Group = group
		M.Pid = common.StringToInt(pid)
		models.Add(M)
	} else {
		id := c.Input().Get("id")
		fmt.Println(id)
		c.TplName = "admin/menu/add.html"
	}

}
//菜单编辑
func (c *MenuController)Edit(){
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		Id := common.StringToInt(id)
		m := models.Menu{}
		m.Id = Id
		title := c.Input().Get("title")
		m.Title = title
		sort := c.Input().Get("sort")
		m.Sort = common.StringToInt(sort)
		url :=c.Input().Get("url")
		m.Url = url
		pid := c.Input().Get("pid")
		m.Pid = common.StringToInt(pid)
		group := c.Input().Get("group")
		m.Group = group
		m.Update(m)
		c.TplName = "admin/menu/index.html"
	}else{
		id := c.Ctx.Input.Param(":id")
		Id,_ := strconv.Atoi(id)
		fmt.Println("Id",id)
		m := models.Menu{}
		Data := m.ListOne(Id)
		menus := models.List()
		fmt.Println(menus)
		for k,v := range  menus {
			fmt.Println(v.Id)
			fmt.Println(k)
		}
		c.Data["menus"] = menus
		c.Data["data"] = Data
		c.TplName = "admin/menu/edit.html"
	}

}

//菜单删除
func (c *MenuController)Del(id string){
	c.TplName = "admin/menu/del.html"
}

type Record struct {
	Rows []*models.Menu `json:"rows"`
	Total int64 `json:"total"`
}
//菜单列表数据
func (c *MenuController)Ajax(){
	pid := c.Input().Get("pid")
	fmt.Println("pid",pid)
	Pid,_ := strconv.Atoi(pid)

	menu := new(models.Menu)
	ml,Total :=menu.Select(Pid)
	r := Record{}
	r.Rows = ml
	r.Total = Total
	c.Data["json"] = r
	c.ServeJSON()
}