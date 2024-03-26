package controller

import (
	"encoding/json"
	"net/http"
	"rest/service"
)

type UserRoleController struct {
	userService *service.UserService
	roleService *service.RoleService
}

func NewUserController(userService *service.UserService, roleService *service.RoleService) *UserRoleController {
	return &UserRoleController{
		userService: userService,
		roleService: roleService,
	}
}

func (c *UserRoleController) CreateUser(w http.ResponseWriter, r *http.Request) {
	token, status, err := service.CreateUser(r, c.userService, c.roleService)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		http.Error(w, "can not parse output: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UserRoleController) LoginUser(w http.ResponseWriter, r *http.Request) {
	token, status, err := c.userService.LoginUser(r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		http.Error(w, "can not parse output: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
