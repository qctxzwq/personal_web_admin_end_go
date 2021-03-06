package routers

import (
	"admin/controllers"
	"admin/filter"

	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, filter.GlobalFilter)
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/upload/img", &controllers.StaticController{}, "post:UpdateImg")
	beego.Router("/user/all", &controllers.UserController{}, "get:All")
	beego.Router("/user/update", &controllers.UserController{}, "post:Update")
}
