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
	Up_title string `orm:"-"` //上级菜单
	Child []*Menu `orm:"-"`
}
//查询菜单
func (m *Menu)Select(pid int) ([]*Menu,int64) {
	o := orm.NewOrm()
	var menu []*Menu
	num,_ := o.QueryTable("th_menu").Filter("pid",pid).OrderBy("Sort").All(&menu)
	menuall := m.MenuAll()

	for _,v := range  menu {
		v.Up_title = menuall[v.Pid]
	}
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
//查询所有菜单
func (m * Menu)MenuAll() map[int]string {
	o := orm.NewOrm()
	menu := [] Menu{}
	num,err := o.QueryTable("th_menu").All(&menu)
	if err != nil {
		fmt.Println(err)
	}
	var mmap = make(map[int]string)
	if num > 0 {
		 for _,v := range menu {
			 mmap[v.Id] = v.Title
		 }
	}

	return mmap
}
//查询一条记录
func (m *Menu) ListOne(Id int) Menu{
	o := orm.NewOrm()
	menu := Menu{Id:Id}
	err := o.Read(&menu)
	if err != nil {
		fmt.Println(err)
	}
	return menu
}

//更新
func (m *Menu)Update(mm Menu) bool{
	o := orm.NewOrm()
	num,err := o.Update(&mm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if num >0 {
		return true
	}
	return  false
}
