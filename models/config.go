package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var ProCfg config.Configer

func InitConfig() (err error) {
	ProCfg, err = config.NewConfig("ini", "conf/app.conf")
	InitStatic()
	if err != nil {
		return
	}
	return
}

func InitStatic() () {
	staticpath := ProCfg.String("static")
	beego.SetStaticPath("/static",staticpath)
}
