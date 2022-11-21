package main

import (
	"mvc-app/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.Home) // routes
	http.HandleFunc("/users", controller.GetUser)
	http.ListenAndServe(":8080", nil)

}
