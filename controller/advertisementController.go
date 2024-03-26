package controller

import (
	"encoding/json"
	"net/http"
	"rest/service"
)

type AdvController struct {
	advService  *service.AdvService
	userService *service.UserService
}

func NewAdvController(advService *service.AdvService, userService *service.UserService) *AdvController {
	return &AdvController{
		advService:  advService,
		userService: userService,
	}
}

func (c *AdvController) AddAdv(w http.ResponseWriter, r *http.Request) {
	adv, status, err := c.advService.AddAdv(r, c.userService)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err = json.NewEncoder(w).Encode(adv)
	if err != nil {
		http.Error(w, "can not parse output: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *AdvController) GetAdv(w http.ResponseWriter, r *http.Request) {
	advsDto, status, err := c.advService.GetAdv(r, c.userService)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err = json.NewEncoder(w).Encode(advsDto)
	if err != nil {
		http.Error(w, "can not parse output: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
