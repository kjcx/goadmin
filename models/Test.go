package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Member struct {
	Id int
	Pid *Plot `orm:"column(Pid);rel(one)"`
} 
type Plot struct {
	Id int
	Name string
}

func (this *Member)List()  {
	m := []*Member{}
	o := orm.NewOrm()
	o.QueryTable("th_member").Filter("id", 1).RelatedSel().All(&m)
	for _, v := range m {
		fmt.Println(v.Pid.Name)
		fmt.Println(v.Pid.Name)
	}
	//fmt.Println(m)
}