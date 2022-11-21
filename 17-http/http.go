package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/home", homePage) //registering the route // and registering the handler which will handle the req
	http.ListenAndServe(":8080", nil)  // this will start the server //

}

// homePage
// http.ResponseWriter is used to write response back to the client ,
// http.Request is used to fetch any request specific details like json, body , or anything related to request data
func homePage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "welcome to our home page")
	w.Write([]byte("welcome to our home page"))
	fmt.Println("home page is up")
}
