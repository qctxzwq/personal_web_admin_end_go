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

// 注册必传字段
type Register struct {
	Name     string `json:"name"`
	Password string `json:"password"`
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

// 无密码的用户信息
type UserNoPwd struct {
	Id     int        `json:"id";gorm:"primaryKey"`
	Name   string     `json:"name"`
	Avatar string     `json:"avatar"`
	Status userStatus `json:"status"`
	Ctime  int64      `json:"ctime"`
}

// 登录成功的用户信息
type UserInfo struct {
	UserInfo *UserNoPwd `json:"user_info"`
	Token    string     `json:"token"`
}

// 获取的全部用户结构体
type List struct {
	Total    int         `json:"total"`
	UserList []UserNoPwd `json:"user_info"`
}

// 接口请求成功
type SuccessMsg struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// 接口请求失败
type FailedMsg struct {
	Code    int
	Message string
}
