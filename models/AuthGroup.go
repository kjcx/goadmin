package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
)

type AuthGroup struct {
	Id int
	Module string
	Type int
	Title string
	Description string
	Status int
	Rules string
	Cuid int

}
/**
	* 根据用户id获取角色
	* @param  uid int     用户id
	* @return array       用户所属的用户组 array(
	* array('uid'=>'用户id','group_id'=>'用户组id','title'=>'用户组名称','rules'=>'用户组拥有的规则id,多个,号隔开'),
	*                                         ...)
	*/
func GetGroups(Uid int) (bool,[]string)  {
	u :=  User{}
	var role_id int
	var bool bool
	role_id,bool = u.GetRoleId(Uid)
	if bool {
		fmt.Println(role_id)
	}else {
		fmt.Println("查询失败")
	}
	o := orm.NewOrm()
	AG := AuthGroup{Id: role_id}



	err := o.Read(&AG)
	var Ids []string
	if err == orm.ErrNoRows {
		bool = false
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		bool = false
		fmt.Println("找不到主键")

	} else {
		fmt.Println(AG.Rules)
		bool = true
		Ids = strings.Split(AG.Rules,",")
	}
	return bool,Ids

	//o := orm.NewOrm()
	//AG := &AuthGroup{Id:role_id}
	//err := o.Read(&AG)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(AG)
	//return &AG
}
//查询角色列表
func (ag *AuthGroup)List(cuid int) (bool,[]AuthGroup){
	o := orm.NewOrm()
	AG := []AuthGroup{}
	var bool bool
	num,err := o.QueryTable("th_auth_group").Filter("cuid", cuid).All(&AG)
	if err == orm.ErrNoRows {
		bool = false
		fmt.Println("查询不到",num)
	} else if err == orm.ErrMissPK {
		bool = false
		fmt.Println("找不到主键",num)

	} else {
		bool = true

	}
	return bool,AG
}


//查询角色列表
func (ag *AuthGroup)ListOne(id int,cuid int) (bool,AuthGroup){
	o := orm.NewOrm()
	AG := AuthGroup{}
	var bool bool
	num,err := o.QueryTable("th_auth_group").Filter("cuid", cuid).Filter("id",id).All(&AG)
	if err == orm.ErrNoRows {
		bool = false
		fmt.Println("查询不到",num)
	} else if err == orm.ErrMissPK {
		bool = false
		fmt.Println("找不到主键",num)

	} else {
		bool = true

	}
	return bool,AG
}

func (ag *AuthGroup)Save(id int,rules[]string){
	o := orm.NewOrm()
	AG := AuthGroup{Id:id}
	AG.Rules = strings.Join(rules,",")
	fmt.Println(AG)
	if num, err := o.Update(&AG,"rules"); err == nil {
		fmt.Println(num)
	}
}