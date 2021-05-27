package models

import "github.com/astaxie/beego/config"

var ProCfg config.Configer

func InitConfig() (err error) {
	ProCfg, err = config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		return
	}
	return
}
