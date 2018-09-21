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
		ar := models.AuthRule{}
		for k,v := range menu {
			bool := ar.Check(v.Url,3,1)
			fmt.Println(v.Url)
			if bool {
				fmt.Println(k,"一级有权限",v.Url)
				for _,vv := range v.Child {
					booll := ar.Check(vv.Url,3,1)
					if booll {
						fmt.Println(k,"二级有权限",vv.Url)
					}else{
						fmt.Println(k,"二级无权限",vv.Url)
					}

				}
			}else {
				//a = append(a[:i], a[i+1:]...)
				menu = append(menu[:k], menu[k+1:]...)
				fmt.Println(k,"无权限",v.Url)
			}
		}
		for _,v := range menu{
			fmt.Println("Url:",v.Url)
		}
		c.Data["Menu"] = menu
	}

	c.TplName = "admin/index/index.html"
}

func (c *IndexController)Indexv1()  {
	ar := models.AuthRule{}
	ar.GetAuthList(3)
	c.TplName = "admin/index/index_v1.html"
}
