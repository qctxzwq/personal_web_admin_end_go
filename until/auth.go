package until

func IsHaveAuth(reqStatue,routeAuth int) bool{
	return reqStatue <= routeAuth
}