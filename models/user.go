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

// 登录必传字段
type LoginMes struct {
	AutoLogin bool   `json:"type"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	LoginType string `json:"type"`
}

// 用户基本信息
type Users struct {
	Id       int        `json:"id";gorm:"primaryKey"`
	Name     string     `json:"name"`
	Avatar   string     `json:"avatar"`
	Password string     `json:"password"`
	Status   userStatus `json:"status"`
	Ctime    int64      `json:"ctime"`
}

// 登录成功的用户信息
type UserInfo struct {
	UserInfo map[string]interface{} `json:"user_info"`
	Token    string                 `json:"token"`
}

// 登录成功
type LoginSuccess struct {
	Code    int      `json:"code"`
	Data    UserInfo `json:"data"`
	Message string   `json:"message"`
}

// 登陆失败
type LoginFailed struct {
	Code    int
	Message string
}

