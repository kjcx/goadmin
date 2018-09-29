package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Menu struct {
	Id int
	Title string
	Pid int
	Sort int
	Url string
	Hide int
	Tip string
	Group string
	Is_dev int
	Child []*Menu `orm:"-"`
}
//查询菜单
func (m *Menu)Select(pid int) ([]*Menu,int64) {
	o := orm.NewOrm()
	var menu []*Menu
	num,_ := o.QueryTable("th_menu").Filter("pid",pid).OrderBy("Sort").All(&menu)
	return menu,num
}
func List() []*Menu {
	//var mapmenu = make(map[int]*Menu,100)
	M := Menu{}
	menu,_ := M.Select(0)

	for _,v:= range menu {
		//fmt.Println(v)
		if v.Pid == 0 {
			childmenu,_ := M.Select(v.Id)
			v.Child = childmenu
		}
	}
	return menu
}
//查询用户对应的菜单列表
func MenuList(Uid int) []*Menu{
	M := Menu{}
	menu,_ := M.Select(0)
	var UserMenu []*Menu
	//过滤掉当前人拥有的
	for _,v := range menu {
		bool := Check(v.Id,3)
		if bool {
			//保留url
			fmt.Println(v.Id,"=>",v.Url)
			UserMenu = append(UserMenu,v)
			//fmt.Println("M",len(UserMenu))
		}
	}
	for _,v:= range UserMenu {
		if v.Pid == 0 {
			childmenu,_ := M.Select(v.Id)
			v.Child = childmenu
		}
	}
	return UserMenu
}
func Add(menu Menu) bool{
	o := orm.NewOrm()
	//mm := Menu{Title:strmenu["title"]}
	num,_ := o.Insert(menu)
	if num > 0 {
		return true
	}else {
		return false
	}
}