package service 

import (

)

type UserResponse struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Displayname string `json:"displayname"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUserByID(int) (*UserResponse, error)
}