package controllers

import (
	"github.com/astaxie/beego"
	"goadmin/models"
)
type MenuController struct {
	beego.Controller
}

//菜单列表
func (c *MenuController)Index(){
	c.TplName = "admin/menu/index.html"
}

//菜单添加
func (c *MenuController)Add(){
	c.TplName = "admin/menu/index.html"
}
//菜单编辑
func (c *MenuController)Edit(id string){
	c.TplName = "admin/menu/edit.html"
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
	menu := new(models.Menu)
	ml,Total :=menu.Select(0)
	r := Record{}
	r.Rows = ml
	r.Total = Total
	c.Data["json"] = r
	c.ServeJSON()
}