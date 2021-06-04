package models

type ErrorCode int
type ErrorMessage string

const (
	UNLOGIN_ERROR_CODE                 ErrorCode = 401
	NOT_HAVE_AUTH_ERROR_CODE           ErrorCode = 402
	DB_ERROR_CODE                      ErrorCode = 500
	USER_ERROR_CODE                    ErrorCode = 501
	PASSWORD_ERROR_CODE                ErrorCode = 502
	TOKEN_CREATE_ERROR_CODE            ErrorCode = 503
	NICK_NAME_ERROR_CODE               ErrorCode = 504
	REGIIST_USERSTATUS_ERROR_CODE_     ErrorCode = 505
	GET_REGIIST_AVATARFILE_ERROR_CODE_ ErrorCode = 506
	PARAMS_ERROR_CODE                  ErrorCode = 507
	UPLOAD_ERROR_CODE                  ErrorCode = 508
	CREATE_FOLDER_ERROR_CODE           ErrorCode = 509
)

const (
	UNLOGIN_ERROR_MSG          ErrorMessage = "未登录！"
	NOT_HAVE_AUTH_ERROR_MSG    ErrorMessage = "用户权限不足！"
	SYSTEM_ERROR_MSG           ErrorMessage = "系统错误，请稍后重试！"
	USER_ERROR_MSG             ErrorMessage = "用户不存在！"
	PASSWORD_ERROR_MSG         ErrorMessage = "密码错误！"
	NICK_NAME_ERROR_MSG        ErrorMessage = "昵称或者密码为空！"
	INVALID_STATUS_ERROR_MSG   ErrorMessage = "无效的用户状态！"
	PARAMS_ERROR_MSG           ErrorMessage = "参数错误！"
	USER_EXIST_ERROR_MSG       ErrorMessage = "用户已存在！"
	USERNAME_MORELEN_ERROR_MSG ErrorMessage = "用户名过长！"
	UPLOAD_FILE_ERROR_MSG      ErrorMessage = "上传失败！"
)

type SystemError struct {
	Code    ErrorCode    `json:"code"`
	Message ErrorMessage `json:"message"`
}
