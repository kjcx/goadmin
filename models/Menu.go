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
	fmt.Println("num:",num)
	fmt.Println("pid",menu)
	return menu,num
}
func List() []*Menu {
	//var mapmenu = make(map[int]*Menu,100)
	M := Menu{}
	menu,_ := M.Select(0)

	var i int
	for _,v:= range menu {
		i++
		fmt.Println(v)
		if v.Pid == 0 {
			childmenu,_ := M.Select(v.Id)
			v.Child = childmenu
		}
	}
	return menu
}
