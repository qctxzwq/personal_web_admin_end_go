package main

import (
	"admin/models"
	_ "admin/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.InitLogger()
	err := models.InitDb()
	if err != nil {
		beego.Error("database connect failed:", err)
		return
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	}
	beego.Run()
	beego.Info("application running success!")
}
