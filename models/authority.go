package models

var RouteAuthMap = make(map[string]UserStatus)

func InitRouteAuth() {
	RouteAuthMap["/login"] = VISITOR
	RouteAuthMap["/register"] = VISITOR
	RouteAuthMap["/upload/img"] = VISITOR
	RouteAuthMap["/user/all"] = VISITOR
	RouteAuthMap["/user/update"] = VISITOR
}
