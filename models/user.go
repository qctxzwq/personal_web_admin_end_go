package models

type UserStatus int

const (
	SUPER_ADMIN       UserStatus = 0
	NORMAL_ADMIN      UserStatus = 1
	SUPER_USER        UserStatus = 2
	NORMAL_USER       UserStatus = 3
	FORBIDEN_USER_ONE UserStatus = 4
	FORBIDEN_USER_TWO UserStatus = 5
	VISITOR           UserStatus = 6
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
	Name     string `json:"name";form:"name"`
	Password string `json:"password";form:"password"`
	Confirm string `json:"confirm";form:"confirm"`
	Status UserStatus `json:"status";form:"status"`
	Avatar string `json:"avatar"`
}

// 用户基本信息
type Users struct {
	Id       int        `json:"id";gorm:"primaryKey"`
	Name     string     `json:"name"`
	Avatar   string     `json:"avatar"`
	Password string     `json:"password"`
	Status   UserStatus `json:"status"`
	Ctime    int64      `json:"ctime"`
}

// 无密码的用户信息
type UserNoPwd struct {
	Id     int        `json:"id";gorm:"primaryKey"`
	Name   string     `json:"name"`
	Avatar string     `json:"avatar"`
	Status UserStatus `json:"status"`
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
