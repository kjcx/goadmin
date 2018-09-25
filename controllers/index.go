package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"goadmin/models"
)
type IndexController struct {
	beego.Controller
}
 type Menu struct {
 	Id int
 	Title string
 	Pid int
 	Sort int
 	Url string
 	Hide int
 	Tip int
 	Group string
 	is_dev int
 	Child [2]*Menu
 }
type MenuList struct {
	Main [1]*Menu
}

// @router /login [get]
func (c *IndexController)Index(){
	M := &Menu{}
	MC := &Menu{}
	M.Title = "权限管理"
	M.Url = "admin/role/role_list"

	ML := &MenuList{}
	MC.Title = "管理员列表"
	MC.Title = "管理员列表"
	MC.Url = "/admin/admin/admin_list"
	M.Child[0] = MC
	M.Child[1] = MC
	ML.Main[0] = M
	c.Data["Website"] = "beego.me"
	//c.Data["Menu"] = ML.Main
	//判断当前人拥有的菜单
	IsAdmin := false
	menu := models.List()
	fmt.Println("zzzz")
	if IsAdmin {
		c.Data["Menu"] = menu
	}else{
		menu := models.MenuList(2)
		c.Data["Menu"] = menu
	}

	c.TplName = "admin/index/index.html"
}

func (c *IndexController)Indexv1()  {
	ar := models.AuthRule{}
	ar.GetAuthList(3)

	c.TplName = "admin/index/index_v1.html"
}
