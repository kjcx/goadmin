package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	User_name string
	Password string
	Phone string
	Role_id int
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
		return user.Role_id,true
	}
}