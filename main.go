package main

import (
	"fmt"
	"log"
	"net/http"
	"rest/controller"
	"rest/database"
	"rest/service"
)

type Ad struct {
	Title   string  `json:"title"`
	Text    string  `json:"text"`
	Image   string  `json:"image"`
	Price   float64 `json:"price"`
	OwnerID string  `json:"owner_id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {

	db, err := database.DbInit()
	if err != nil {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}

	userService := service.NewUserService(db)
	roleService := service.NewRoleService(db)
	advService := service.NewAdvService(db)

	userController := controller.NewUserController(userService, roleService)
	advController := controller.NewAdvController(advService, userService)

	http.HandleFunc("/login", userController.LoginUser)
	http.HandleFunc("/register", userController.CreateUser)
	http.HandleFunc("/add/adv", advController.AddAdv)
	http.HandleFunc("/show/adv", advController.GetAdv)

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
