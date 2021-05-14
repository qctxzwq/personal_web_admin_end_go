package routers

import (
	"admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.MainController{},"get:AdminHome")
	beego.Router("/login",&controllers.MainController{},"post:AdminHome")
}
