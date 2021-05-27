package models

var RouteAuthMap = make(map[string]UserStatus)

func InitRouteAuth() {
	RouteAuthMap["/login"] = VISITOR
	RouteAuthMap["/user/all"] = VISITOR
}
