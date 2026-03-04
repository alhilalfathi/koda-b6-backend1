package models

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var NextId = 1

var UserList []Users
