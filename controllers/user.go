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
	info := &models.UserNoPwd{
		Id:     destUser.Id,
		Name:   destUser.Name,
		Avatar: destUser.Avatar,
		Status: destUser.Status,
		Ctime:  destUser.Ctime,
	}
	userInfo := models.UserInfo{
		UserInfo: info,
		Token:    token,
	}
	mes := &models.SuccessMsg{
		Code:    0,
		Data:    userInfo,
		Message: "登录成功!",
	}
	u.Data["json"] = mes
	u.ServeJSON()
}

func (u *UserController) Register() {
	name := u.GetString("name")
	password := u.GetString("password")
	confirm := u.GetString("confirm")
	status, err := u.GetInt("status")
	if err != nil {
		beego.Error("get status err:", err)
		errRes := models.SystemError{
			Code:    models.REGIIST_USERSTATUS_ERROR_CODE_,
			Message: models.SYSTEM_ERROR_MSG,
		}
		u.Data["json"] = errRes
		u.ServeJSON()
		return
	}
	// status验证
	if status < 0 && status > 7 {
		errRes := models.SystemError{
			Code:    models.PARAMS_ERROR_CODE,
			Message: models.INVALID_STATUS_ERROR_MSG,
		}
		u.Data["json"] = errRes
		u.ServeJSON()
		return
	}
	// 密码、确认密码验证
	if password != confirm {
		errRes := models.SystemError{
			Code:    models.PARAMS_ERROR_CODE,
			Message: models.PARAMS_ERROR_MSG,
		}
		u.Data["json"] = errRes
		u.ServeJSON()
		return
	}

	// 用户名长度验证
	if len([]rune(name)) > 8 {
		errRes := models.SystemError{
			Code:    models.PARAMS_ERROR_CODE,
			Message: models.USERNAME_MORELEN_ERROR_MSG,
		}
		u.Data["json"] = errRes
		u.ServeJSON()
		return
	}

	var userBox []models.Users
	result := models.Db.Where("name = ?", name).Find(&userBox)

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

	// 用户名重复验证
	if result.RowsAffected != 0 {
		errRes := models.SystemError{
			Code:    models.USER_ERROR_CODE,
			Message: models.USER_EXIST_ERROR_MSG,
		}
		beego.Info("user existed", name)
		u.Data["json"] = errRes
		u.ServeJSON()
		return
	}

	newUser := models.Users{
		Name:     name,
		Avatar:   "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Password: until.HashAndSalt(password),
		Status:   models.UserStatus(status),
		Ctime:    time.Now().Unix(),
	}

	result = models.Db.Create(&newUser)
	if result.Error != nil {
		errRes := models.SystemError{
			Code:    models.DB_ERROR_CODE,
			Message: models.SYSTEM_ERROR_MSG,
		}
		beego.Error("create user err:", result.Error)
		u.Data["json"] = errRes
		u.ServeJSON()
		return
	}

	logRes := map[string]interface{}{
		"code":    0,
		"message": "创建成功！",
	}
	u.Data["json"] = logRes
	u.ServeJSON()
	return
}

// 获取全部用户
func (u *UserController) All() {
	var users []models.UserNoPwd
	var usersAll []models.UserNoPwd
	resultCount := models.Db.Table("users").
		Scopes(models.FilterName(u.Ctx.Request)).
		Scopes(models.FilterId(u.Ctx.Request)).
		Scopes(models.FilterStatus(u.Ctx.Request)).
		Find(&usersAll)
	if resultCount.Error != nil {
		errRes := models.SystemError{
			Code:    models.DB_ERROR_CODE,
			Message: models.SYSTEM_ERROR_MSG,
		}
		errLogMsg := fmt.Sprintln("query all users err:", resultCount.Error)
		beego.Error(errLogMsg)
		u.Data["json"] = errRes
		u.ServeJSON()
	}
	result := models.Db.Table("users").
		Scopes(models.Paginate(u.Ctx.Request)).
		Scopes(models.FilterId(u.Ctx.Request)).
		Scopes(models.FilterName(u.Ctx.Request)).
		Scopes(models.FilterStatus(u.Ctx.Request)).
		Find(&users)
	if result.Error != nil {
		errRes := models.SystemError{
			Code:    models.DB_ERROR_CODE,
			Message: models.SYSTEM_ERROR_MSG,
		}
		errLogMsg := fmt.Sprintln("query all users err:", result.Error)
		beego.Error(errLogMsg)
		u.Data["json"] = errRes
		u.ServeJSON()
	}
	u.Data["json"] = models.SuccessMsg{
		Code: 0,
		Data: models.List{
			Total:    int(resultCount.RowsAffected),
			UserList: users,
		},
		Message: "获取用户列表成功！",
	}
	u.ServeJSON()
}

func (u *UserController) Update(){

}
