package models

type userStatus int

const (
	SUPER_ADMIN       userStatus = 0
	NORMAL_ADMIN      userStatus = 1
	SUPER_USER        userStatus = 2
	NORMAL_USER       userStatus = 3
	FORBIDEN_USER_ONE userStatus = 4
	FORBIDEN_USER_TWO userStatus = 5
	VISITOR           userStatus = 6
)

var (
	UserList map[string]*Users
)

type LoginMes struct {
	AutoLogin bool   `json:"type"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	LoginType string `json:"type"`
}

type Users struct {
	Id       int        `json:"id";gorm:"primaryKey"`
	Name     string     `json:"name"`
	Avatar   string     `json:"avatar"`
	Password string     `json:"password"`
	Status   userStatus `json:"status"`
	Ctime    int64      `json:"ctime"`
}

type UserInfo struct {
	UserInfo Users `json:"user_info"`
}

type LoginSuccess struct {
	Code    int      `json:"code"`
	Data    UserInfo `json:"data"`
	Message string   `json:"message"`
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
