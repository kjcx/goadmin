package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

/**
 * 权限认证类
 * 功能特性：
 * 1，是对规则进行认证，不是对节点进行认证。用户可以把节点当作规则名称实现对节点进行认证。
 *      $auth=new Auth();  $auth->check('规则名称','用户id')
 * 2，可以同时对多条规则进行认证，并设置多条规则的关系（or或者and）
 *      $auth=new Auth();  $auth->check('规则1,规则2','用户id','and')
 *      第三个参数为and时表示，用户需要同时具有规则1和规则2的权限。 当第三个参数为or时，表示用户值需要具备其中一个条件即可。默认为or
 * 3，一个用户可以属于多个用户组(think_auth_group_access表 定义了用户所属用户组)。我们需要设置每个用户组拥有哪些规则(think_auth_group 定义了用户组权限)
 *
 * 4，支持规则表达式。
 *      在think_auth_rule 表中定义一条规则时，如果type为1， condition字段就可以定义规则表达式。 如定义{score}>5  and {score}<100  表示用户的分数在5-100之间时这条规则才会通过。
 */
type AuthRule struct {
	Id int
	Module string
	Type int
	Name string
	Title string
	Status int
	Condition string
}
/**
   * 检查权限
   * @param name string|array  需要验证的规则列表,支持逗号分隔的权限规则或索引数组
   * @param uid  int           认证用户的id
   * @param string mode        执行check的模式
   * @param relation string    如果为 'or' 表示满足任一条规则即通过验证;如果为 'and'则表示需满足所有规则才能通过验证
   * @return boolean           通过验证返回true;失败返回false
  */
func (u *AuthRule)Check(Url string,Uid int,Mode int) bool{
	authlist := u.GetAuthList(Uid)
	fmt.Println("Url",Url)
	fmt.Println("authlist",authlist)
	for _,v := range authlist {
		if v.Name == Url {
			return true
		}
	}
	return false
}
//获取权限列表
func (ar *AuthRule)GetAuthList(Uid int) []AuthRule {
	o := orm.NewOrm()
	bool,Ids := GetGroups(Uid)
	AR := []AuthRule{}
	if bool {

		num,err := o.QueryTable("th_auth_rule").Filter("id__in", Ids).All(&AR)
		if err != nil {
			fmt.Println(err)
		}else {
			fmt.Println("NUM",num)
		}
	}
	return AR

}


