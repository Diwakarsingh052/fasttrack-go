package main

import (
	"mvc-app/controllers"
	"net/http"
)

/*
	m -> model // to interact with the db
	v -> views
	c -> controllers // this is the entry point of our app // it will have HandleFuncs
*/

func main() {
	http.HandleFunc("/user", controllers.GetUser)
	http.ListenAndServe(":8080", nil)

}
