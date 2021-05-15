package controllers

import (
	"admin/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var loginMes models.LoginMes
	json.Unmarshal(u.Ctx.Input.RequestBody, &loginMes)
	fmt.Println(loginMes)


	if models.Login(loginMes.Username,loginMes.Password) {
		user := models.User{
			Id:       "10",
			Username: "zhangwangqian",
			Password: "123456",
		}
		mes := &models.LoginSuccess{
			Code:    0,
			Data:    models.UserInfo{user},
			Message: "登录成功!",
		}
		u.Data["json"] = mes
	} else {
		mes := models.LoginFailed{
			Code:    0,
			Message: "用户名或密码错误！",
		}
		u.Data["json"] = mes
	}
	u.ServeJSON()
}
