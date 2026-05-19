package rout

import (
	"fmt"
	"net/http"
	sozd "proj1/func"
	handler "proj1/handler"
)

func Rout() {
	http.HandleFunc("/", sozd.AuthMiddleware(handler.FormHandler))
	http.HandleFunc("/add", sozd.AuthMiddleware(handler.AddHandler))
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/users", sozd.AuthMiddleware(handler.UsersHandler))
	http.HandleFunc("/logout", handler.LogoutHandler)
	http.HandleFunc("/registr", handler.RegistrHandler)
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
