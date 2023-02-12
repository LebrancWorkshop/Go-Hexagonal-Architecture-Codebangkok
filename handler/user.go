package handler 

import (
	"github.com/LebrancWorkshop/Go-Hexagonal-Architecture-Codebangkok/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}; 
}