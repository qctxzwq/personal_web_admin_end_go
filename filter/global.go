package filter

import (
	"admin/models"
	"admin/until"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// 全局过滤器（判断token、用户路由权限）
func GlobalFilter(ctx *context.Context) {
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")

	if ctx.Input.URL() != "/login" {
		//判断是否携带AUTHORIZATION字段
		token := ctx.Input.Header("AUTHORIZATION")
		token = token[7:]
		beego.Debug(token)
		// 未携带token
		if len(token) == 0 || token == "null" {
			ctx.Output.SetStatus(401)
			data := models.SystemError{
				Code:    models.UNLOGIN_ERROR_CODE,
				Message: models.UNLOGIN_ERROR_MSG,
			}
			_ = ctx.Output.JSON(data, true, true)
			return
		}
		authMap, err := until.ParseToken(token)
		// token解析失败
		if err != nil {
			beego.Error("parse token failed,err:%v", err)
			data := models.SystemError{
				Code:    models.DB_ERROR_CODE,
				Message: models.SYSTEM_ERROR_MSG,
			}
			_ = ctx.Output.JSON(data, true, true)
			return
		}

		role := authMap["status"].(float64)

		if until.IsHaveAuth(int(models.RouteAuthMap[ctx.Input.URL()]), int(role)) {
			data := models.SystemError{
				Code:    models.NOT_HAVE_AUTH_ERROR_CODE,
				Message: models.NOT_HAVE_AUTH_ERROR_MSG,
			}
			_ = ctx.Output.JSON(data, true, true)
		}
	}

}
