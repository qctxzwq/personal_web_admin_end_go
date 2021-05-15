package models

type errorCode int
type errorMessage string

const (
	DB_ERROR_CODE       errorCode = 500
	USER_ERROR_CODE     errorCode = 501
	PASSWORD_ERROR_CODE errorCode = 502
)

const (
	SYSTEM_ERROR_MSG   errorMessage = "系统错误，请稍后重试！"
	USER_ERROR_MSG     errorMessage = "用户不存在！"
	PASSWORD_ERROR_MSG errorMessage = "密码错误！"
)

type SystemError struct {
	Code    errorCode    `json:"code"`
	Message errorMessage `json:"message"`
}