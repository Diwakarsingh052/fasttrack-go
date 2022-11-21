package main

import (
	"fmt"
	"log"
	"net/http"
)

/*

	log mid ->  handlerFunc -> DOSomeWorkFunc()
	//log mid ->  error mid -> handlerFunc -> DOSomeWorkFunc()
*/

func main() {

	http.HandleFunc("/home", LoggingMid(home))
	panic(http.ListenAndServe(":8080", nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("In home Page handler")
	fmt.Fprintln(w, "hello this is our first home page")

}

func LoggingMid(next http.HandlerFunc) http.HandlerFunc { // return type means exec next thing in the chain
	return func(w http.ResponseWriter, r *http.Request) {
		//do middleware specific stuff first and when it is over then go to the actual handler func to exec it
		log.Println("started")
		defer log.Println("ended")
		log.Println(r.Method)

		if r.Method != http.MethodGet {
			http.Error(w, "method must be get", http.StatusInternalServerError)
			return
		}

		// actual handler request
		next(w, r) // closure
	}

}
