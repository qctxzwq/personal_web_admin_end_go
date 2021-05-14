package models

type Home struct{
	Code	int 	`json:"code"`
	Data 	HomeMessage `json:"data"`
	ReqTime string 	`json:"request_time"`
}

type HomeMessage struct{
	Username string
	Message string
}
