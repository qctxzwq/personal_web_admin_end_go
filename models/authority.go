package models

var RouteAuthMap = make(map[string]userStatus)

func InitRouteAuth() {
	RouteAuthMap["/login"] = VISITOR
	RouteAuthMap["/user/all"] = VISITOR
}
