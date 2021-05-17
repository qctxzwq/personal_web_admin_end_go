package controllers

import (
	"admin/models"
	"admin/until"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
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

	var userBox []models.Users
	var destUser models.Users
	result := models.Db.Where("name = ?", loginMes.Name).Find(&userBox)

	// 检索错误
	if result.Error != nil {
		errRes := models.SystemError{
			Code:    models.DB_ERROR_CODE,
			Message: models.SYSTEM_ERROR_MSG,
		}
		beego.Info("query user error:", result.Error)
		u.Data["json"] = errRes
		u.ServeJSON()
		return
	}

	// 用户不存在
	if result.RowsAffected == 0 {
		errRes := models.SystemError{
			Code:    models.USER_ERROR_CODE,
			Message: models.USER_ERROR_MSG,
		}
		beego.Info("user not exist", loginMes)
		u.Data["json"] = errRes
		u.ServeJSON()
		return
	}

	destUser = userBox[0]

	// 密码错误
	if !until.ComparePasswords(destUser.Password, loginMes.Password) {
		errRes := models.SystemError{
			Code:    models.PASSWORD_ERROR_CODE,
			Message: models.PASSWORD_ERROR_MSG,
		}
		beego.Debug("user password error:", loginMes)
		u.Data["json"] = errRes
		u.ServeJSON()
		return
	}

	// 创建token
	token, err := until.CreateToken(destUser)
	if err != nil {
		errRes := models.SystemError{
			Code:    models.TOKEN_CREATE_ERROR_CODE,
			Message: models.SYSTEM_ERROR_MSG,
		}
		errLogMsg := fmt.Sprintf("userId:%v,create token failed,err:%v", destUser.Id, err)
		beego.Error(errLogMsg)
		u.Data["json"] = errRes
		u.ServeJSON()
		return
	}
	// 组合数据
	info :=  make(map[string]interface{})
	info["id"] = destUser.Id
	info["name"] = destUser.Name
	info["status"] = destUser.Status
	info["avatar"] = destUser.Avatar
	info["ctime"] = destUser.Ctime
	userInfo := models.UserInfo{
		UserInfo: info,
		Token:    token,
	}
	mes := &models.LoginSuccess{
		Code:    0,
		Data:    userInfo,
		Message: "登录成功!",
	}
	u.Data["json"] = mes

	u.ServeJSON()
}

func (u *UserController) register() {
	models.Db.AutoMigrate(&models.Users{})
	pass1 := until.HashAndSalt("123456")
	user1 := &models.Users{
		Name:     "admin2",
		Avatar:   "",
		Password: pass1,
		Status:   models.SUPER_ADMIN,
		Ctime:    time.Now().Unix(),
	}
	fmt.Println(user1)
	result := models.Db.Create(user1)
	if result.Error != nil {
		beego.Error("insert user%v to database failed,err:%v", result.Error)
		return
	}
}
