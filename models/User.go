package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int `json:"id"`
	User_name string `json:"user_name"`
	Password string `json:"-"`
	Phone string `json:"phone"`
	Role_id *AuthGroup `orm:"column(role_id);rel(one)"`
	Company string `json:"company"`
	Email string `json:"email"`
	Parent_id int `json:"Parent_id"`
	Name string `json:"name"`
}




//获取用户角色id
func (u *User)GetRoleId(Uid int) (int,bool){
	o := orm.NewOrm()
	user := User{Id: Uid}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return 0,false
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return 0,false
	} else {
		fmt.Println(user.Role_id)
		return user.Role_id.Id,true
	}
}

func (u * User)List(offset int) (int64,[]*User){
	o := orm.NewOrm()
	user := []*User{}
	count,err := o.QueryTable("th_user").Filter("parent_id",1).Count()
	num,err := o.QueryTable("th_user").Filter("parent_id",1).Offset(offset).Limit(20).RelatedSel().All(&user)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("NUM",num)
	}
	return count,user
}


func (u * User)ListOne(id int) (User){
	o := orm.NewOrm()
	user := User{}
	err := o.QueryTable("th_user").Filter("id",id).One(&user)
	if err != nil {
		fmt.Println(err)
	}else {

	}
	return user
}
/**
	按id更新
 */
func UserUpdate(Id int,update map[string]interface{}) int64  {
	o := orm.NewOrm()
	num,err := o.QueryTable("th_user").Filter("id",Id).Update(update)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("修改",num)
	return num
}