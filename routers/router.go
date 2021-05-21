package routers

import (
	"admin/controllers"
	"admin/filter"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, filter.GlobalFilter)
	beego.Router("/login",&controllers.UserController{},"post:Login")
	beego.Router("/user/all",&controllers.UserController{},"get:All")
}
