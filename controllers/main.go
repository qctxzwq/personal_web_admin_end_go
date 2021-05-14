package controllers

import (
	"admin/models"
	"github.com/astaxie/beego"
	"time"
)

type MainController struct {
	beego.Controller
}


func (m *MainController) AdminHome () {
	username := m.GetString("username")
	homeMes :=  models.HomeMessage{
		Username:username,
		Message:"首页请求成功",
	}
	response := &models.Home{
		Code:    0,
		Data:    homeMes,
		ReqTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	m.Data["json"] = response
	m.ServeJSON()
	return
}