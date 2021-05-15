package models

var (
	UserList map[string]*User
)

type LoginMes struct {
	AutoLogin bool   `json:"type"`
	Username  string `json:"name"`
	Password  string `json:"password"`
	LoginType string `json:"type"`
}

type User struct {
	Id       string `json:"id"`
	Username string
	Password string
}

type UserInfo struct {
	UserInfo User
}

type LoginSuccess struct {
	Code    int
	Data    UserInfo
	Message string
}

type LoginFailed struct {
	Code    int
	Message string
}

func Login(username, password string) bool {
	if username == "zhang" && password == "123" {
		return true
	}
	return false
}
