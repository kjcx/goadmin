package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goadmin/models"
	_ "goadmin/routers"
)
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:shenyangtaihuakeji2017@tcp(sjs.jygeili.com:3306)/kada?charset=utf8", 30)
	//// register model
	orm.RegisterModelWithPrefix("th_", new(models.Menu))
	orm.RegisterModelWithPrefix("th_", new(models.User))
	orm.RegisterModelWithPrefix("th_", new(models.AuthGroup))
	orm.RegisterModelWithPrefix("th_", new(models.AuthRule))
	orm.RegisterModelWithPrefix("th_", new(models.Member))
	orm.RegisterModelWithPrefix("th_", new(models.Plot))
}
func FunStatus(status int)(out string){
	if status == 1 {
		return  "是"
	}else{
		return "否"
	}
}
func main() {
	orm.Debug = true
	beego.SetStaticPath("/__CSS__", "static/admin/css")
	beego.SetStaticPath("/__JS__", "static/admin/js")
	beego.SetStaticPath("/__IMG__", "static/admin/img")
	beego.SetStaticPath("/fonts", "static/admin/fonts")


	beego.AddFuncMap("FunStatus",FunStatus)
	beego.Run()
}

