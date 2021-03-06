package routers

import (
	"goadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//beego.Router("/admin/login", &controllers.MainController{},"get:Login")
	beego.Router("/admin/login/login", &controllers.LoginController{},"*:Login")
	beego.Router("/admin/index/index", &controllers.IndexController{},"*:Index")
	beego.Router("/admin/index/index_v1", &controllers.IndexController{},"*:Indexv1")
	beego.Router("/admin/menu/index", &controllers.MenuController{},"*:Index")
	beego.Router("/admin/menu/ajax", &controllers.MenuController{},"*:Ajax")
	beego.Router("/admin/menu/add/:id", &controllers.MenuController{},"*:Add")
	beego.Router("/admin/menu/edit/:id", &controllers.MenuController{},"*:Edit")
	beego.Router("/admin/menu/del/:id", &controllers.MenuController{},"*:Del")
	beego.Router("/admin/role/role_list", &controllers.RoleController{},"*:List")
	beego.Router("/admin/role/role_access/:group_id", &controllers.RoleController{},"*:Access")
	beego.Router("/admin/role/WriteGroup", &controllers.RoleController{},"*:WriteGroup")
	beego.Router("/admin/role/role_ajax", &controllers.RoleController{},"*:Ajax")
	beego.Router("/admin/role/role_edit/:id", &controllers.RoleController{},"*:Edit")
	beego.Router("/admin/role/role_add", &controllers.RoleController{},"*:Add")
	beego.Router("/admin/role/role_del", &controllers.RoleController{},"*:Del")
	beego.Router("/admin/admin/admin_list", &controllers.AdminController{},"*:List")
	beego.Router("/admin/admin/ajax", &controllers.AdminController{},"*:Ajax")
	beego.Router("/admin/admin/admin_add", &controllers.AdminController{},"*:Add")
	beego.Router("/admin/admin/admin_edit/:id", &controllers.AdminController{},"*:Edit")
	beego.Router("/admin/user/user_list", &controllers.UserController{},"*:List") //用户列表
	beego.Router("/admin/member/member_list", &controllers.MemberController{},"*:List") //用户列表
}
